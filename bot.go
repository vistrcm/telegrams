package telegrams

import (
	"strings"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)

type telegrambot struct {
	baseUrl string
	token   string
	httpClient *http.Client
}

// NewTelegramBot create new instance of telegrambot.
func NewTelegramBot(token string) telegrambot {
	telegrambot := telegrambot{}
	telegrambot.token = token
	telegrambot.baseUrl = "https://api.telegram.org/bot" + token

	telegrambot.httpClient = &http.Client{
		Timeout: time.Second * 30,
	}
	return telegrambot
}

func (bot telegrambot)request(path string) *http.Response {
	trimmedPath := strings.Trim(path, "/") // remove '/' from method
	url := bot.baseUrl + "/" + trimmedPath

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicf("Error happened %v, req: %v", err, req)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := bot.httpClient.Do(req)
	if err != nil {
		log.Panicf("error on sending request: %v, resp: %v", err, resp)
	}
	return resp
}

func (bot telegrambot) GetMe() User {
	resp := bot.request("getMe")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error on reading: %v. body: %v", err, body)
	}

	var apiResp UserAPIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Panicf("something happened during unmarshall: %v. engineers: %v", err, apiResp)
	}

	if ! apiResp.Ok {
		log.Panicf("API Response is not ok. %+v", apiResp)
	}
	return apiResp.Result
}

// set webhook to get messages
func (bot telegrambot) SetWebhook(url string){
	// first unregister all webhooks
	bot.request("setWebhook")
	// now register readl webhook url
	bot.request("setWebhook")
}
