// Package drive code generated by oapi sdk gen
package larkdrive

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go.v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *DriveService {
	d := &DriveService{config: config}
	d.ExportTask = &exportTask{service: d}
	d.File = &file{service: d}
	d.FileComment = &fileComment{service: d}
	d.FileCommentReply = &fileCommentReply{service: d}
	d.FileStatistics = &fileStatistics{service: d}
	d.FileSubscription = &fileSubscription{service: d}
	d.ImportTask = &importTask{service: d}
	d.Media = &media{service: d}
	d.Meta = &meta{service: d}
	d.PermissionMember = &permissionMember{service: d}
	d.PermissionPublic = &permissionPublic{service: d}
	return d
}

// 业务域服务定义
type DriveService struct {
	config           *larkcore.Config
	ExportTask       *exportTask
	File             *file
	FileComment      *fileComment
	FileCommentReply *fileCommentReply
	FileStatistics   *fileStatistics
	FileSubscription *fileSubscription
	ImportTask       *importTask
	Media            *media
	Meta             *meta
	PermissionMember *permissionMember
	PermissionPublic *permissionPublic
}

// 资源服务定义
type exportTask struct {
	service *DriveService
}
type file struct {
	service *DriveService
}
type fileComment struct {
	service *DriveService
}
type fileCommentReply struct {
	service *DriveService
}
type fileStatistics struct {
	service *DriveService
}
type fileSubscription struct {
	service *DriveService
}
type importTask struct {
	service *DriveService
}
type media struct {
	service *DriveService
}
type meta struct {
	service *DriveService
}
type permissionMember struct {
	service *DriveService
}
type permissionPublic struct {
	service *DriveService
}

