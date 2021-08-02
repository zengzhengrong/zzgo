package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zstr"
)

func TestHideString(t *testing.T) {
	result := zstr.HideString("11111111111111", 5)
	fmt.Println(result)
}
