package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zshuffle"
)

func TestShuffle(t *testing.T) {
	els := []string{"1", "2", "311", "5"}
	zshuffle.Shuffle(els)
	fmt.Println(els)
}
