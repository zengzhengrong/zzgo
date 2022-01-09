package request

import (
	"net/http"
	"time"
)

const (
	defaultMultipartMemory             = 32 << 20 // 32 MB
	defaultContentType                 = "application/json"
	jsonContectType                    = defaultContentType
	formContectType                    = "application/x-www-form-urlencoded"
	defaultDebug                       = false
	defaultTimeout                     = 60 * time.Second
	defaultTLSConfigInsecureSkipVerify = true
)

var MaxUploadThreads int = 20

var defaultCheckRedirect = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

type Options interface {
	*ClientOptions | *ReqOptions
}

type GenericOption[T Options] interface {
	apply(T)
}
