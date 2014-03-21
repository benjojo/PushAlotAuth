package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"log"
	"regexp"
	"strings"
)

func WatchFileSystem(path string, triggerwords []string, ntf Notifiers) { //token string, user string, banner string
	fRgx, err := regexp.Compile("[\\w\\d]+ from [\\.\\d]+")
	if err != nil {
		log.Fatalf("The regex didn't compile. This shouldn't be happening...")
	}

	t, err := tail.TailFile(path, tail.Config{Follow: true, ReOpen: true})
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
					title := fmt.Sprintf("Log from %s", GetHostName())
					if banner != "" {
						title = banner
					}
					SendPushAlot(title, token, line.Text)

					if fRgx.Match([]byte(line.Text)) {
						matches := strings.Split(fRgx.FindString(line.Text), " from ")
						notice := fmt.Sprintf("%s@%s logged into %s", matches[0], matches[1], GetHostName())
						Push(ntf, fmt.Sprintf("Login at %s", GetHostName()), notice)
					} else {
						Push(ntf, fmt.Sprintf("Login at %s", GetHostName()), line.Text)
					}
				}
			}
		}
	}
}

func main() {
	log.Println("PushOver Auth Notifier")
	cfg := GetCFG()
	for _, v := range cfg.Watches {
		log.Printf("Setting up watching for %s", v.Path)
		go WatchFileSystem(v.Path, v.TriggerWords, cfg.Notifications)
		// go WatchFileSystem(v.Path, v.TriggerWords, cfg.Token, v.Banner)
	}
	select {}
}
