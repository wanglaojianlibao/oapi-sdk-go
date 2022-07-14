// Package translation code generated by oapi sdk gen
package larktranslation

import (
	"github.com/larksuite/oapi-sdk-go.v3/core"
)

// 生成枚举值

// 生成数据类型

type Term struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

// builder开始
type TermBuilder struct {
	from     string
	fromFlag bool
	to       string
	toFlag   bool
}

func NewTermBuilder() *TermBuilder {
	builder := &TermBuilder{}
	return builder
}

func (builder *TermBuilder) From(from string) *TermBuilder {
	builder.from = from
	builder.fromFlag = true
	return builder
}
func (builder *TermBuilder) To(to string) *TermBuilder {
	builder.to = to
	builder.toFlag = true
	return builder
}

func (builder *TermBuilder) Build() *Term {
	req := &Term{}
	if builder.fromFlag {
		req.From = &builder.from

	}
	if builder.toFlag {
		req.To = &builder.to

	}
	return req
}

// builder结束

type Text struct {
}

// builder开始
// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

type DetectTextReqBodyBuilder struct {
	text     string
	textFlag bool
}

// 生成body的New构造器
func NewDetectTextReqBodyBuilder() *DetectTextReqBodyBuilder {
	builder := &DetectTextReqBodyBuilder{}
	return builder
}

// 1.2 生成body的builder属性方法
func (builder *DetectTextReqBodyBuilder) Text(text string) *DetectTextReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

// 1.3 生成body的build方法
func (builder *DetectTextReqBodyBuilder) Build() *DetectTextReqBody {
	req := &DetectTextReqBody{}
	if builder.textFlag {
		req.Text = &builder.text
	}
	return req
}

// 上传文件path开始
type DetectTextPathReqBodyBuilder struct {
	text     string
	textFlag bool
}

func NewDetectTextPathReqBodyBuilder() *DetectTextPathReqBodyBuilder {
	builder := &DetectTextPathReqBodyBuilder{}
	return builder
}
func (builder *DetectTextPathReqBodyBuilder) Text(text string) *DetectTextPathReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

func (builder *DetectTextPathReqBodyBuilder) Build() (*DetectTextReqBody, error) {
	req := &DetectTextReqBody{}
	if builder.textFlag {
		req.Text = &builder.text
	}
	return req, nil
}

// 上传文件path结束

// 1.4 生成请求的builder结构体
type DetectTextReqBuilder struct {
	body     *DetectTextReqBody
	bodyFlag bool
}

// 生成请求的New构造器
func NewDetectTextReqBuilder() *DetectTextReqBuilder {
	builder := &DetectTextReqBuilder{}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *DetectTextReqBuilder) Body(body *DetectTextReqBody) *DetectTextReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *DetectTextReqBuilder) Build() *DetectTextReq {
	req := &DetectTextReq{}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type DetectTextReqBody struct {
	Text *string `json:"text,omitempty"`
}

type DetectTextReq struct {
	Body *DetectTextReqBody `body:""`
}

type DetectTextRespData struct {
	Language *string `json:"language,omitempty"`
}

type DetectTextResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *DetectTextRespData `json:"data"`
}

func (resp *DetectTextResp) Success() bool {
	return resp.Code == 0
}

type TranslateTextReqBodyBuilder struct {
	sourceLanguage     string
	sourceLanguageFlag bool
	text               string
	textFlag           bool
	targetLanguage     string
	targetLanguageFlag bool
	glossary           []*Term
	glossaryFlag       bool
}

// 生成body的New构造器
func NewTranslateTextReqBodyBuilder() *TranslateTextReqBodyBuilder {
	builder := &TranslateTextReqBodyBuilder{}
	return builder
}

