package main

import (
	"fmt"
	"math"
)

type Visitor interface {
	visitForBall(*ball)
	visitForCube(*cube)
	visitForCylinder(*cylinder)
}

type volumeCalculator struct {
	volume float64
}

func (v *volumeCalculator) visitForBall(b *ball) {
	v.volume = (math.Pi * math.Pow(b.radius, 3) * 4) / 3
	fmt.Println("Volume for ball =", v.volume)
}

func (v *volumeCalculator) visitForCube(c *cube) {
	v.volume = math.Pow(c.sideLen, 3)
	fmt.Println("Volume for cube =", v.volume)
}

func (v *volumeCalculator) visitForCylinder(cy *cylinder) {
	v.volume = math.Pi * math.Pow(cy.radius, 2) * cy.height
	fmt.Println("Volume for cylinder =", v.volume)
}

type planArea struct {
	area float64
}

func (base *planArea) visitForBall(b *ball) {
	base.area = math.Pi * math.Pow(b.radius, 2)
	fmt.Println("Plan area for ball =", base.area)
}

func (base *planArea) visitForCube(c *cube) {
	base.area = math.Pow(c.sideLen, 2)
	fmt.Println("Plan area for cube =", base.area)
}

func (base *planArea) visitForCylinder(cy *cylinder) {
	base.area = math.Pi * math.Pow(cy.radius, 2)
	fmt.Println("Plan area for cylinder =", base.area)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type ball struct {
	radius float64
}

func (b *ball) accept(v Visitor) {
	v.visitForBall(b)
}

func (b *ball) getType() string {
	return "ball"
}

type cube struct {
	sideLen float64
}

func (c *cube) accept(v Visitor) {
	v.visitForCube(c)
}

func (c *cube) getType() string {
	return "cube"
}

type cylinder struct {
	radius float64
	height float64
}

func (cy *cylinder) accept(v Visitor) {
	v.visitForCylinder(cy)
}

func (cy *cylinder) getType() string {
	return "cylinder"
}

func main() {
	ball := &ball{radius: 5}
	cube := &cube{sideLen: 8}
	cylinder := &cylinder{radius: 9, height: 12}

	volumeCalculator := &volumeCalculator{}

	ball.accept(volumeCalculator)
	cube.accept(volumeCalculator)
	cylinder.accept(volumeCalculator)

	fmt.Println()

	planArea := &planArea{}
	ball.accept(planArea)
	cube.accept(planArea)
	cylinder.accept(planArea)
}
