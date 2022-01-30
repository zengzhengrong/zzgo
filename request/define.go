package request

import (
	"net/http"
	"os"
	"time"

	"github.com/zengzhengrong/zzgo/zbool"
)

const (
	DefaultMultipartMemory             = 32 << 20 // 32 MB
	DefaultContentType                 = "application/json"
	JsonContectType                    = DefaultContentType
	FormContectType                    = "application/x-www-form-urlencoded"
	DefaultTimeout                     = 60 * time.Second
	DefaultTLSConfigInsecureSkipVerify = true
	PiplineCtxValueKey                 = "values"
)

func SetDefaultDebug() bool {
	debug := os.Getenv("REQUEST_DEBUG")
	return zbool.BoolFlagMap.Check(debug)
}

var MaxUploadThreads int = 20
var DefaultDebug = SetDefaultDebug
var DefaultCheckRedirect = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
