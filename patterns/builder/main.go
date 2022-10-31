package main

import "fmt"

type appleBuilder struct {
	chipset string
	OS      string
	body    string
}

func newAppleBuilder() *appleBuilder {
	return &appleBuilder{}
}

func (b *appleBuilder) setChipset() {
	b.chipset = "M1"
}

func (b *appleBuilder) setOS() {
	b.OS = "IOS"
}

func (b *appleBuilder) setBody() {
	b.body = "IPhone body"
}

func (b *appleBuilder) getPhone() Phone {
	return Phone{
		chipset: b.chipset,
		OS:      b.OS,
		body:    b.body,
	}
}

type Director struct {
	builder builderI
}

func newDirector(b builderI) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b builderI) {
	d.builder = b
}

func (d *Director) assemblePhone() Phone {
	d.builder.setChipset()
	d.builder.setOS()
	d.builder.setBody()
	return d.builder.getPhone()
}

type samsungBuilder struct {
	chipset string
	OS      string
	body    string
}

func newSamsungBuilder() *samsungBuilder {
	return &samsungBuilder{}
}

func (b *samsungBuilder) setChipset() {
	b.chipset = "Qualcomm"
}

func (b *samsungBuilder) setOS() {
	b.OS = "Android"
}

func (b *samsungBuilder) setBody() {
	b.body = "Samsung body"
}

func (b *samsungBuilder) getPhone() Phone {
	return Phone{
		chipset: b.chipset,
		OS:      b.OS,
		body:    b.body,
	}
}

type Phone struct {
	chipset string
	OS      string
	body    string
}

type builderI interface {
	setChipset()
	setOS()
	setBody()
	getPhone() Phone
}

func makeBuilder(buildType string) builderI {
	if buildType == "apple" {
		return newAppleBuilder()
	} else if buildType == "samsung" {
		return newSamsungBuilder()
	}
	return nil
}

func main() {
	appleBuilder := makeBuilder("apple")
	samsungBuilder := makeBuilder("samsung")
	director := newDirector(appleBuilder)
	IPhone := director.assemblePhone()
	fmt.Printf("Iphone tech details:\nChipset type - %s\nOS type - %s\nBody type - %s\n\n",
		IPhone.chipset, IPhone.OS, IPhone.body)

	director.setBuilder(samsungBuilder)
	samsungGalaxy := director.assemblePhone()
	fmt.Printf("Samsung tech details:\nChipset type - %s\nOS type - %s\nBody type - %s\n",
		samsungGalaxy.chipset, samsungGalaxy.OS, samsungGalaxy.body)
}
