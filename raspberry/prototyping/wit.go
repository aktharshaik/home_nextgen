package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const WIT_VERSION = "20141122"
const ACCESS_TOKEN = "MVNSQA57DWYL2G6XIOGRXFUGIY4LDLJE"

func main() {
	askWit("Switch on the light in bedroom")
}

func askWit(msg string) {
	msg, err := sanitizeQuerryString(msg)
	if err != nil {
		fmt.Println("Some thing went wrong")
		return
	}
	url := fmt.Sprintf("https://api.wit.ai/message?v=%s&q=%s", WIT_VERSION, msg)
	fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ACCESS_TOKEN))
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("Requesting wit's api gave: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("Something went really wrong with the response from Wit.ai")
		fmt.Println("Sorry, the machine learning service I use for my brain went down, @Diego: check the logs, there may be something for you there.")
		return
	}
	intent, _ := ioutil.ReadAll(res.Body)

	jsonString := string(intent[:])
	_ = jsonString

	var jsonResponse WitMessage
	err = json.Unmarshal(intent, &jsonResponse)
	if err != nil {
		log.Println("error parsing json: ", err)
		log.Printf("plain text json was %+v", jsonString)
	}
	//fmt.Println(res)
	//fmt.Println(res.Body)
	fmt.Println(jsonResponse)
}

func sanitizeQuerryString(str string) (string, error) {
	if len(url.QueryEscape(str)) > 255 {
		log.Println("Somebody talked too much, more than the 256 characters I can read.")
		errMsg := "Sorry, I can only read up to 256 characters and I didn't want to just ignore the end of your message."
		return "", errors.New(errMsg)
	}
	return url.QueryEscape(str), nil
}

type WitMessage struct {
	Text     string `json:"_text"`
	MsgID    string `json:"msg_id"`
	Outcomes []struct {
		Text       string  `json:"_text"`
		Confidence float64 `json:"confidence"`
		Entities   struct {
			Device []struct {
				Value string `json:"value"`
			} `json:"device"`
			OnOff []struct {
				Value string `json:"value"`
			} `json:"on_off"`
			Room []struct {
				Value string `json:"value"`
			} `json:"room"`
		} `json:"entities"`
		Intent string `json:"intent"`
	} `json:"outcomes"`
}
