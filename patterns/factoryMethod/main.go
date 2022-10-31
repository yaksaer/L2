package main

import "fmt"

type IHero interface {
	setName(name string)
	setPower(power string)
	getName() string
	getPower() string
}

type SuperHero struct {
	name  string
	power string
}

func (s *SuperHero) setName(name string) {
	s.name = name
}

func (s *SuperHero) getName() string {
	return s.name
}

func (s *SuperHero) setPower(power string) {
	s.name = power
}

func (s *SuperHero) getPower() string {
	return s.power
}

type Flash struct {
	SuperHero
}

func newFlash() IHero {
	return &Flash{
		SuperHero: SuperHero{
			name:  "Barry Alan",
			power: "Super speed",
		},
	}
}

type Batman struct {
	SuperHero
}

func newBatman() IHero {
	return &Batman{
		SuperHero: SuperHero{
			name:  "Bruce Wayne",
			power: "Money",
		},
	}
}

func getHero(heroType string) (IHero, error) {
	if heroType == "flash" {
		return newFlash(), nil
	}
	if heroType == "batman" {
		return newBatman(), nil
	}
	return nil, fmt.Errorf("wrong hero type")
}

func main() {
	flash, _ := getHero("flash")
	batman, _ := getHero("batman")
	fmt.Printf("Hero name: %s\nHero power: %s\n\n", flash.getName(), flash.getPower())
	fmt.Printf("Hero name: %s\nHero power: %s\n\n", batman.getName(), batman.getPower())
}
