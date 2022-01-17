package request

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type DebugOption bool
type TimeoutOption time.Duration
type TransportOption struct{ http.RoundTripper }
type CheckRedirectOption func(req *http.Request, via []*http.Request) error
type TLSClientConfigOption struct{ tls.Config }
type Default struct{}
type Client struct {
	Opts       *ClientOptions
	HttpClient *http.Client
}
type ClientOptions struct {
	Debug           bool
	Timeout         time.Duration
	Transport       http.RoundTripper
	CheckRedirect   func(req *http.Request, via []*http.Request) error
	TLSClientConfig TLSClientConfigOption
}

func (t TimeoutOption) apply(opts *ClientOptions) {
	opts.Timeout = time.Duration(t)
}

func (d DebugOption) apply(opts *ClientOptions) {
	opts.Debug = bool(d)
}

func (t TransportOption) apply(opts *ClientOptions) {
	opts.Transport = TransportOption{t}
}

func (c CheckRedirectOption) apply(opts *ClientOptions) {
	opts.CheckRedirect = CheckRedirectOption(c)
}

func (t TLSClientConfigOption) apply(opts *ClientOptions) {
	opts.TLSClientConfig = TLSClientConfigOption(t)
	opts.Transport = &http.Transport{TLSClientConfig: &opts.TLSClientConfig.Config}
}

func (d Default) apply(opts *ClientOptions) {
	// no processing
}

func WithTransport(t http.RoundTripper) GenericOption[*ClientOptions] {
	return TransportOption{t}
}

// WithInsecureSkipVerify is will override Transport
func WithInsecureSkipVerify() GenericOption[*ClientOptions] {
	return &TLSClientConfigOption{tls.Config{InsecureSkipVerify: true}}
}

func WithDebug(debug bool) GenericOption[*ClientOptions] {
	return DebugOption(debug)
}
func WithTimeOut(timeout time.Duration) GenericOption[*ClientOptions] {
	return TimeoutOption(timeout)
}

func WithCheckRedirect(f func(req *http.Request, via []*http.Request) error) GenericOption[*ClientOptions] {
	return CheckRedirectOption(f)
}

func WithDefault() GenericOption[*ClientOptions] {
	return Default{}
}

func NewClient[T GenericOption[*ClientOptions]](opts ...T) *Client {
	options := &ClientOptions{
		Debug:         setDefaultDebug(),
		Timeout:       defaultTimeout,
		Transport:     http.DefaultTransport,
		CheckRedirect: defaultCheckRedirect,
	}
	for _, o := range opts {
		o.apply(options)
	}

	client := &http.Client{
		Transport:     options.Transport,
		CheckRedirect: options.CheckRedirect, // 获取301重定向
		Timeout:       options.Timeout,
	}
	return &Client{
		Opts:       options,
		HttpClient: client,
	}
}

// Do is ShortCut http client do method
func (client *Client) Do(r *Request) (*http.Response, error) {
	resp, err := client.HttpClient.Do(r.HttpReq)
	return resp, err
}

func getqueryheader(args ...map[string]string) (map[string]string, map[string]string) {
	var (
		query  map[string]string
		header map[string]string
	)

	if len(args) > 0 {
		query = args[0]

	}

	if len(args) == 2 {
		header = args[1]
	}
	return query, header
}

// GET is ShortCut get http method but not reuse tpc connect
// The first args[0] is query , args[1] is header
func GET(url string, args ...map[string]string) Response {
	query, header := getqueryheader(args...)
	r, err := NewReuqest(
		http.MethodGet,
		url,
		WithQuery(query),
		WithHeader(header),
	)
	if err != nil {
		return Response{nil, nil, err}
	}

	client := NewClient(WithDefault())
	resp, err := client.Do(r)
	if err != nil {
		return Response{resp, nil, err}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{resp, nil, err}
	}
	resp.Body.Close()
	return Response{resp, body, nil}

}

