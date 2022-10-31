package main

import "fmt"

type Handler interface {
	execute(*Request)
	setNext(Handler)
}

type VoiceAssist struct {
	next Handler
}

func (v *VoiceAssist) execute(r *Request) {
	fmt.Println("Find out what the client wants")
	r.voiceAssistDone = true
	if r.request == "Not working" {
		r.request = "Need tech help"
	}
	v.next.execute(r)
}

func (v *VoiceAssist) setNext(next Handler) {
	v.next = next
}

type Operator struct {
	next Handler
}

func (o *Operator) execute(r *Request) {
	if r.operatorDone {
		fmt.Println("Regular operator can't help client")
		o.next.execute(r)
		return
	}
	fmt.Println("Operator is trying to help client")
	r.operatorDone = true
	o.next.execute(r)
}

func (o *Operator) setNext(next Handler) {
	o.next = next
}

type SeniorOperator struct {
	next Handler
}

func (s *SeniorOperator) execute(r *Request) {
	if r.SeniorOperatorDone {
		fmt.Println("Senior operator can't help client")
		s.next.execute(r)
		return
	}
	fmt.Println("Senior operator is trying to help client")
	r.SeniorOperatorDone = true
	s.next.execute(r)
}

func (s *SeniorOperator) setNext(next Handler) {
	s.next = next
}

type TechSpec struct {
	next Handler
}

func (t *TechSpec) execute(r *Request) {
	if r.techSpecDone {
		fmt.Println("Tech specialist has already helped help client")
		return
	}
	fmt.Println("Tech specialist has helped client")
	r.techSpecDone = true
}

func (t *TechSpec) setNext(next Handler) {
	t.next = next
}

type Request struct {
	request            string
	voiceAssistDone    bool
	operatorDone       bool
	SeniorOperatorDone bool
	techSpecDone       bool
}

func main() {
	techSpec := &TechSpec{}

	senior := &SeniorOperator{}
	senior.setNext(techSpec)

	operator := &Operator{}
	operator.setNext(senior)

	voiceAssist := &VoiceAssist{}
	voiceAssist.setNext(operator)

	request := &Request{request: "Not working"}
	voiceAssist.execute(request)
	fmt.Println()
	voiceAssist.execute(request)
}
