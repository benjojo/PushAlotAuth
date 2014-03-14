package main

import (
	"net/http"
	"net/url"
)

func Push(ntf Notifiers, Title string, Message string) {
	if IsToken(ntf.PushOver.UserToken) && IsToken(ntf.PushOver.AppToken) {
		SendPushOver(Title, Message, ntf.PushOver.AppToken, ntf.PushOver.UserToken)
	} else if IsToken(ntf.PushAlot.Token) {
		SendPushAlot(Title, Message, ntf.PushAlot.Token)
	}
}

func IsToken(token string) bool {
	return token != "" && token != "token"
}

func SendPushOver(Title string, Message string, Token string, User string) {
	http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{
			"token":   {Token},
			"user":    {User},
			"message": {Message},
			"Title":   {Title},
		})
}

func SendPushAlot(Title string, Body string, Token string) {
	http.PostForm("https://pushalot.com/api/sendmessage",
		url.Values{
			"AuthorizationToken": {Token},
			"Body": {Body},
			"Title": {Title},
		})
}
