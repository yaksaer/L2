package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

type Pwd struct {
	currentPath string
}

func (p *Pwd) Execute() string {
	p.currentPath, _ = os.Getwd()
	return p.currentPath
}

type Cd struct {
	Pwd
}

func (c *Cd) Execute(path string) {
	var sb strings.Builder
	var pwd Pwd
	sb.WriteString(pwd.Execute())
	sb.WriteString("/")
	if len(path) == 0 {
		sb.WriteString("..")
	} else {
		sb.WriteString(path)
	}
	err := os.Chdir(sb.String())
	if err != nil {
		fmt.Println(err)
	}
}

func Kill(pid string) {
	pidInt, _ := strconv.Atoi(pid)
	s, err := os.FindProcess(pidInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Kill()
}

type Ps struct {
	info string
}

func (p *Ps) Executable() {
	processList, _ := ps.Processes()
	var sb strings.Builder
	sb.WriteString("PID\tName\n")
	for _, j := range processList {
		sb.WriteString(strconv.Itoa(j.Pid()))
		sb.WriteString("\t")
		sb.WriteString(j.Executable())
		sb.WriteString("\n")
	}
	p.info = sb.String()

}

type Shell struct {
	currentPath string
	Cd
	Ps
}

type pipeline struct {
	command string
	args    []string
}

func (s *Shell) Start() {
	var inputTmp strings.Builder
	var sb strings.Builder
	tmp := pipeline{
		args: make([]string, 0, 1),
	}
	pipe := make([]pipeline, 0, 1)
	scanner := bufio.NewScanner(os.Stdin)
	isCommand := true
	for {
		s.currentPath = s.Pwd.Execute()
		sb.WriteString("Minishell> ")
		fmt.Printf("%s", sb.String())
		scanner.Scan()
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), " ")
		line = strings.ReplaceAll(line, " |", "|")
		line = strings.ReplaceAll(line, "| ", "|")
		for _, v := range line {
			if string(v) == " " && isCommand {
				isCommand = false
				tmp.command = inputTmp.String()
				inputTmp.Reset()
				continue
			}
			if string(v) == " " && !isCommand {
				tmp.args = append(tmp.args, inputTmp.String())
				inputTmp.Reset()
				continue
			}
			if string(v) == "|" {
				if isCommand {
					tmp.command = inputTmp.String()
				}
				if !isCommand {
					tmp.args = append(tmp.args, inputTmp.String())
				}
				pipe = append(pipe, tmp)
				tmp.args = nil
				isCommand = true
				inputTmp.Reset()
				continue
			}
			inputTmp.WriteRune(v)

		}
		if isCommand {
			tmp.command = inputTmp.String()
			tmp.args = nil
			pipe = append(pipe, tmp)
			inputTmp.Reset()
		}
		if !isCommand {
			tmp.args = append(tmp.args, inputTmp.String())
			pipe = append(pipe, tmp)
			tmp.args = nil
			inputTmp.Reset()
		}
		for _, cmd := range pipe {

			switch cmd.command {
			case "\\quit":
				os.Exit(0)
			case "pwd":
				fmt.Println(s.Cd.Pwd.Execute())
			case "cd":
				if len(cmd.args) == 0 {
					s.Cd.Execute("")
				} else {
					s.Cd.Execute(cmd.args[0])
				}
			case "echo":
				for _, val := range cmd.args {
					fmt.Printf("%s ", val)
				}
				fmt.Println()
			case "ps":
				s.Ps.Executable()
				fmt.Println(s.Ps.info)
			case "kill":
				Kill(cmd.args[0])
			default:
				fmt.Println("Command:", line, "not found")
			}

		}
		pipe = nil
		sb.Reset()
		isCommand = true
		inputTmp.Reset()
	}

}

func main() {

	var s Shell
	s.Start()

}
