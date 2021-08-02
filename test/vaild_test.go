package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zvalid"
)

func TestValidatePhoneNumber(t *testing.T) {
	ok, err := zvalid.ValidatePhoneNumber("1328888118")
	if err != nil {
		fmt.Println("ko")
	}
	if ok {
		fmt.Println("ok")
	}

}
