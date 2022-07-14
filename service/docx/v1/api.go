// Package docx code generated by oapi sdk gen
package larkdocx

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go.v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *DocxService {
	d := &DocxService{config: config}
	d.Document = &document{service: d}
	d.DocumentBlock = &documentBlock{service: d}
	d.DocumentBlockChildren = &documentBlockChildren{service: d}
	return d
}

// 业务域服务定义
type DocxService struct {
	config                *larkcore.Config
	Document              *document
	DocumentBlock         *documentBlock
	DocumentBlockChildren *documentBlockChildren
}

// 资源服务定义
type document struct {
	service *DocxService
}
type documentBlock struct {
	service *DocxService
}
type documentBlockChildren struct {
	service *DocxService
}

// 资源服务方法定义
func (d *document) Create(ctx context.Context, req *CreateDocumentReq, options ...larkcore.RequestOptionFunc) (*CreateDocumentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/docx/v1/documents", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDocumentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *document) Get(ctx context.Context, req *GetDocumentReq, options ...larkcore.RequestOptionFunc) (*GetDocumentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *document) RawContent(ctx context.Context, req *RawContentDocumentReq, options ...larkcore.RequestOptionFunc) (*RawContentDocumentResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/raw_content", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RawContentDocumentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) BatchUpdate(ctx context.Context, req *BatchUpdateDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*BatchUpdateDocumentBlockResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPatch,
		"/open-apis/docx/v1/documents/:document_id/blocks/batch_update", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchUpdateDocumentBlockResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) Get(ctx context.Context, req *GetDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentBlockResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) List(ctx context.Context, req *ListDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*ListDocumentBlockResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDocumentBlockResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) ListByIterator(ctx context.Context, req *ListDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*ListDocumentBlockIterator, error) {
	return &ListDocumentBlockIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *documentBlock) Patch(ctx context.Context, req *PatchDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*PatchDocumentBlockResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPatch,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDocumentBlockResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) BatchDelete(ctx context.Context, req *BatchDeleteDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*BatchDeleteDocumentBlockChildrenResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodDelete,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteDocumentBlockChildrenResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Create(ctx context.Context, req *CreateDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*CreateDocumentBlockChildrenResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDocumentBlockChildrenResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Get(ctx context.Context, req *GetDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockChildrenResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentBlockChildrenResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) GetByIterator(ctx context.Context, req *GetDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockChildrenIterator, error) {
	return &GetDocumentBlockChildrenIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.Get,
		options:  options,
		limit:    req.Limit}, nil
}
