// Code generated by zenrpc v2.2.9; DO NOT EDIT.

package vfs

import (
	"context"
	"encoding/json"

	"github.com/vmkteam/zenrpc/v2"
	"github.com/vmkteam/zenrpc/v2/smd"
)

var RPC = struct {
	Service struct{ GetFolder, GetFolderBranch, GetFiles, CountFiles, MoveFiles, DeleteFiles, SetFilePhysicalName, SearchFolderByFileId, SearchFolderByFile, GetFavorites, ManageFavorites, CreateFolder, DeleteFolder, MoveFolder, RenameFolder, HelpUpload, UrlByHash, UrlByHashList, DeleteHash string }
}{
	Service: struct{ GetFolder, GetFolderBranch, GetFiles, CountFiles, MoveFiles, DeleteFiles, SetFilePhysicalName, SearchFolderByFileId, SearchFolderByFile, GetFavorites, ManageFavorites, CreateFolder, DeleteFolder, MoveFolder, RenameFolder, HelpUpload, UrlByHash, UrlByHashList, DeleteHash string }{
		GetFolder:            "getfolder",
		GetFolderBranch:      "getfolderbranch",
		GetFiles:             "getfiles",
		CountFiles:           "countfiles",
		MoveFiles:            "movefiles",
		DeleteFiles:          "deletefiles",
		SetFilePhysicalName:  "setfilephysicalname",
		SearchFolderByFileId: "searchfolderbyfileid",
		SearchFolderByFile:   "searchfolderbyfile",
		GetFavorites:         "getfavorites",
		ManageFavorites:      "managefavorites",
		CreateFolder:         "createfolder",
		DeleteFolder:         "deletefolder",
		MoveFolder:           "movefolder",
		RenameFolder:         "renamefolder",
		HelpUpload:           "helpupload",
		UrlByHash:            "urlbyhash",
		UrlByHashList:        "urlbyhashlist",
		DeleteHash:           "deletehash",
	},
}

