// Package application code generated by oapi sdk gen
package larkapplication

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go.v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *ApplicationService {
	a := &ApplicationService{config: config}
	a.Application = &application{service: a}
	a.ApplicationAppUsage = &applicationAppUsage{service: a}
	a.ApplicationAppVersion = &applicationAppVersion{service: a}
	a.ApplicationFeedback = &applicationFeedback{service: a}
	a.ApplicationVisibility = &applicationVisibility{service: a}
	return a
}

// 业务域服务定义
type ApplicationService struct {
	config                *larkcore.Config
	Application           *application
	ApplicationAppUsage   *applicationAppUsage
	ApplicationAppVersion *applicationAppVersion
	ApplicationFeedback   *applicationFeedback
	ApplicationVisibility *applicationVisibility
}

// 资源服务定义
type application struct {
	service *ApplicationService
}
type applicationAppUsage struct {
	service *ApplicationService
}
type applicationAppVersion struct {
	service *ApplicationService
}
type applicationFeedback struct {
	service *ApplicationService
}
type applicationVisibility struct {
	service *ApplicationService
}

// 资源服务方法定义
func (a *application) Get(ctx context.Context, req *GetApplicationReq, options ...larkcore.RequestOptionFunc) (*GetApplicationResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/application/v6/applications/:app_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApplicationResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *application) Patch(ctx context.Context, req *PatchApplicationReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodPatch,
		"/open-apis/application/v6/applications/:app_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *application) Underauditlist(ctx context.Context, req *UnderauditlistApplicationReq, options ...larkcore.RequestOptionFunc) (*UnderauditlistApplicationResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/application/v6/applications/underauditlist", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnderauditlistApplicationResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *application) UnderauditlistByIterator(ctx context.Context, req *UnderauditlistApplicationReq, options ...larkcore.RequestOptionFunc) (*UnderauditlistApplicationIterator, error) {
	return &UnderauditlistApplicationIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.Underauditlist,
		options:  options,
		limit:    req.Limit}, nil
}
func (a *applicationAppUsage) Overview(ctx context.Context, req *OverviewApplicationAppUsageReq, options ...larkcore.RequestOptionFunc) (*OverviewApplicationAppUsageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodPost,
		"/open-apis/application/v6/applications/:app_id/app_usage/overview", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &OverviewApplicationAppUsageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *applicationAppVersion) Get(ctx context.Context, req *GetApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*GetApplicationAppVersionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/application/v6/applications/:app_id/app_versions/:version_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApplicationAppVersionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *applicationAppVersion) Patch(ctx context.Context, req *PatchApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationAppVersionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodPatch,
		"/open-apis/application/v6/applications/:app_id/app_versions/:version_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationAppVersionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *applicationFeedback) List(ctx context.Context, req *ListApplicationFeedbackReq, options ...larkcore.RequestOptionFunc) (*ListApplicationFeedbackResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/application/v6/applications/:app_id/feedbacks", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListApplicationFeedbackResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *applicationFeedback) Patch(ctx context.Context, req *PatchApplicationFeedbackReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationFeedbackResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, a.service.config, http.MethodPatch,
		"/open-apis/application/v6/applications/:app_id/feedbacks/:feedback_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationFeedbackResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
