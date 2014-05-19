package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
)

var cyan string = "0;36"
var yellow string = "1;33"

var pumpkinName = flag.String("pumpkin", ".Pumpkin", "The location of your pumpkin")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pumpkin := NewPumpkin(*pumpkinName)
	logger := make(chan string)
	watcher, err := fsnotify.NewWatcher()
	check(err)

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if !ev.IsAttrib() {
					pumpkin.Carve(ev.Name, logger)
				}
			}
		}
	}()

	go func() {
		useCyan := true
		for {
			select {
			case output := <-logger:
				var color string
				if useCyan {
					color = cyan
				} else {
					color = yellow
				}
				if len(output) > 0 {
					out := fmt.Sprintf("\033[%sm%s\033[0m", color, output)
					useCyan = !useCyan
					fmt.Println(out)
				}
			}
		}
	}()

	err = watcher.Watch(".")
	check(err)

	<-done

	watcher.Close()

}
