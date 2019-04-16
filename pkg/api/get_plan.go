package api

import (
	"encoding/json"
	"fmt"
	"github.com/hou-rong/dfss-reservation/pkg/config"
	"io/ioutil"
	"net/http"
)

type Plan struct {
	DatingCarDate               string `json:"datingCarDate"`
	LessonID                    string `json:"lessonID"`
	LessonName                  string `json:"lessonName"`
	TrainingTimeSlotId          string `json:"trainingTimeSlotId"`
	TrainingTimeSlotName        string `json:"trainingTimeSlotName"`
	DayOfWeek                   string `json:"dayOfWeek"`
	DisableGetVehicleByTimeSlot bool   `json:"disableGetVehicleByTimeSlot"`
}

type PlanData struct {
	Data struct {
		Plans  []Plan `json:"plans"`
		Status int64  `json:"status"`
	} `json:"data"`
}

func getPlan(timestamp, checksum, authToken, dating string) *PlanData {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.dfssclub.cn/api/v2/CarAppt/Dfss/AvailablePlans_New?fromDate=%s&toDate=%s&vehicleNum", dating, dating),
		nil,
	)

	req.Header.Add("apikey", "2bd45eff35cb4942babe02cbcb7a6c26")
	req.Header.Add("authtoken", authToken)
	req.Header.Add("apichecksum", checksum)
	req.Header.Add("timestamp", timestamp)
	req.Header.Add("clientversioncode", "78")
	req.Header.Add("appownerid", "1")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return nil
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	resData := PlanData{}
	err = json.Unmarshal(respBody, &resData)
	if err != nil {
		return nil
	}
	return &resData
}

func GetPlan(timestamp, checksum, authToken, dating string) *PlanData {
	for i := 0; i < 4; i++ {
		plan := getPlan(timestamp, checksum, authToken, dating)
		if plan != nil {
			return plan
		}
	}
	return nil
}

func GetBestPlan(timestamp, checksum, authToken, dating string) *Plan {
	planData := GetPlan(timestamp, checksum, authToken, dating)
	if planData == nil {
		return nil
	}
	var maxScore int64 = 0
	var maxScoreIdx = 0
	for idx, plan := range planData.Data.Plans {
		if score, ok := config.TrainingTimeSlotScore[plan.TrainingTimeSlotName]; ok && score > maxScore {
			maxScoreIdx = idx
		}
	}

	plan := planData.Data.Plans[maxScoreIdx]
	fmt.Printf("获取到最佳方案：%+v\n", plan)
	return &plan
}
