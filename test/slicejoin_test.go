package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zengzhengrong/zzgo/zslice"
)

func TestSliceJoin(t *testing.T) {
	s := zslice.SliceJoin(",", struct{}{}, 2, 3, 4, 5)
	fmt.Println(s)
	sp := strings.Split(s, ",")
	fmt.Println(sp)
}
