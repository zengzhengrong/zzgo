package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zdup"
)

func TestRemoveDuplicateElement(t *testing.T) {
	result := zdup.RemoveDuplicateString([]string{"1", "2", "3", "4", "5", "6", "7", "8", "8", "1"})
	fmt.Println(result)
}
