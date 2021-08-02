package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zstr"
)

func TestStr2Bytes(t *testing.T) {
	b := zstr.Str2Bytes("zzr")
	fmt.Println(b)
}

//TestBytes2Str is bytes转字符串不需要通过拷贝
func TestBytes2Str(t *testing.T) {
	s := zstr.Bytes2Str([]byte{122, 122, 114})
	fmt.Println(s)
}
