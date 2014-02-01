package main

import (
	"net/http"
	"net/url"
)

func SendPushAlot(Title string, Token string, Body string) {
	http.PostForm("https://pushalot.com/api/sendmessage",
		url.Values{
			"AuthorizationToken": {Token},
			"Body":               {Body},
			"Title":              {Title},
		})
}
