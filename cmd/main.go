package main

import (
	"fmt"
	"github.com/hou-rong/dfss-reservation/pkg/api"
	"time"
)

func main() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if now.Hour() < 7 || now.Hour() > 21 {
		fmt.Println(now, "当前时间非驾校约车时间，等待中")
	} else {
		// 7 天后日期
		sevenDaysLater := today.AddDate(0, 0, 6)
		dating := sevenDaysLater.Format("2006-01-02")
		fmt.Println(now, "时间 OK，开始预约 7 天后的车，日期: ", dating)
		timestamp, checksum, token := api.Login()
		api.Reservation(timestamp, checksum, token, dating, "01", "02")
	}
	//if today.Weekday() == 0 || now.Weekday() == 6 {
	//	// 如果是周日或者周六：周日预定下周六的车，周六预约下周日的车
	//
	//	// 7 天后日期
	//	sevenDaysLater := today.AddDate(0, 0, 6)
	//	dating := sevenDaysLater.Format("20060102")
	//	fmt.Println(now, "时间 OK，开始预约 7 天后的车，日期: ", dating)
	//	timestamp, checksum, token := api.Login()
	//	api.Reservation(timestamp, checksum, token, dating)
	//} else {
	//	fmt.Println(now, "时间不 OK，等周六日再干吧")
	//}
}
