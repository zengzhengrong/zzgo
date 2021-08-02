package curl

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/zengzhengrong/zzgo/zstr"
)

var Debug bool

// 只支持form 与json 提交, 请留意body的类型, 支持string, []byte, map[string]string map[string]interface{} map类型是form表单提交
func Get(router string, header map[string]string) ([]byte, error) {
	return curl(http.MethodGet, router, nil, header)
}

// 只支持form 与json 提交, 请留意body的类型, 支持string, []byte, map[string]string map[string]interface{} map类型是form表单提交
func Post(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPost, router, body, header)
}

// 只支持form 与json 提交, 请留意body的类型, 支持string, []byte, map[string]string  map[string]interface{} map类型是form表单提交
func Put(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPut, router, body, header)
}

// 只支持form 与json 提交, 请留意body的类型, 支持string, []byte, map[string]string  map[string]interface{} map类型是form表单提交
func Patch(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPatch, router, body, header)
}

// 只支持form 与json 提交, 请留意body的类型, 支持string, []byte, map[string]string  map[string]interface{} map类型是form表单提交
func Delete(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodDelete, router, body, header)
}

func curl(method, router string, body interface{}, header map[string]string) ([]byte, error) {
	var reqBody io.Reader
	contentType := "application/json"
	switch v := body.(type) {
	case string:
		reqBody = strings.NewReader(v)
	case []byte:
		reqBody = bytes.NewReader(v)
	case map[string]string:
		val := url.Values{}
		for k, v := range v {
			val.Set(k, v)
		}
		reqBody = strings.NewReader(val.Encode())
		contentType = "application/x-www-form-urlencoded"
	case map[string]interface{}:
		val := url.Values{}
		for k, v := range v {
			val.Set(k, v.(string))
		}
		reqBody = strings.NewReader(val.Encode())
		contentType = "application/x-www-form-urlencoded"
	}
	if header == nil {
		header = map[string]string{"Content-Type": contentType}
	}
	if _, ok := header["Content-Type"]; !ok {
		header["Content-Type"] = contentType
	}
	resp, er := Do(method, router, reqBody, header)
	if er != nil {
		return nil, er
	}
	res, err := ioutil.ReadAll(resp.Body)
	if Debug {
		blob := zstr.SerializeStr(body)
		if contentType != "application/json" {
			blob = HttpBuild(body)
		}
		fmt.Printf("\n\n=====================\n[url]: %s\n[time]: %s\n[method]: %s\n[content-type]: %v\n[req_header]: %s\n[req_body]: %#v\n[resp_err]: %v\n[resp_header]: %v\n[resp_body]: %v\n=====================\n\n",
			router,
			time.Now().Format("2006-01-02 15:04:05.000"),
			method,
			contentType,
			HttpBuildQuery(header),
			blob,
			err,
			zstr.SerializeStr(resp.Header),
			string(res),
		)
	}
	resp.Body.Close()
	return res, err
}

func Do(method, router string, reqBody io.Reader, header map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest(method, router, reqBody)

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		// 获取301重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client.Do(req)
}

// 组建get请求参数,sortAsc true为小到大,false为大到小,nil不排序  a=123&b=321
func HttpBuildQuery(args map[string]string, sortAsc ...bool) string {
	str := ""
	if len(args) == 0 {
		return str
	}
	if len(sortAsc) > 0 {
		keys := make([]string, 0, len(args))
		for k := range args {
			keys = append(keys, k)
		}
		if sortAsc[0] {
			sort.Strings(keys)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		}
		for _, k := range keys {
			str += "&" + k + "=" + args[k]
		}
	} else {
		for k, v := range args {
			str += "&" + k + "=" + v
		}
	}
	return str[1:]
}

func HttpBuild(body interface{}, sortAsc ...bool) string {
	params := map[string]string{}
	if args, ok := body.(map[string]interface{}); ok {
		for k, v := range args {
			params[k] = cast.ToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	if args, ok := body.(map[string]string); ok {
		for k, v := range args {
			params[k] = cast.ToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	if args, ok := body.(map[string]int); ok {
		for k, v := range args {
			params[k] = cast.ToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	return cast.ToString(body)
}
