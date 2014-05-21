package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

type Pumpkin struct {
	Regex    *regexp.Regexp
	Pattern  string   `json:"pattern"`
	Commands []string `json:"commands"`
}

func NewPumpkinFromFile(filename string) Pumpkin {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return NewPumpkin(bytes)
}

func NewPumpkin(bytes []byte) Pumpkin {
	var pumpkin Pumpkin
	err := json.Unmarshal(bytes, &pumpkin)
	check(err)

	pumpkin.Regex = regexp.MustCompile(pumpkin.Pattern)
	return pumpkin
}

func (p Pumpkin) Carve(filename string, output chan Message) {
	if p.Check(filename) {
		p.Process(output)
	}
}

func (p Pumpkin) Check(filename string) bool {
	return p.Regex.MatchString(filename)
}

func (p Pumpkin) Process(output chan Message) {
	var stdout, stderr bytes.Buffer
	output <- Message{Color: yellow, Content: "---------------"}
	for _, command := range p.Commands {
		args := strings.Split(command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()

		msg := &Message{Color: green, Content: stdout.String()}
		if msg.IsAvailable() {
			output <- Message{Color: yellow, Content: command}
		}

		if err != nil {
			msg.Color = red
			output <- *msg
			msg.Content = stderr.String()
		}

		output <- *msg
	}
}
