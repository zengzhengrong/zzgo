package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zengzhengrong/zzgo/tag"
	"github.com/zengzhengrong/zzgo/zdup"
)

func TestConenct(t *testing.T) {
	opt, err := redis.ParseURL("redis://192.168.2.100:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	if err := tag.Init(client); err != nil {
		panic(err)
	}
}

func TestTagCount(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bigtg.Count(nil))

}

func TestLen(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bigtg.Len())
}

func TestGetFullName(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tg.GetFullName())

}

func TestAll(t *testing.T) {
	TestConenct(t)
	bigtg, err := tag.NewTag("man", nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {

	}
	rr := bigtg.All(time.Second * 20)
	fmt.Println(rr)
	fmt.Println(zdup.DuplicateInt64(rr))
}

func TestAND(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg2, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgAND := tag.AND(false, tg, tg2)
	fmt.Println(tgAND.All())
}

func TestOR(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg2, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgOR := tag.OR(false, tg, tg2)
	fmt.Println(tgOR.All())
}

func TestXOR(t *testing.T) {
	TestConenct(t)
	tg, err := tag.NewTag("not_man", nil)
	tg3, err := tag.NewTag("woman", nil)
	if err != nil {
		panic(err)
	}
	tgXOR := tag.XOR(false, tg3, tg)
	fmt.Println(tgXOR.All())
}
