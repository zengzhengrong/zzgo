package logz

import (
	"os"
	"runtime/debug"

	"github.com/zengzhengrong/zzgo/zbool"

	"github.com/sirupsen/logrus"
)

var stackOpen = os.Getenv("STACK_OPEN")

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	if os.Getenv("LOG_FORMAT") == "json" {
		log.Formatter = &logrus.JSONFormatter{}
	}
	if log == nil {
		log = logrus.New()
	}
}

// Info is 打印 info 级别的栈 不返回error
func Info(info interface{}) {
	if zbool.BoolFlagMap.Check(stackOpen) {
		logrus.Infof("%v:%s", info, string(debug.Stack()))
	} else {
		logrus.Infof("%v", info)
	}
}

// Err is 打印 err级别的栈 并返回error
func Err(err error) error {
	if zbool.BoolFlagMap.Check(stackOpen) {
		logrus.Errorf("%v:%s", err, string(debug.Stack()))
	}
	return err
}

// Warn is 打印 war 级别的栈 不返回error
func Warn(war interface{}) {
	if zbool.BoolFlagMap.Check(stackOpen) {
		logrus.Warnf("%v:%s", war, string(debug.Stack()))
	} else {
		logrus.Warnf("%v", war)
	}
}
