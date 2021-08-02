package ztime

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//判断时间是当年的第几周
func WeekByDate(t time.Time) (int, int) {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return t.Year(), week
}

type WeekDate struct {
	StartTime time.Time
	EndTime   time.Time
}

// 将开始时间和结束时间分割为周为单位
func GetWeeks(startTime, endTime time.Time) []WeekDate {
	weekDate := make([]WeekDate, 0)
	diffDuration := endTime.Sub(startTime)
	days := int(math.Ceil(float64(diffDuration/(time.Hour*24)))) + 1

	currentWeekDate := WeekDate{}

	currentWeekDate.EndTime = endTime
	currentWeekDay := int(endTime.Weekday())
	if currentWeekDay == 0 {
		currentWeekDay = 7
	}
	currentWeekDate.StartTime = endTime.AddDate(0, 0, -currentWeekDay+1)
	nextWeekEndTime := currentWeekDate.StartTime
	weekDate = append(weekDate, currentWeekDate)

	for i := 0; i < (days-currentWeekDay)/7; i++ {
		weekData := WeekDate{}
		weekData.EndTime = nextWeekEndTime
		weekData.StartTime = nextWeekEndTime.AddDate(0, 0, -7)

		nextWeekEndTime = weekData.StartTime
		weekDate = append(weekDate, weekData)
	}

	if lastDays := (days - currentWeekDay) % 7; lastDays > 0 {
		lastData := WeekDate{}
		lastData.EndTime = nextWeekEndTime
		lastData.StartTime = nextWeekEndTime.AddDate(0, 0, -lastDays)

		weekDate = append(weekDate, lastData)
	}

	return weekDate
}

// 获取时间戳   return:1464059756
func GetTimeUnix(t time.Time) int64 {
	return t.Local().Unix()
}

// 获取当前时间的时间戳 return:1464059756
func GetNowTimeUnix() int64 {
	return GetTimeUnix(time.Now())
}

// 获取当日晚上24点（次日0点）的时间
func Get24Time(t time.Time) time.Time {
	dateStr := TimeToDate(t.Add(time.Hour * 24))
	return DateStrToTime(dateStr)
}

// 获取当日晚上24点（次日0点）的时间戳
func Get24TimeUnix(t time.Time) int64 {
	t24 := Get24Time(t)
	return GetTimeUnix(t24)
}

// 获取今天晚上24点（次日0点）的时间
func GetToday24Time() time.Time {
	return Get24Time(time.Now())
}

// 获取今天晚上24点（次日0点）的时间戳
func GetToday24TimeUnix() int64 {
	return Get24TimeUnix(time.Now())
}

// 时间转换成日期字符串 (time.Time to "2006-01-02")
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// 获取当前时间的日期字符串（"2006-01-02"）
func GetNowDateStr() string {
	return TimeToDate(time.Now())
}

// 时间转换成日期+时间字符串 (time.Time to "2006-01-02 15:04:05")
func TimeToDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 获取当前的时期+时间字符串
func GetNowStr() string {
	return TimeToDateTime(time.Now())
}

// 日期字符串转换成时间 ("2006-01-02" to time.Time)
func DateStrToTime(d string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", d, time.Local)
	return t
}

// 日期+时间字符串转换成时间 ("2006-01-02 15:04:05" to time.Time)
func DateTimeStrToTime(dt string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dt, time.Local)
	return t
}

// 时间字符串转换成时间 ("15:04:05" to time.Time)
func TimeTodayStrToTime(dt string) time.Time {
	now := time.Now()
	strNowDate := TimeToDate(now)
	return DateTimeStrToTime(strNowDate + " " + dt)
}

// 是否是周末
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return true
	}
	return false
}

// 时间戳传时间字符串
func TimeUnixsToTimeStrs(unixs ...int64) []string {
	timestrs := make([]string, len(unixs))
	for _, unix := range unixs {
		t := time.Unix(unix, 0)
		timestrs = append(timestrs, TimeToDateTime(t))
	}
	return timestrs
}

// 时分秒字符串转时间戳，传入示例：8:40 or 8:40:10
func HmsToUnix(str string) (int64, error) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Now()
	arr := strings.Split(str, ":")
	if len(arr) < 2 {
		return 0, errors.New("time format error")
	}
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	s := 0
	if len(arr) == 3 {
		s, _ = strconv.Atoi(arr[3])
	}
	formatted1 := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), h, m, s)
	res, err := time.ParseInLocation("20060102150405", formatted1, l)
	if err != nil {
		return 0, err
	} else {
		return res.Unix(), nil
	}
}

// 获取一个当前时间和通用单位的时间间隔 返回开始时间和截至时间的时间戳
// unit : years,months,days,hours
// amount : the number of unit
func GetTimeInterval(unit string, amount int) (startTime, endTime int64) {
	t := time.Now()
	nowTime := t.Unix()
	tmpTime := int64(0)
	switch unit {
	case "years":
		tmpTime = time.Date(t.Year()+amount, t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local).Unix()
	case "months":
		tmpTime = time.Date(t.Year(), t.Month()+time.Month(amount), t.Day(), t.Hour(), 0, 0, 0, time.Local).Unix()
	case "days":
		tmpTime = time.Date(t.Year(), t.Month(), t.Day()+amount, t.Hour(), 0, 0, 0, time.Local).Unix()
	case "hours":
		tmpTime = time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+amount, 0, 0, 0, time.Local).Unix()
	}
	if amount > 0 {
		startTime = nowTime
		endTime = tmpTime
	} else {
		startTime = tmpTime
		endTime = nowTime
	}
	return
}
