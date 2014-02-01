package main

import (
	"github.com/ActiveState/tail"
	"log"
)

func main() {
	log.Println("PushAlot Auth Notifier")
	t, err := tail.TailFile("/var/log/auth.log", tail.Config{
		Follow: true,
		ReOpen: true})
	if err != nil {
		log.Fatal("Uhh I can't read /var/log/auth.log ... Maybe I am not root? Maybe you are on windows?")
	}
	for line := range t.Lines {
		log.Println(line.Text)
	}
}
