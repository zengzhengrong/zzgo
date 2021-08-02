package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zsystem"
)

// CurPath 获取当前运行目录
func TestCurPath(t *testing.T) {
	fmt.Println(zsystem.CurPath())
}

// Sep 获取系统分隔符
func TestSep(t *testing.T) {

	fmt.Println(zsystem.Sep())
}

// IsDirExists 判断目录是否存在
func TestIsDirExists(t *testing.T) {
	fmt.Println(zsystem.IsDirExists("../test"))
}

// IsFileExists 判断文件是否存在
func TestIsFileExists(t *testing.T) {
	fmt.Println(zsystem.IsFileExists("./system.go"))
}
