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
	var data map[string]interface{}
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	err = json.Unmarshal(bytes, &data)
	check(err)

	patternString := data["pattern"].(string)
	pattern, err := regexp.Compile(patternString)
	check(err)

	commandArray := data["commands"].([]interface{})
	commands := make([]string, len(commandArray))
	for i := 0; i < len(commandArray); i++ {
		commands[i] = commandArray[i].(string)
	}
	return Pumpkin{pattern: pattern, commands: commands}
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