// GETBind is bind struct with Get method
func GETBind(v any, url string, args ...map[string]string) error {
	resp := GET(url, args...)
	if !resp.OK() && resp.GetError() != nil {
		return resp.GetError()
	}
	if err := resp.GetStruct(&v); err != nil {
		return err
	}
	return nil

}

// POST is shortcut post method with json
func POST(url string, postbody any, args ...map[string]string) Response {
	query, header := getqueryheader(args...)
	r, err := NewReuqest(
		http.MethodPost,
		url,
		WithBody(postbody),
		WithQuery(query),
		WithHeader(header),
	)
	if err != nil {
		return Response{nil, nil, err}
	}
	client := NewClient(WithDefault())
	resp, err := client.Do(r)
	if err != nil {
		return Response{resp, nil, err}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{resp, nil, err}
	}
	resp.Body.Close()
	return Response{resp, body, nil}

}

func POSTForm(url string, postbody any, args ...map[string]string) Response {
	query, header := getqueryheader(args...)
	r, err := NewReuqest(
		http.MethodPost,
		url,
		WithBody(postbody),
		WithQuery(query),
		WithHeader(header),
		WithContentType(FormContectType),
	)
	if err != nil {
		return Response{nil, nil, err}
	}
	client := NewClient(WithDefault())
	resp, err := client.Do(r)
	if err != nil {
		return Response{resp, nil, err}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{resp, nil, err}
	}
	resp.Body.Close()
	return Response{resp, body, nil}

}

// POSTBinaryBody is binary body upload
func POSTBinaryBody(url string, binfile io.Reader, timeout time.Duration, args ...map[string]string) Response {
	query, header := getqueryheader(args...)

	r, err := NewReuqest(
		http.MethodPost,
		url,
		WithContentType("binary/octet-stream"),
		WithQuery(query),
		WithHeader(header),
	)
	if err != nil {
		return Response{nil, nil, err}
	}
	client := NewClient(WithTimeOut(timeout))

	resp, err := client.Do(r)
	if err != nil {
		return Response{resp, nil, err}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{resp, nil, err}
	}
	resp.Body.Close()
	return Response{resp, body, nil}

}

// POSTMultiPartUpload is upload file , files key is fieldname of file ,file name is in fields key
func POSTMultiPartUpload(url string, files map[string]io.Reader, fields map[string]string, timeout time.Duration, args ...map[string]string) Response {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for name, file := range files {
		filename, ok := fields[name]
		if !ok {
			return Response{nil, nil, errors.New(name + "is not found in the fields")}
		}
		writer, err := writer.CreateFormFile(name, filename)
		if err != nil {
			return Response{nil, nil, err}
		}
		_, err = io.Copy(writer, file)
		if err != nil {
			return Response{nil, nil, err}
		}
		delete(fields, name)
	}
	for k, v := range fields {
		err := writer.WriteField(k, v)
		if err != nil {
			return Response{nil, nil, err}
		}
	}
	err := writer.Close()
	query, header := getqueryheader(args...)
	r, err := NewReuqest(
		http.MethodPost,
		url,
		WithContentType(writer.FormDataContentType()),
		WithQuery(query),
		WithHeader(header),
	)
	if err != nil {
		return Response{nil, nil, err}
	}
	client := NewClient(WithTimeOut(timeout))
	resp, err := client.Do(r)
	if err != nil {
		return Response{resp, nil, err}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{resp, nil, err}
	}
	resp.Body.Close()
	return Response{resp, body, nil}
}

func POSTBind(v any, url string, postbody any, args ...map[string]string) error {
	resp := POST(url, postbody, args...)
	if !resp.OK() && resp.GetError() != nil {
		return resp.GetError()
	}
	if err := resp.GetStruct(&v); err != nil {
		return err
	}
	return nil
}

func POSTFormBind(v any, url string, postbody any, args ...map[string]string) error {
	resp := POSTForm(url, postbody, args...)
	if !resp.OK() && resp.GetError() != nil {
		return resp.GetError()
	}
	if err := resp.GetStruct(&v); err != nil {
		return err
	}
	return nil
}
