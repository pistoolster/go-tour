package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

const FORMAT string = "2006-01-02 15:04:05-07:00"

func FixReadTimeFromMysql(t time.Time, tz string) (fixed time.Time, err error) {
	// 传入参数是rows.Scan后的Time对象, 在go中sql.Open时连接字符串里loc选项为空的情况下, 会被读取成utc, location为time.UTC
	if t.Location() != time.UTC {
		// 兼容其他情况
		_, offset := t.Zone()
		t = t.Add(time.Second * time.Duration(offset)).UTC()
	}
	// 因为Format不识别+8:00/8:00, 须将+8:00/8:00修复为+08:00  +08:00/8:00/+8:00 ==> +08:00
	switch len(tz) {
	case 6:
		// +08:00
	case 5:
		// +8:00 08:00
		if strings.HasPrefix(tz, "+") || strings.HasPrefix(tz, "-") {
			tz = fmt.Sprintf("%s0%s", string(tz[0]), tz[1:])
		} else {
			tz = fmt.Sprintf("%s0%s", "+", tz)
		}
	case 4:
		// 8:00
		tz = fmt.Sprintf("%s0%s", "+", tz)
	case 0:
		tz = "+00:00"
	default:
		return t, errors.New("invalid param tz")
	}
	fixed, err = time.Parse(FORMAT, t.Format("2006-01-02 15:04:05")+tz)
	if err != nil {
		return t, errors.New("time.Parse error, t:%s, tz:%s")
	}
	return fixed, nil
}

func TimeInLocation(t time.Time, tz string) (fixed time.Time, err error) {
	fixed, err = FixReadTimeFromMysql(t, tz)
	if err != nil {
		return
	}
	fixed = t.In(fixed.Location())
	return
}

func main() {
	timeStr := "2019-10-25 04:40:15"
	timeZone := "+08:00"
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(t.Location())
	// fmt.Println(t.Zone())
	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	// // fmt.Println(t.In(loc).Format("2006-01-02 15:04:05-07:00"))
	// // fmt.Println(t.Local().Format("2006-01-02 15:04:05-07:00"))
	// // fmt.Println(t.Format(FORMAT))
	// // fmt.Println(t.In(loc).Format(FORMAT))
	// fix, err := FixReadTimeFromMysql(t, timeZone)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(fix.Local())
	// fmt.Println(fix.Format(FORMAT))
	// fmt.Println(fix.In(loc).Format(FORMAT))
	fix, err := TimeInLocation(t, timeZone)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(fix.Location())
	fmt.Println(fix.Format(FORMAT))
	fmt.Println(fix.In(loc).Format(FORMAT))
	fmt.Println(time.Now().UTC())
}
