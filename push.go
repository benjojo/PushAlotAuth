package main

import (
	"net/http"
	"net/url"
)

func SendPushOver(Title string, Token string, User string, Message string) {
	http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{
			"token":   {Token},
			"user":    {User},
			"message": {Message},
			"Title":   {Title},
		}
	)
}

func SendPushAlot(Title string, Token string, Body string) {
	http.PostForm("https://pushalot.com/api/sendmessage",
		url.Values{
			"AuthorizationToken": {Token},
			"Body": {Body},
			"Title": {Title},
		}
	)
}
