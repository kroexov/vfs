package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/namsral/flag"
	"github.com/semrush/zenrpc"

	"github.com/vmkteam/vfs"
	"github.com/vmkteam/vfs/db"
)

var (
	fs            = flag.NewFlagSetWithEnvPrefix(os.Args[0], "VFS", 0)
	flAddr        = fs.String("addr", "localhost:9999", "listen address")
	flDir         = fs.String("dir", "testdata", "storage path")
	flNamespaces  = fs.String("ns", "items,test", "namespaces, separated by comma")
	flWebPath     = fs.String("webpath", "/media/", "web path to files")
	flPreviewPath = fs.String("preview-path", "/media/small/", "preview path to image files")
	flDbConn      = fs.String("conn", "postgresql://localhost:5432/vfs?sslmode=disable", "database connection dsn")
	flJWTKey      = fs.String("jwt-key", "QuiuNae9OhzoKohcee0h", "JWT key")
	flJWTHeader   = fs.String("jwt-header", "AuthorizationJWT", "JWT header")
	flFileSize    = fs.Int64("maxsize", 32<<20, "max file size in bytes")
	flVerboseSQL  = fs.Bool("verbose-sql", false, "log all sql queries")
	version       string
)

func main() {
	err := fs.Parse(os.Args[1:])
	checkErr(err)

	v, err := vfs.New(vfs.Config{
		MaxFileSize:    *flFileSize,
		Path:           *flDir,
		WebPath:        *flWebPath,
		PreviewPath:    *flPreviewPath,
		UploadFormName: "Filedata",
		Namespaces:     strings.Split(*flNamespaces, ","),
		Database:       nil,
	})
	checkErr(err)

	log.Printf("starting vfssrv version=%v on %s jwt.header=%v", version, *flAddr, *flJWTHeader)

	// use rpc only when dbconn is set
	repo, dbc := initRepo()
	if repo != nil {
		rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true, AllowCORS: true})
		rpc.Use(zenrpc.Logger(log.New(os.Stdout, "", log.LstdFlags)))
		rpc.Register("", vfs.NewService(*repo, v, dbc))
		rpc.Register("vfs", vfs.NewService(*repo, v, dbc))

		http.Handle("/rpc", corsMiddleware(authMiddleware(rpc)))
		http.Handle("/upload/file", corsMiddleware(authMiddleware(v.UploadHandler(*repo))))
		http.Handle("/rpc/api.ts", corsMiddleware(typeScriptClientHandler(rpc)))
	}

	http.HandleFunc("/auth-token", issueTokenHandler)

	http.Handle("/upload/hash", corsMiddleware(authMiddleware(http.HandlerFunc(v.HashUploadHandler))))
	http.Handle(*flWebPath, http.StripPrefix(*flWebPath, http.FileServer(http.Dir(*flDir))))

	checkErr(http.ListenAndServe(*flAddr, nil))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		allowHeaders := "Authorization, Authorization2, Origin, X-Requested-With, Content-Type, Accept, Platform, Version"
		if *flJWTHeader != "" {
			allowHeaders += ", " + *flJWTHeader
		}
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// typeScriptClientHandler returns typescript API client
func typeScriptClientHandler(srv zenrpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bb, err := vfs.NewTypeScriptClient(srv.SMD()).Run()
		if err != nil {
			log.Printf("failed to convert type script err=%q", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		_, err = w.Write(bb)
		if err != nil {
			log.Printf("failed to write type script err=%q", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

// issueTokenHandler issues new jwt token for 1 hour. Subject can be set by id GET/POST param
func issueTokenHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "vfs",
		Subject:   id,
	})

	key := []byte(*flJWTKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, tokenString)
		fmt.Printf("issued new token=%v for id=%v", tokenString, id)
	}
}

// authMiddleware checks JWT token if set in flag jwt.header.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			isOK    = true
			errMsg  = ""
			errCode = http.StatusUnauthorized
		)

		defer func() {
			if isOK {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, errMsg, errCode)
			}
		}()

		if *flJWTHeader != "" {
			isOK = false
			tokenString := r.Header.Get(*flJWTHeader)
			token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(*flJWTKey), nil
			})

			if token == nil {
				errMsg = "missing token"
				return
			} else if err != nil {
				errMsg = err.Error()
			}

			// validate token
			if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
				if err = claims.Valid(); err != nil {
					errMsg, errCode = err.Error(), http.StatusForbidden
				} else {
					isOK = true
				}
			} else {
				errMsg = "bad token"
			}
		}
	})
}

// initRepo connects to postgres and inits vfs db repo.
func initRepo() (*db.VfsRepo, *pg.DB) {
	if flDbConn == nil {
		return nil, nil
	}

	cfg, err := pg.ParseURL(*flDbConn)
	checkErr(err)

	dbc := pg.Connect(cfg)
	d := db.New(dbc)
	v, err := d.Version()
	checkErr(err)

	log.Println(v)
	repo := db.NewVfsRepo(d)

	if *flVerboseSQL {
		dbc.AddQueryHook(dbLogger{})
	}

	return &repo, dbc
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}
func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	log.Println(q.FormattedQuery())
	return nil
}
