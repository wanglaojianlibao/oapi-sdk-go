// Package sheets code generated by oapi sdk gen
package larksheets

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go.v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *SheetsService {
	s := &SheetsService{config: config}
	s.Spreadsheet = &spreadsheet{service: s}
	s.SpreadsheetSheet = &spreadsheetSheet{service: s}
	s.SpreadsheetSheetFilter = &spreadsheetSheetFilter{service: s}
	s.SpreadsheetSheetFilterView = &spreadsheetSheetFilterView{service: s}
	s.SpreadsheetSheetFilterViewCondition = &spreadsheetSheetFilterViewCondition{service: s}
	s.SpreadsheetSheetFloatImage = &spreadsheetSheetFloatImage{service: s}
	return s
}

// 业务域服务定义
type SheetsService struct {
	config                              *larkcore.Config
	Spreadsheet                         *spreadsheet
	SpreadsheetSheet                    *spreadsheetSheet
	SpreadsheetSheetFilter              *spreadsheetSheetFilter
	SpreadsheetSheetFilterView          *spreadsheetSheetFilterView
	SpreadsheetSheetFilterViewCondition *spreadsheetSheetFilterViewCondition
	SpreadsheetSheetFloatImage          *spreadsheetSheetFloatImage
}

// 资源服务定义
type spreadsheet struct {
	service *SheetsService
}
type spreadsheetSheet struct {
	service *SheetsService
}
type spreadsheetSheetFilter struct {
	service *SheetsService
}
type spreadsheetSheetFilterView struct {
	service *SheetsService
}
type spreadsheetSheetFilterViewCondition struct {
	service *SheetsService
}
type spreadsheetSheetFloatImage struct {
	service *SheetsService
}

// 资源服务方法定义
func (s *spreadsheet) Create(ctx context.Context, req *CreateSpreadsheetReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheet) Find(ctx context.Context, req *FindSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*FindSpreadsheetSheetResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &FindSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheet) MoveDimension(ctx context.Context, req *MoveDimensionSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*MoveDimensionSpreadsheetSheetResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveDimensionSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheet) Replace(ctx context.Context, req *ReplaceSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*ReplaceSpreadsheetSheetResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ReplaceSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Get(ctx context.Context, req *GetSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPut,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterView) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterView) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterView) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterView) Patch(ctx context.Context, req *PatchSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*PatchSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPatch,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterView) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterViewCondition) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterViewCondition) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterViewCondition) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterViewCondition) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilterViewCondition) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPut,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFloatImage) Create(ctx context.Context, req *CreateSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFloatImage) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFloatImage) Get(ctx context.Context, req *GetSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFloatImage) Patch(ctx context.Context, req *PatchSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*PatchSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodPatch,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFloatImage) Query(ctx context.Context, req *QuerySpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
