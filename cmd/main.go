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

		// 先登录
		timestamp, checksum, token := api.Login()
		for i := 0; i < 9; i++ {
			// 获取最佳方案
			bestPlan := api.GetBestPlan(timestamp, checksum, token, dating)
			if bestPlan != nil {
				// 如果有最佳方案，则下订单
				api.Reservation(timestamp, checksum, token, dating, bestPlan.TrainingTimeSlotId, bestPlan.LessonID)
			} else {
				fmt.Println(now, "今天没的选了，改天吧", dating)
			}
		}
	}
}
