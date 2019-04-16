package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func reservation(timestamp, checksum, authToken, dating, timeIdx, lessonId string) {
	json := []byte(fmt.Sprintf(`{"vehicleNum": "","datingCarDate": "%s","trainingTimeSlotId": "%s","lessonId": "%s"}`, dating, timeIdx, lessonId))
	body := bytes.NewBuffer(json)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.dfssclub.cn/api/v2/CarAppt/Dfss/AddPlan", body)
	req.Header.Add("apikey", "2bd45eff35cb4942babe02cbcb7a6c26")
	req.Header.Add("authtoken", authToken)
	req.Header.Add("apichecksum", checksum)
	req.Header.Add("timestamp", timestamp)
	req.Header.Add("clientversioncode", "78")
	req.Header.Add("appownerid", "1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("reservation timestamp: ", timestamp, "checksum: ", checksum, "authToken: ", authToken, "dating: ", dating)
	fmt.Println("reservation Status : ", resp.Status)
	fmt.Println("reservation Headers : ", resp.Header)
	fmt.Println("reservation Body : ", string(respBody))
}

func Reservation(timestamp, checksum, authToken, dating, timeIdx, lessonId string) {
	reservation(timestamp, checksum, authToken, dating, timeIdx, lessonId)
	time.Sleep(time.Second * 2)
}
