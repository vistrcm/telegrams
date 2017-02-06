package telegrams

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type telegrambot struct {
	baseUrl    string
	token      string
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

func (bot telegrambot) request(path string, params url.Values) (APIResponse, error) {
	trimmedPath := strings.Trim(path, "/") // remove '/' from method
	url := bot.baseUrl + "/" + trimmedPath

	resp, err := bot.httpClient.PostForm(url, params)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, err
	}

	var apiResp APIResponse
	err = json.Unmarshal(bytes, &apiResp)

	if err != nil {
		log.Printf("something happened during unmarshall: %v.\n", err)
		return APIResponse{}, err
	}

	if !apiResp.Ok {
		return APIResponse{}, errors.New(apiResp.Description)
	}

	return apiResp, nil
}

func (bot telegrambot) GetMe() (User, error) {
	resp, err := bot.request("getMe", nil)

	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(resp.Result, &user)

	if err != nil {
		log.Printf("something happened during unmarshall: %v.\n", err)
		return User{}, err
	}

	return user, nil
}