// 1.2 生成body的builder属性方法
func (builder *TranslateTextReqBodyBuilder) SourceLanguage(sourceLanguage string) *TranslateTextReqBodyBuilder {
	builder.sourceLanguage = sourceLanguage
	builder.sourceLanguageFlag = true
	return builder
}
func (builder *TranslateTextReqBodyBuilder) Text(text string) *TranslateTextReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}
func (builder *TranslateTextReqBodyBuilder) TargetLanguage(targetLanguage string) *TranslateTextReqBodyBuilder {
	builder.targetLanguage = targetLanguage
	builder.targetLanguageFlag = true
	return builder
}
func (builder *TranslateTextReqBodyBuilder) Glossary(glossary []*Term) *TranslateTextReqBodyBuilder {
	builder.glossary = glossary
	builder.glossaryFlag = true
	return builder
}

// 1.3 生成body的build方法
func (builder *TranslateTextReqBodyBuilder) Build() *TranslateTextReqBody {
	req := &TranslateTextReqBody{}
	if builder.sourceLanguageFlag {
		req.SourceLanguage = &builder.sourceLanguage
	}
	if builder.textFlag {
		req.Text = &builder.text
	}
	if builder.targetLanguageFlag {
		req.TargetLanguage = &builder.targetLanguage
	}
	if builder.glossaryFlag {
		req.Glossary = builder.glossary
	}
	return req
}

// 上传文件path开始
type TranslateTextPathReqBodyBuilder struct {
	sourceLanguage     string
	sourceLanguageFlag bool
	text               string
	textFlag           bool
	targetLanguage     string
	targetLanguageFlag bool
	glossary           []*Term
	glossaryFlag       bool
}

func NewTranslateTextPathReqBodyBuilder() *TranslateTextPathReqBodyBuilder {
	builder := &TranslateTextPathReqBodyBuilder{}
	return builder
}
func (builder *TranslateTextPathReqBodyBuilder) SourceLanguage(sourceLanguage string) *TranslateTextPathReqBodyBuilder {
	builder.sourceLanguage = sourceLanguage
	builder.sourceLanguageFlag = true
	return builder
}
func (builder *TranslateTextPathReqBodyBuilder) Text(text string) *TranslateTextPathReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}
func (builder *TranslateTextPathReqBodyBuilder) TargetLanguage(targetLanguage string) *TranslateTextPathReqBodyBuilder {
	builder.targetLanguage = targetLanguage
	builder.targetLanguageFlag = true
	return builder
}
func (builder *TranslateTextPathReqBodyBuilder) Glossary(glossary []*Term) *TranslateTextPathReqBodyBuilder {
	builder.glossary = glossary
	builder.glossaryFlag = true
	return builder
}

func (builder *TranslateTextPathReqBodyBuilder) Build() (*TranslateTextReqBody, error) {
	req := &TranslateTextReqBody{}
	if builder.sourceLanguageFlag {
		req.SourceLanguage = &builder.sourceLanguage
	}
	if builder.textFlag {
		req.Text = &builder.text
	}
	if builder.targetLanguageFlag {
		req.TargetLanguage = &builder.targetLanguage
	}
	if builder.glossaryFlag {
		req.Glossary = builder.glossary
	}
	return req, nil
}

// 上传文件path结束

// 1.4 生成请求的builder结构体
type TranslateTextReqBuilder struct {
	body     *TranslateTextReqBody
	bodyFlag bool
}

// 生成请求的New构造器
func NewTranslateTextReqBuilder() *TranslateTextReqBuilder {
	builder := &TranslateTextReqBuilder{}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *TranslateTextReqBuilder) Body(body *TranslateTextReqBody) *TranslateTextReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *TranslateTextReqBuilder) Build() *TranslateTextReq {
	req := &TranslateTextReq{}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type TranslateTextReqBody struct {
	SourceLanguage *string `json:"source_language,omitempty"`
	Text           *string `json:"text,omitempty"`
	TargetLanguage *string `json:"target_language,omitempty"`
	Glossary       []*Term `json:"glossary,omitempty"`
}

type TranslateTextReq struct {
	Body *TranslateTextReqBody `body:""`
}

type TranslateTextRespData struct {
	Text *string `json:"text,omitempty"`
}

type TranslateTextResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *TranslateTextRespData `json:"data"`
}

func (resp *TranslateTextResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体
