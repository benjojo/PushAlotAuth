package main

import (
	"github.com/ActiveState/tail"
	"log"
	"strings"
)

func WatchFileSystem(path string, triggerwords []string, token string) {

	t, err := tail.TailFile(path, tail.Config{
		Follow: true,
		ReOpen: true})
	if err != nil {
		log.Fatalf("Uhh I can't read %s ... Maybe I am not root? Maybe you are on windows?", path)
	}
	lc := CountLines(path) - 1
	cnt := 0
	log.Printf("Now watching %s", path)
	for line := range t.Lines {
		if cnt < lc {
			cnt++
		} else {
			for _, v := range triggerwords {
				if strings.Contains(line.Text, v) {
					SendPushAlot("Login from server", token, line.Text)
				}
			}
			log.Println(line.Text)
		}
	}
}

func main() {
	log.Println("PushAlot Auth Notifier")
	cfg := GetCFG()
	for _, v := range cfg.Watches {
		log.Printf("Setting up watching for %s", v.Path)
		go WatchFileSystem(v.Path, v.TriggerWords, cfg.Token)
	}
	select {}
}
