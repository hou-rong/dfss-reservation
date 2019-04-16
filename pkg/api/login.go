package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hou-rong/dfss-reservation/pkg/config"
	"github.com/hou-rong/dfss-reservation/pkg/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type LoginInfo struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

func login(timestamp, checksum string) *LoginInfo {
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{"mobileOrStudentNum": "%s","drivingSchoolId": "1","password": "%s"}`, config.USERNAME, config.PASSWORD)))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.dfssclub.cn/api/v2/User/MixedLogin", body)

	req.Header.Add("apikey", "2bd45eff35cb4942babe02cbcb7a6c26")
	req.Header.Add("timestamp", timestamp)
	req.Header.Add("appownerid", "1")
	req.Header.Add("clientversioncode", "78")
	req.Header.Add("apichecksum", checksum)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
		return nil
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	info := LoginInfo{}
	fmt.Println("Response", string(respBody))
	err = json.Unmarshal(respBody, &info)
	if err != nil {
		fmt.Println("Failure : ", err)
		return nil
	}
	return &info
}

func Login() (string, string, string) {
	now := time.Now().Unix()
	timestamp := strconv.FormatInt(now, 10)
	fmt.Println("Start Login Timestamp", timestamp, "Checksum", "")
	info := login(timestamp, "")
	message := info.Message
	fmt.Println(message)
	checksum := util.GetCheckSum(message)
	fmt.Println(checksum)
	time.Sleep(time.Second * 2)

	fmt.Println("Start Login Timestamp", timestamp, "Checksum", checksum)
	info = login(timestamp, checksum)
	token := info.Data.Token
	fmt.Println("Login Success Token", token)
	return timestamp, checksum, token
}
