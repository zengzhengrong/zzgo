package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zengzhengrong/zzgo/zslice"
	"github.com/zengzhengrong/zzgo/ztime"
)

func TestWeekByDate(t *testing.T) {

	fmt.Println(ztime.WeekByDate(time.Now()))
}

func TestGet(t *testing.T) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation("2006-01-02", "2021-07-01", l)
	endTime, _ := time.ParseInLocation("2006-01-02", "2021-07-25", l)
	weeks := ztime.GetWeeks(startTime, endTime)
	for _, item := range weeks {
		fmt.Println(item)
	}

	args, err := zslice.ToSliceInterface(weeks)
	if err != nil {
		panic(err)
	}

	list := zslice.SliceRever(args)

	for _, item := range list {
		fmt.Println(item)
	}
}

func TestHmsToUnix(t *testing.T) {
	uninx, err := ztime.HmsToUnix("16:10")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uninx)
}

func TestGetTimeIntervalToTimeStr(t *testing.T) {
	start, end := ztime.GetTimeInterval("days", 1)
	times := ztime.TimeUnixsToTimeStrs(start, end)
	fmt.Println(times)
}
