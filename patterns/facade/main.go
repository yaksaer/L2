package main

import (
	"fmt"
	"math/rand"
)

func (a *Aircraft) takeOff() {
	fmt.Println("Take off starts")
	if a.airport.checkAircraft() == 1 {
		fmt.Println("The aircraft didn't pass the pre-flight checkAircraft")
		return
	}
	fmt.Println("Aircraft is ready to take off")
	fmt.Println("Boarding begins")
	a.steward.startBoarding()
	a.steward.makeAnnounce()
	a.pilots.takeOff()
	fmt.Println("Take off has been completed")
	a.inAir = true
}

func (a *Aircraft) landing() {
	if a.inAir == false {
		fmt.Println("The aircraft is not in the air")
		return
	}
	fmt.Println("Landing starts")
	a.steward.makeAnnounce()
	a.pilots.landing()
	a.airport.meetAircraft()
	a.steward.landing()
	a.inAir = false
	fmt.Println("Landing ends")
}

func newAircraft(model string, age int) *Aircraft {
	fmt.Println("Creating new aircraft")
	return &Aircraft{
		model: model,
		age:   age,
		airport: &airport{
			age: age,
		},
		steward: &steward{},
		pilots:  &pilots{},
		inAir:   false,
	}
}

func (p airport) meetAircraft() {
	fmt.Println("Adjust the ladder")
}

func (s steward) landing() {
	fmt.Println("Open the door")
	fmt.Println("Say goodbye to passengers")
}

func (p pilots) landing() {
	fmt.Println("Turn off autopilot")
	fmt.Println("Start descent")
	fmt.Println("Landing")
	fmt.Println("Drive to the parking place")
	fmt.Println("Park the aircraft")
}

type Aircraft struct {
	model   string
	age     int
	airport *airport
	steward *steward
	pilots  *pilots
	inAir   bool
}

type pilots struct {
}

type airport struct {
	age int
}

type steward struct {
}

func (p pilots) takeOff() {
	fmt.Println("Start engine")
	fmt.Println("Drive to runway")
	fmt.Println("Taking off")
}

func (p airport) checkAircraft() int {
	fmt.Println("Start pre-flight check")
	fmt.Println("Checking aircraft's age")
	if p.age > 20 {
		fmt.Println("The aircraft is too old for flights")
		return 1
	}
	fmt.Println("Checking aircraft's systems")
	if rand.Int() < 100 {
		fmt.Println("The aircraft needs repairs")
		return 1
	}
	fmt.Println("Checking pilots")
	if rand.Int() < 50 {
		fmt.Println("One of pilots can't fly")
		return 1
	}
	fmt.Println("The aircraft has passed the pre-flight check")
	return 0
}

func (s steward) startBoarding() {
	fmt.Println("Stewards are welcoming passengers")
	fmt.Println("Boarding ends")
}

func (s steward) makeAnnounce() {
	fmt.Println("Stewards make an announcement")
}

func main() {
	aircraft := newAircraft("Boeing 737", 10)
	aircraft.takeOff()
	fmt.Println()
	aircraft.landing()
	fmt.Println()
	aircraft = newAircraft("Sesna", 21)
	aircraft.takeOff()
	fmt.Println()
	aircraft.landing()
}
