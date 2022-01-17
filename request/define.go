package request

import (
	"net/http"
	"os"
	"time"

	"github.com/zengzhengrong/zzgo/zbool"
)

const (
	defaultMultipartMemory             = 32 << 20 // 32 MB
	defaultContentType                 = "application/json"
	JsonContectType                    = defaultContentType
	FormContectType                    = "application/x-www-form-urlencoded"
	defaultTimeout                     = 60 * time.Second
	defaultTLSConfigInsecureSkipVerify = true
)

func setDefaultDebug() bool {
	debug := os.Getenv("REQUEST_DEBUG")
	return zbool.BoolFlagMap.Check(debug)
}

var MaxUploadThreads int = 20
var DefaultDebug = setDefaultDebug
var defaultCheckRedirect = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

type Options interface {
	*ClientOptions | *ReqOptions
}

type GenericOption[T Options] interface {
	apply(T)
}
