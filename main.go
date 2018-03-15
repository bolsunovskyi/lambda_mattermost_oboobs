package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ChannelID   string `json:"channel_id" schema:"channel_id"`
	Text        string `json:"text" schema:"text"`
	ChannelName string `json:"channel_name" schema:"channel_name"`
	Command     string `json:"command" schema:"command"`
	ResponseURL string `json:"response_url" schema:"response_url"`
	TeamDomain  string `json:"team_domain" schema:"team_domain"`
	TeamID      string `json:"team_id" schema:"team_id"`
	Token       string `json:"token" schema:"token"`
	UserID      string `json:"user_id" schema:"user_id"`
	UserName    string `json:"user_name"`
}

const baseURL = "http://media.oboobs.ru"

var boobSource = []string{"http://api.oboobs.ru/noise/1", "http://api.oboobs.ru/boobs/0/1/random"}

func getBoobsLink() (string, error) {
	source := boobSource[rand.Int31n(int32(len(boobSource)))]
	rq, err := http.NewRequest("GET", source, nil)
	if err != nil {
		return "", err
	}

	hCl := http.Client{Timeout: time.Second * 5}
	rsp, err := hCl.Do(rq)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	var m []map[string]interface{}
	if err := json.NewDecoder(rsp.Body).Decode(&m); err != nil {
		return "", err
	}

	if preview, ok := m[0]["preview"]; ok {
		if pStr, ok := preview.(string); ok {
			return baseURL + "/" + pStr, nil
		}
	}

	return "", errors.New("preview not found in response")
}

type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func Handler(_ Request) (interface{}, error) {
	boobLink, err := getBoobsLink()
	if err != nil {
		return nil, err
	}

	return Response{
		ResponseType: "in_channel",
		Text:         "![boobs](" + boobLink + ")",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
