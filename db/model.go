//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"time"
)

var Columns = struct {
	VfsFile struct {
		ID, FolderID, Title, Path, Params, IsFavorite, MimeType, FileSize, FileExists, CreatedAt, StatusID string

		Folder string
	}
	VfsFolder struct {
		ID, ParentFolderID, Title, IsFavorite, CreatedAt, StatusID string

		ParentFolder string
	}
}{
	VfsFile: struct {
		ID, FolderID, Title, Path, Params, IsFavorite, MimeType, FileSize, FileExists, CreatedAt, StatusID string

		Folder string
	}{
		ID:         "fileId",
		FolderID:   "folderId",
		Title:      "title",
		Path:       "path",
		Params:     "params",
		IsFavorite: "isFavorite",
		MimeType:   "mimeType",
		FileSize:   "fileSize",
		FileExists: "fileExists",
		CreatedAt:  "createdAt",
		StatusID:   "statusId",

		Folder: "Folder",
	},
	VfsFolder: struct {
		ID, ParentFolderID, Title, IsFavorite, CreatedAt, StatusID string

		ParentFolder string
	}{
		ID:             "folderId",
		ParentFolderID: "parentFolderId",
		Title:          "title",
		IsFavorite:     "isFavorite",
		CreatedAt:      "createdAt",
		StatusID:       "statusId",

		ParentFolder: "ParentFolder",
	},
}

var Tables = struct {
	VfsFile struct {
		Name, Alias string
	}
	VfsFolder struct {
		Name, Alias string
	}
}{
	VfsFile: struct {
		Name, Alias string
	}{
		Name:  "vfsFiles",
		Alias: "t",
	},
	VfsFolder: struct {
		Name, Alias string
	}{
		Name:  "vfsFolders",
		Alias: "t",
	},
}

type VfsFile struct {
	tableName struct{} `pg:"\"vfsFiles\",alias:t,,discard_unknown_columns"`

	ID         int            `pg:"fileId,pk"`
	FolderID   int            `pg:"folderId,use_zero"`
	Title      string         `pg:"title,use_zero"`
	Path       string         `pg:"path,use_zero"`
	Params     *VfsFileParams `pg:"params"`
	IsFavorite *bool          `pg:"isFavorite"`
	MimeType   string         `pg:"mimeType,use_zero"`
	FileSize   *int           `pg:"fileSize"`
	FileExists bool           `pg:"fileExists,use_zero"`
	CreatedAt  time.Time      `pg:"createdAt,use_zero"`
	StatusID   int            `pg:"statusId,use_zero"`

	Folder *VfsFolder `pg:"fk:folderId"`
}

type VfsFolder struct {
	tableName struct{} `pg:"\"vfsFolders\",alias:t,,discard_unknown_columns"`

	ID             int       `pg:"folderId,pk"`
	ParentFolderID *int      `pg:"parentFolderId"`
	Title          string    `pg:"title,use_zero"`
	IsFavorite     *bool     `pg:"isFavorite"`
	CreatedAt      time.Time `pg:"createdAt,use_zero"`
	StatusID       int       `pg:"statusId,use_zero"`

	ParentFolder *VfsFolder `pg:"fk:parentFolderId"`
}
