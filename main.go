package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
)

var red string = "0;31"
var green string = "0;32"
var cyan string = "0;36"
var yellow string = "1;33"

type Message struct {
	Color   string
	Content string
}

func (m Message) String() string {
	return fmt.Sprintf("\033[%sm%s\033[0m", m.Color, m.Content)
}

func (m Message) IsAvailable() bool {
	return len(m.Content) > 0
}

var pumpkinName = flag.String("pumpkin", ".Pumpkin", "The location of your pumpkin")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pumpkin := NewPumpkinFromFile(*pumpkinName)
	logger := make(chan Message)
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
		for {
			select {
			case output := <-logger:
				if output.IsAvailable() {
					fmt.Println(output.String())
				}
			}
		}
	}()

	err = watcher.Watch(".")
	check(err)

	<-done

	watcher.Close()

}
