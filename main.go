package main

import (
	"github.com/ActiveState/tail"
	"log"
)

func main() {
	watchtarget := "/var/log/auth.log"
	log.Println("PushAlot Auth Notifier")
	t, err := tail.TailFile(watchtarget, tail.Config{
		Follow: true,
		ReOpen: true})
	if err != nil {
		log.Fatal("Uhh I can't read /var/log/auth.log ... Maybe I am not root? Maybe you are on windows?")
	}
	lc := CountLines(watchtarget) - 1
	cnt := 0
	for line := range t.Lines {
		if cnt < lc {
			cnt++
		} else {
			log.Println(line.Text)
		}
	}
}
