package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const telegramUrlFormat = "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s"

var telegramToken = ""
var telegramChatId = ""

func log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if telegramToken == "" && telegramChatId == "" {
		println(msg)

		return
	}

	if response, err := http.Post(fmt.Sprintf(telegramUrlFormat, telegramToken, telegramChatId, url.QueryEscape(msg)), "", nil); err != nil || response.StatusCode/100 != 2 {
		println(msg)
	}
}
