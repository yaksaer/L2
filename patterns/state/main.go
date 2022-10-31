package main

import (
	"fmt"
	"log"
	"time"
)

type State interface {
	addResource(int) error
	receiveRequest(string) error
	processRequest() error
	answerRequest() error
}

type process struct {
	noResource     State
	receiveRequest State
	processRequest State
	answerRequest  State

	currentState State

	resources int
	request   string
}

func newProcess(resources int) *process {
	p := &process{
		resources: resources,
	}
	noResource := &noResourcesState{
		process: p,
	}
	receiveState := &waitRequestState{
		process: p,
	}
	processState := &processRequest{
		process: p,
	}
	answerState := &answerRequest{
		process: p,
	}

	p.setState(receiveState)
	p.noResource = noResource
	p.receiveRequest = receiveState
	p.processRequest = processState
	p.answerRequest = answerState
	return p
}

func (p *process) requestReceive(str string) error {
	return p.currentState.receiveRequest(str)
}

func (p *process) resourceAdd(num int) error {
	return p.currentState.addResource(num)
}

func (p *process) procRequest() error {
	return p.currentState.processRequest()
}

func (p *process) answRequest() error {
	return p.currentState.answerRequest()
}

func (p *process) setState(s State) {
	p.currentState = s
}

func (p *process) incrementItemCount(count int) {
	fmt.Printf("Adding %d resources\n", count)
	p.resources = p.resources + count
}

type noResourcesState struct {
	process *process
}

func (i *noResourcesState) addResource(num int) error {
	i.process.resources += num
	fmt.Println(num, "resources have been added")
	i.process.setState(i.process.receiveRequest)
	return nil
}

func (i *noResourcesState) receiveRequest(str string) error {
	return fmt.Errorf("no resources left\n")
}

func (i *noResourcesState) processRequest() error {
	return fmt.Errorf("no resources left\n")
}

func (i *noResourcesState) answerRequest() error {
	return fmt.Errorf("no resources left\n")
}

type waitRequestState struct {
	process *process
}

func (i *waitRequestState) receiveRequest(str string) error {
	if i.process.resources == 0 {
		i.process.setState(i.process.noResource)
		fmt.Println("no resources left\n")
		return nil
	}
	i.process.request = str
	fmt.Printf("request has been recieved\n")
	i.process.resources -= 1
	i.process.setState(i.process.processRequest)
	return nil
}

func (i *waitRequestState) addResource(num int) error {
	return fmt.Errorf("no request recieved yet\n")
}

func (i *waitRequestState) processRequest() error {
	return fmt.Errorf("no request recieved yet\n")
}

func (i *waitRequestState) answerRequest() error {
	return fmt.Errorf("no request recieved yet\n")
}

type processRequest struct {
	process *process
}

func (i *processRequest) receiveRequest(str string) error {
	return fmt.Errorf("request has been already recieved\n")
}

func (i *processRequest) addResource(num int) error {
	return fmt.Errorf("resources not needed\n")
}

func (i *processRequest) processRequest() error {
	fmt.Printf("processing request\n")
	fmt.Print("#")
	for j := 0; j < 10; j++ {
		fmt.Print("=")
		time.Sleep(1000000)
	}
	fmt.Println()
	i.process.setState(i.process.answerRequest)
	return nil
}

func (i *processRequest) answerRequest() error {
	return fmt.Errorf("request hasn't processed yet\n")
}

type answerRequest struct {
	process *process
}

func (i *answerRequest) receiveRequest(str string) error {
	return fmt.Errorf("request has been already recieved\n")
}

func (i *answerRequest) addResource(num int) error {
	return fmt.Errorf("resources not needed\n")
}

func (i *answerRequest) processRequest() error {
	return fmt.Errorf("request has been already processed\n")
}

func (i *answerRequest) answerRequest() error {
	fmt.Printf("answering request\n")
	fmt.Println("Answer to", i.process.request, "is", 42)
	i.process.setState(i.process.receiveRequest)
	return nil
}

func main() {
	process := newProcess(1)

	err := process.requestReceive("Meaning of life")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = process.procRequest()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = process.answRequest()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = process.requestReceive("Meaning of life")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = process.resourceAdd(5)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = process.requestReceive("2 * 21")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = process.procRequest()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = process.answRequest()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
