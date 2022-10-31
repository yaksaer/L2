package main

import "fmt"

type weapon interface {
	useWeapon(target *character)
}

type axe struct {
	name   string
	damage int
}

func newAxe(name string, damage int) *axe {
	return &axe{
		name:   name,
		damage: damage,
	}
}

func (a *axe) useWeapon(c *character) {
	fmt.Printf("hits the %s with a %s\n", c.name, a.name)
	c.health -= a.damage
}

type sword struct {
	name   string
	damage int
}

func newSword(name string, damage int) *sword {
	return &sword{
		name:   name,
		damage: damage,
	}
}

func (a *sword) useWeapon(c *character) {
	fmt.Printf("hits the %s with a %s\n", c.name, a.name)
	c.health -= a.damage
}

type bow struct {
	name   string
	damage int
}

func newBow(name string, damage int) *bow {
	return &bow{
		name:   name,
		damage: damage,
	}
}

func (a *bow) useWeapon(c *character) {
	fmt.Printf("shoots the %s with a %s\n", c.name, a.name)
	c.health -= a.damage
}

type character struct {
	name   string
	health int
	weapon weapon
	damage int
}

func newCharacter(name string) *character {
	return &character{
		name:   name,
		health: 100,
		damage: 1,
	}
}

func (c *character) equipWeapon(w weapon) {
	c.weapon = w
}

func (c *character) attack(target *character) {
	fmt.Printf("The %s ", c.name)
	c.weapon.useWeapon(target)
}

func printStats(c *character) {
	fmt.Printf("The %s has %d health left.\n", c.name, c.health)
}

func main() {
	unforgedAxe := newAxe("The Unforged", 50)
	shaperSword := newSword("Summit Shaper", 46)
	polarBow := newBow("Polar Star", 44)

	player := newCharacter("Orodreth")
	player.equipWeapon(shaperSword)
	player2 := newCharacter("Aegnor")
	player2.equipWeapon(polarBow)
	enemy := newCharacter("The Dwimmerlaik")
	enemy.equipWeapon(unforgedAxe)

	player.attack(enemy)
	printStats(enemy)
	enemy.attack(player2)
	printStats(player2)
	player2.attack(enemy)
	printStats(enemy)
}