func (Service) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"GetFolder": {
				Description: `GetFolder returns Folder with sub folders.`,
				Parameters: []smd.JSONSchema{
					{
						Name:     "rootFolderId",
						Optional: true,
						Type:     smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					TypeName: "Folder",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name:     "parentId",
							Optional: true,
							Type:     smd.Integer,
						},
						{
							Name: "folders",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/Folder",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"Folder": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name:     "parentId",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name: "folders",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Folder",
									},
								},
							},
						},
					},
				},
				Errors: map[int]string{
					404: "Folder not found",
				},
			},
			"GetFolderBranch": {
				Description: `GetFolderBranch returns Folder branch.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "folderId",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type:     smd.Array,
					TypeName: "[]Folder",
					Items: map[string]string{
						"$ref": "#/definitions/Folder",
					},
					Definitions: map[string]smd.Definition{
						"Folder": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name:     "parentId",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name: "folders",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Folder",
									},
								},
							},
						},
					},
				},
			},
			"GetFiles": {
				Description: `GetFiles returns list of files.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "folderId",
						Description: `root folder id`,
						Type:        smd.Integer,
					},
					{
						Name:        "query",
						Optional:    true,
						Description: `file name`,
						Type:        smd.String,
					},
					{
						Name:        "sortField",
						Optional:    true,
						Description: `createdAt, title or fileSize`,
						Type:        smd.String,
					},
					{
						Name:        "isDescending",
						Optional:    true,
						Description: `asc = false, desc = true`,
						Type:        smd.Boolean,
					},
					{
						Name:        "page",
						Optional:    true,
						Description: `current page`,
						Type:        smd.Integer,
					},
					{
						Name:        "pageSize",
						Optional:    true,
						Description: `current pageSize`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type:     smd.Array,
					TypeName: "[]File",
					Items: map[string]string{
						"$ref": "#/definitions/File",
					},
					Definitions: map[string]smd.Definition{
						"File": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "path",
									Type: smd.String,
								},
								{
									Name: "previewpath",
									Type: smd.String,
								},
								{
									Name: "relpath",
									Type: smd.String,
								},
								{
									Name: "size",
									Type: smd.Integer,
								},
								{
									Name: "sizeH",
									Type: smd.Array,
									Items: map[string]string{
										"type": smd.String,
									},
								},
								{
									Name: "date",
									Type: smd.String,
								},
								{
									Name: "type",
									Type: smd.String,
								},
								{
									Name: "extension",
									Type: smd.String,
								},
								{
									Name: "params",
									Ref:  "#/definitions/FileParams",
									Type: smd.Object,
								},
								{
									Name: "shortpath",
									Type: smd.String,
								},
								{
									Name:     "width",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name:     "height",
									Optional: true,
									Type:     smd.Integer,
								},
							},
						},
						"FileParams": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "width",
									Type: smd.Integer,
								},
								{
									Name: "height",
									Type: smd.Integer,
								},
							},
						},
					},
				},
			},
			"CountFiles": {
				Description: `CountFiles returns count of files.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "folderId",
						Description: `root folder id`,
						Type:        smd.Integer,
					},
					{
						Name:        "query",
						Optional:    true,
						Description: `file name`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Integer,
				},
			},
			"MoveFiles": {
				Description: `MoveFiles move files to destination folder.`,
				Parameters: []smd.JSONSchema{
					{
						Name:     "fileIds",
						Type:     smd.Array,
						TypeName: "[]",
						Items: map[string]string{
							"type": smd.Integer,
						},
					},
					{
						Name: "destinationFolderId",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
				Errors: map[int]string{
					400: "empty file ids",
				},
			},
			"DeleteFiles": {
				Description: `DeleteFiles remove files.`,
				Parameters: []smd.JSONSchema{
					{
						Name:     "fileIds",
						Type:     smd.Array,
						TypeName: "[]",
						Items: map[string]string{
							"type": smd.Integer,
						},
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"SetFilePhysicalName": {
				Description: `SetFilePhysicalName renames File on server.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "fileId",
						Type: smd.Integer,
					},
					{
						Name: "name",
						Type: smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"SearchFolderByFileId": {
				Description: `SearchFolderByFileId return Folder by File id.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "fileId",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					TypeName: "Folder",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name:     "parentId",
							Optional: true,
							Type:     smd.Integer,
						},
						{
							Name: "folders",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/Folder",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"Folder": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name:     "parentId",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name: "folders",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Folder",
									},
								},
							},
						},
					},
				},
			},
			"SearchFolderByFile": {
				Description: `SearchFolderByFile return Folder by File name.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "filename",
						Type: smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					TypeName: "Folder",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name:     "parentId",
							Optional: true,
							Type:     smd.Integer,
						},
						{
							Name: "folders",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/Folder",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"Folder": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name:     "parentId",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name: "folders",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Folder",
									},
								},
							},
						},
					},
				},
			},
			"GetFavorites": {
				Description: `GetFavorites return favorites list.`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type:     smd.Array,
					TypeName: "[]Folder",
					Items: map[string]string{
						"$ref": "#/definitions/Folder",
					},
					Definitions: map[string]smd.Definition{
						"Folder": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name:     "parentId",
									Optional: true,
									Type:     smd.Integer,
								},
								{
									Name: "folders",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Folder",
									},
								},
							},
						},
					},
				},
			},
			"ManageFavorites": {
				Description: `ManageFavorites manage favorite virtual folders.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "folderId",
						Type: smd.Integer,
					},
					{
						Name: "isInFavorites",
						Type: smd.Boolean,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"CreateFolder": {
				Description: `CreateFolder create virtual folder.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "rootFolderId",
						Type: smd.Integer,
					},
					{
						Name: "name",
						Type: smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"DeleteFolder": {
				Description: `DeleteFolder removes Folder.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "folderId",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"MoveFolder": {
				Description: `MoveFolder move Folder to destination folder.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "folderId",
						Type: smd.Integer,
					},
					{
						Name: "destinationFolderId",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"RenameFolder": {
				Description: `RenameFolder change Folder name.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "folderId",
						Type: smd.Integer,
					},
					{
						Name: "name",
						Type: smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"HelpUpload": {
				Description: `HelpUpload returns a uploader help info.`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type:     smd.Object,
					TypeName: "HelpUploadResponse",
					Properties: smd.PropertyList{
						{
							Name: "temp",
							Ref:  "#/definitions/HelpUploadItem",
							Type: smd.Object,
						},
						{
							Name: "queue",
							Ref:  "#/definitions/HelpUploadItem",
							Type: smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"HelpUploadItem": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "url",
									Type: smd.String,
								},
								{
									Name: "params",
									Type: smd.Array,
									Items: map[string]string{
										"type": smd.String,
									},
								},
							},
						},
					},
				},
			},
			"UrlByHash": {
				Description: `UrlByHash get Url by hash, namespace and media type`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "hash",
						Description: `media hash`,
						Type:        smd.String,
					},
					{
						Name:        "namespace",
						Description: `media namespace`,
						Type:        smd.String,
					},
					{
						Name:        "mediaType",
						Description: `type of media (possible values: small, medium, big, empty string)`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
			"UrlByHashList": {
				Description: `UrlByHashList get Urls by hash list, with namespace and media type`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "hashList",
						Description: `media hash list`,
						Type:        smd.Array,
						TypeName:    "[]",
						Items: map[string]string{
							"type": smd.String,
						},
					},
					{
						Name:        "namespace",
						Description: `media namespace`,
						Type:        smd.String,
					},
					{
						Name:        "mediaType",
						Description: `type of media (possible values: small, medium, big, empty string)`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type:     smd.Array,
					TypeName: "[]UrlByHashListResponse",
					Items: map[string]string{
						"$ref": "#/definitions/UrlByHashListResponse",
					},
					Definitions: map[string]smd.Definition{
						"UrlByHashListResponse": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "hash",
									Type: smd.String,
								},
								{
									Name: "webPath",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"DeleteHash": {
				Description: `DeleteHash delete file by namespace, hash and extension.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `media namespace`,
						Type:        smd.String,
					},
					{
						Name:        "hash",
						Description: `media hash`,
						Type:        smd.String,
					},
					{
						Name:        "ext",
						Description: `media extension`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
				Errors: map[int]string{
					404: "File not found by hash",
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s Service) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.Service.GetFolder:
		var args = struct {
			RootFolderId *int `json:"rootFolderId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"rootFolderId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:rootFolderId=1
		if args.RootFolderId == nil {
			var v int = 1
			args.RootFolderId = &v
		}

		resp.Set(s.GetFolder(ctx, *args.RootFolderId))

	case RPC.Service.GetFolderBranch:
		var args = struct {
			FolderId int `json:"folderId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GetFolderBranch(ctx, args.FolderId))

	case RPC.Service.GetFiles:
		var args = struct {
			FolderId     int     `json:"folderId"`
			Query        *string `json:"query"`
			SortField    *string `json:"sortField"`
			IsDescending *bool   `json:"isDescending"`
			Page         *int    `json:"page"`
			PageSize     *int    `json:"pageSize"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId", "query", "sortField", "isDescending", "page", "pageSize"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:isDescending=true asc = false, desc = true
		if args.IsDescending == nil {
			var v bool = true
			args.IsDescending = &v
		}

		//zenrpc:page=0 current page
		if args.Page == nil {
			var v int = 0
			args.Page = &v
		}

		//zenrpc:pageSize=100 current pageSize
		if args.PageSize == nil {
			var v int = 100
			args.PageSize = &v
		}

		//zenrpc:sortField="createdAt" createdAt, title or fileSize
		if args.SortField == nil {
			var v string = "createdAt"
			args.SortField = &v
		}

		resp.Set(s.GetFiles(ctx, args.FolderId, args.Query, *args.SortField, *args.IsDescending, *args.Page, *args.PageSize))

	case RPC.Service.CountFiles:
		var args = struct {
			FolderId int     `json:"folderId"`
			Query    *string `json:"query"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId", "query"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.CountFiles(ctx, args.FolderId, args.Query))

	case RPC.Service.MoveFiles:
		var args = struct {
			FileIds             []int64 `json:"fileIds"`
			DestinationFolderId int     `json:"destinationFolderId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"fileIds", "destinationFolderId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.MoveFiles(ctx, args.FileIds, args.DestinationFolderId))

	case RPC.Service.DeleteFiles:
		var args = struct {
			FileIds []int64 `json:"fileIds"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"fileIds"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.DeleteFiles(ctx, args.FileIds))

	case RPC.Service.SetFilePhysicalName:
		var args = struct {
			FileId int    `json:"fileId"`
			Name   string `json:"name"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"fileId", "name"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.SetFilePhysicalName(ctx, args.FileId, args.Name))

	case RPC.Service.SearchFolderByFileId:
		var args = struct {
			FileId int `json:"fileId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"fileId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.SearchFolderByFileId(ctx, args.FileId))

	case RPC.Service.SearchFolderByFile:
		var args = struct {
			Filename string `json:"filename"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"filename"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.SearchFolderByFile(ctx, args.Filename))

	case RPC.Service.GetFavorites:
		resp.Set(s.GetFavorites(ctx))

	case RPC.Service.ManageFavorites:
		var args = struct {
			FolderId      int  `json:"folderId"`
			IsInFavorites bool `json:"isInFavorites"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId", "isInFavorites"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.ManageFavorites(ctx, args.FolderId, args.IsInFavorites))

	case RPC.Service.CreateFolder:
		var args = struct {
			RootFolderId int    `json:"rootFolderId"`
			Name         string `json:"name"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"rootFolderId", "name"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.CreateFolder(ctx, args.RootFolderId, args.Name))

	case RPC.Service.DeleteFolder:
		var args = struct {
			FolderId int `json:"folderId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.DeleteFolder(ctx, args.FolderId))

	case RPC.Service.MoveFolder:
		var args = struct {
			FolderId            int `json:"folderId"`
			DestinationFolderId int `json:"destinationFolderId"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId", "destinationFolderId"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.MoveFolder(ctx, args.FolderId, args.DestinationFolderId))

	case RPC.Service.RenameFolder:
		var args = struct {
			FolderId int    `json:"folderId"`
			Name     string `json:"name"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"folderId", "name"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.RenameFolder(ctx, args.FolderId, args.Name))

	case RPC.Service.HelpUpload:
		resp.Set(s.HelpUpload())

	case RPC.Service.UrlByHash:
		var args = struct {
			Hash      string `json:"hash"`
			Namespace string `json:"namespace"`
			MediaType string `json:"mediaType"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"hash", "namespace", "mediaType"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.UrlByHash(ctx, args.Hash, args.Namespace, args.MediaType))

	case RPC.Service.UrlByHashList:
		var args = struct {
			HashList  []string `json:"hashList"`
			Namespace string   `json:"namespace"`
			MediaType string   `json:"mediaType"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"hashList", "namespace", "mediaType"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.UrlByHashList(ctx, args.HashList, args.Namespace, args.MediaType))

	case RPC.Service.DeleteHash:
		var args = struct {
			Namespace string `json:"namespace"`
			Hash      string `json:"hash"`
			Ext       string `json:"ext"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "hash", "ext"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.DeleteHash(ctx, args.Namespace, args.Hash, args.Ext))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
