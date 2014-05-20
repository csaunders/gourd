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
	pattern  *regexp.Regexp
	commands []string
}

func NewPumpkin(filename string) Pumpkin {

	bytes, err := ioutil.ReadFile(filename)
	check(err)

	var pumpkin struct {
		Pattern  string   `json:"pattern"`
		Commands []string `json:"commands"`
	}
	err = json.Unmarshal(bytes, &pumpkin)
	check(err)

	return Pumpkin{
		pattern:  regexp.MustCompile(pumpkin.Pattern),
		commands: pumpkin.Commands,
	}
}

func (p Pumpkin) Carve(filename string, output chan string) {
	if p.check(filename) {
		p.process(output)
	}
}

func (p Pumpkin) check(filename string) bool {
	return p.pattern.MatchString(filename)
}

func (p Pumpkin) process(output chan string) {
	for _, command := range p.commands {
		output <- command
		args := strings.Split(command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {

		}
		output <- out.String()
	}
}
