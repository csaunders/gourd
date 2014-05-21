package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var basicPumpkin string = `
{
  "pattern": ".*\\.go",
  "commands": [
    "echo hello"
  ]
}`

func PumpkinToTest() Pumpkin {
	return NewPumpkin([]byte(basicPumpkin))
}

func TestExtractingAPumpkin(t *testing.T) {
	pumpkin := PumpkinToTest()
	assert.Equal(t, ".*\\.go", pumpkin.Regex.String())
}

func TestCheckingAFilename(t *testing.T) {
	pumpkin := PumpkinToTest()
	assert.Equal(t, true, pumpkin.Check("main.go"))
	assert.Equal(t, false, pumpkin.Check("main.c"))
	assert.Equal(t, false, pumpkin.Check("maindgo"))
}

func TestRunningCommands(t *testing.T) {
	var msg Message
	pumpkin := PumpkinToTest()
	output := make(chan Message)
	go func() {
		pumpkin.Process(output)
	}()
	validateMessage(t, <-output, yellow, "---------------")
	validateMessage(t, <-output, yellow, "echo hello")
	validateMessage(t, <-output, green, "hello\n")
	select {
	case msg = <-output:
		assert.Fail(t, "Received an unexpected message", msg)
	default:
		// Success
	}
}

func validateMessage(t *testing.T, msg Message, color, content string) {
	assert.Equal(t, color, msg.Color)
	assert.Equal(t, content, msg.Content)
}
