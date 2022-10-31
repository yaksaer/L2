package main

import "fmt"

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type InputCommand struct {
	system System
}

func (c *InputCommand) execute() {
	c.system.inputText()
}

type OutputCommand struct {
	system System
}

func (c *OutputCommand) execute() {
	c.system.outputText()
}

type System interface {
	inputText()
	outputText()
}

type Computer struct {
	isWorking bool
}

func (c *Computer) inputText() {
	c.isWorking = true
	fmt.Println("Computer is receiving message")
	c.isWorking = false
}

func (c *Computer) outputText() {
	c.isWorking = true
	fmt.Println("Computer is answering")
	c.isWorking = false
}

func main() {
	comp := &Computer{}
	input := &InputCommand{
		system: comp,
	}
	output := &OutputCommand{
		system: comp,
	}
	inButton := &Button{
		command: input,
	}
	outButton := &Button{
		command: output,
	}
	inButton.press()
	outButton.press()
}
