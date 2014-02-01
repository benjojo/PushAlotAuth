package main

import (
	"net/http"
	"net/url"
)

func SendPushOver(Title string, Token string, User string, Body string) {
	http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{
			"token": {Token},
			"Body": {Body},
			"Title": {Title},
		}
	)
}