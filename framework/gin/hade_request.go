// Copyright 2024 Caohaitao.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)

// const defaultMultipartMemory = 32 << 20 // 32 MB

// 代表请求包含的方法
type IRequest interface {
	// 请求地址url中带的参数
	// 形如：foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)
	DefaultQuery(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultForm(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// // #endregion

// 请求地址url中带的参数
// 形如: foo.com?a=1&b=bar&c[]=bar

// 获取请求地址中所有参数
func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return map[string][]string(ctx.queryCache)
}

// 获取Int类型的请求参数
func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals, true
		}
	}
	return def, false
}

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.HadeParam(key); val != nil {
		// 通过cast进行类型转换
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := ctx.HadeParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// 获取路由参数
func (ctx *Context) HadeParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return map[string][]string(ctx.formCache)
}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFile(key string) (*multipart.FileHeader, error) {
	if ctx.Request.MultipartForm == nil {
		if err := ctx.Request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.Request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// 将body文本解析到obj结构体中
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request != nil {
		// 读取文本
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		// 重新填充request.Body，为后续的逻辑二次读取做准备
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.Request != nil {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// 其他格式
//func (ctx *Context) GetRawData() ([]byte, error) {
//	if ctx.request != nil {
//		body, err := ioutil.ReadAll(ctx.request.Body)
//		if err != nil {
//			return nil, err
//		}
//		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
//		return body, nil
//	}
//	return nil, errors.New("ctx.request empty")
//}

// 基础信息
func (ctx *Context) Uri() string {
	return ctx.Request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.Request.Method
}

func (ctx *Context) Host() string {
	return ctx.Request.URL.Host
}

func (ctx *Context) ClientIp() string {
	r := ctx.Request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// header
func (ctx *Context) Headers() map[string][]string {
	return map[string][]string(ctx.Request.Header)
}

//func (ctx *Context) Header(key string) (string, bool) {
//	vals := ctx.request.Header.Values(key)
//	if vals == nil || len(vals) <= 0 {
//		return "", false
//	}
//	return vals[0], true
//}

// cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.Request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

//func (ctx *Context) Cookie(key string) (string, bool) {
//	panic("not implemented") // TODO: Implement
//}
