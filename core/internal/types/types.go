// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type MailCodeSendReply struct {
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime            int    `json:"expired_time"`
}

type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity,optional"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileDeleteReply struct {
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFileListRequest struct {
	Identity string `json:"identity,optional"`
	Page     int    `json:"page,optional"`
	Size     int    `json:"size,optional"`
}

type UserFileMoveReply struct {
}

type UserFileMoveRequest struct {
	Idnetity       string `json:"identity"`
	ParentIdnetity string `json:"parent_identity"`
}

type UserFileNameUpdateReply struct {
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFolder struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFolderCreateReply struct {
	Identity string `json:"identity"`
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderListReply struct {
	List []*UserFolder `json:"list"`
}

type UserFolderListRequest struct {
	Identity string `json:"identity,optional"`
}

type UserRegisterReply struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRepositorySaveReply struct {
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}