// 资源服务方法定义
func (e *exportTask) Create(ctx context.Context, req *CreateExportTaskReq, options ...larkcore.RequestOptionFunc) (*CreateExportTaskResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/drive/v1/export_tasks", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exportTask) Download(ctx context.Context, req *DownloadExportTaskReq, options ...larkcore.RequestOptionFunc) (*DownloadExportTaskResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/drive/v1/export_tasks/file/:file_token/download", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadExportTaskResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exportTask) Get(ctx context.Context, req *GetExportTaskReq, options ...larkcore.RequestOptionFunc) (*GetExportTaskResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/drive/v1/export_tasks/:ticket", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetExportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Copy(ctx context.Context, req *CopyFileReq, options ...larkcore.RequestOptionFunc) (*CopyFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/copy", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CopyFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) CreateFolder(ctx context.Context, req *CreateFolderFileReq, options ...larkcore.RequestOptionFunc) (*CreateFolderFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/create_folder", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFolderFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Delete(ctx context.Context, req *DeleteFileReq, options ...larkcore.RequestOptionFunc) (*DeleteFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodDelete,
		"/open-apis/drive/v1/files/:file_token", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Download(ctx context.Context, req *DownloadFileReq, options ...larkcore.RequestOptionFunc) (*DownloadFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/download", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadFileResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) List(ctx context.Context, req *ListFileReq, options ...larkcore.RequestOptionFunc) (*ListFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Move(ctx context.Context, req *MoveFileReq, options ...larkcore.RequestOptionFunc) (*MoveFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/move", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Subscribe(ctx context.Context, req *SubscribeFileReq, options ...larkcore.RequestOptionFunc) (*SubscribeFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/subscribe", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) TaskCheck(ctx context.Context, req *TaskCheckFileReq, options ...larkcore.RequestOptionFunc) (*TaskCheckFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/task_check", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TaskCheckFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadAll(ctx context.Context, req *UploadAllFileReq, options ...larkcore.RequestOptionFunc) (*UploadAllFileResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_all", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadAllFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadFinish(ctx context.Context, req *UploadFinishFileReq, options ...larkcore.RequestOptionFunc) (*UploadFinishFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_finish", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFinishFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadPart(ctx context.Context, req *UploadPartFileReq, options ...larkcore.RequestOptionFunc) (*UploadPartFileResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_part", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPartFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadPrepare(ctx context.Context, req *UploadPrepareFileReq, options ...larkcore.RequestOptionFunc) (*UploadPrepareFileResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_prepare", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPrepareFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) Create(ctx context.Context, req *CreateFileCommentReq, options ...larkcore.RequestOptionFunc) (*CreateFileCommentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/comments", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) Get(ctx context.Context, req *GetFileCommentReq, options ...larkcore.RequestOptionFunc) (*GetFileCommentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) List(ctx context.Context, req *ListFileCommentReq, options ...larkcore.RequestOptionFunc) (*ListFileCommentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/comments", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) ListByIterator(ctx context.Context, req *ListFileCommentReq, options ...larkcore.RequestOptionFunc) (*ListFileCommentIterator, error) {
	return &ListFileCommentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: f.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (f *fileComment) Patch(ctx context.Context, req *PatchFileCommentReq, options ...larkcore.RequestOptionFunc) (*PatchFileCommentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPatch,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileCommentReply) Delete(ctx context.Context, req *DeleteFileCommentReplyReq, options ...larkcore.RequestOptionFunc) (*DeleteFileCommentReplyResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodDelete,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFileCommentReplyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileCommentReply) Update(ctx context.Context, req *UpdateFileCommentReplyReq, options ...larkcore.RequestOptionFunc) (*UpdateFileCommentReplyResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPut,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFileCommentReplyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileStatistics) Get(ctx context.Context, req *GetFileStatisticsReq, options ...larkcore.RequestOptionFunc) (*GetFileStatisticsResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/statistics", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileStatisticsResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Create(ctx context.Context, req *CreateFileSubscriptionReq, options ...larkcore.RequestOptionFunc) (*CreateFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/subscriptions", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Get(ctx context.Context, req *GetFileSubscriptionReq, options ...larkcore.RequestOptionFunc) (*GetFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Patch(ctx context.Context, req *PatchFileSubscriptionReq, options ...larkcore.RequestOptionFunc) (*PatchFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, f.service.config, http.MethodPatch,
		"/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *importTask) Create(ctx context.Context, req *CreateImportTaskReq, options ...larkcore.RequestOptionFunc) (*CreateImportTaskResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, i.service.config, http.MethodPost,
		"/open-apis/drive/v1/import_tasks", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateImportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *importTask) Get(ctx context.Context, req *GetImportTaskReq, options ...larkcore.RequestOptionFunc) (*GetImportTaskResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, i.service.config, http.MethodGet,
		"/open-apis/drive/v1/import_tasks/:ticket", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetImportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) BatchGetTmpDownloadUrl(ctx context.Context, req *BatchGetTmpDownloadUrlMediaReq, options ...larkcore.RequestOptionFunc) (*BatchGetTmpDownloadUrlMediaResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/drive/v1/medias/batch_get_tmp_download_url", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchGetTmpDownloadUrlMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) Download(ctx context.Context, req *DownloadMediaReq, options ...larkcore.RequestOptionFunc) (*DownloadMediaResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/drive/v1/medias/:file_token/download", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadMediaResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadAll(ctx context.Context, req *UploadAllMediaReq, options ...larkcore.RequestOptionFunc) (*UploadAllMediaResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_all", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadAllMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadFinish(ctx context.Context, req *UploadFinishMediaReq, options ...larkcore.RequestOptionFunc) (*UploadFinishMediaResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_finish", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFinishMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadPart(ctx context.Context, req *UploadPartMediaReq, options ...larkcore.RequestOptionFunc) (*UploadPartMediaResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_part", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPartMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadPrepare(ctx context.Context, req *UploadPrepareMediaReq, options ...larkcore.RequestOptionFunc) (*UploadPrepareMediaResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_prepare", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPrepareMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meta) BatchQuery(ctx context.Context, req *BatchQueryMetaReq, options ...larkcore.RequestOptionFunc) (*BatchQueryMetaResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/metas/batch_query", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryMetaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Create(ctx context.Context, req *CreatePermissionMemberReq, options ...larkcore.RequestOptionFunc) (*CreatePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, p.service.config, http.MethodPost,
		"/open-apis/drive/v1/permissions/:token/members", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreatePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Delete(ctx context.Context, req *DeletePermissionMemberReq, options ...larkcore.RequestOptionFunc) (*DeletePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, p.service.config, http.MethodDelete,
		"/open-apis/drive/v1/permissions/:token/members/:member_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeletePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Update(ctx context.Context, req *UpdatePermissionMemberReq, options ...larkcore.RequestOptionFunc) (*UpdatePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, p.service.config, http.MethodPut,
		"/open-apis/drive/v1/permissions/:token/members/:member_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdatePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionPublic) Get(ctx context.Context, req *GetPermissionPublicReq, options ...larkcore.RequestOptionFunc) (*GetPermissionPublicResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, p.service.config, http.MethodGet,
		"/open-apis/drive/v1/permissions/:token/public", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetPermissionPublicResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionPublic) Patch(ctx context.Context, req *PatchPermissionPublicReq, options ...larkcore.RequestOptionFunc) (*PatchPermissionPublicResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, p.service.config, http.MethodPatch,
		"/open-apis/drive/v1/permissions/:token/public", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchPermissionPublicResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
