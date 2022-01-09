package request

import (
	"crypto/tls"
	"net/http"
	"time"
)

type DebugOption bool
type TimeoutOption time.Duration
type TransportOption struct{ http.RoundTripper }
type CheckRedirectOption func(req *http.Request, via []*http.Request) error
type TLSClientConfigOption struct{ tls.Config }
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

func NewClient[T GenericOption[*ClientOptions]](opts ...T) *Client {
	options := &ClientOptions{
		Debug:         defaultDebug,
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
