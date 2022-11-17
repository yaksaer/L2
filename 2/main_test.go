package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

type data struct {
	name   string
	input  string
	output string
	err    bool
}

func newData(name string, input string, output string, err bool) *data {
	return &data{
		name:   name,
		input:  input,
		output: output,
		err:    err,
	}
}

func TestMain_Unpack(t *testing.T) {
	data := make([]data, 5)
	data[0] = *newData("OK", "a4bc2d5e", "aaaabccddddde", false)
	data[1] = *newData("OK", "he3r9m8l2\\4\\5", "heeerrrrrrrrrmmmmmmmmll45", false)
	data[2] = *newData("Bad string", "93he3r9m8l2", "heeerrrrrrrrrmmmmmmmmll44444", true)
	data[3] = *newData("OK", "\\2\\2\\3he3r9m8l2", "223heeerrrrrrrrrmmmmmmmmll", false)
	data[4] = *newData("OK", "", "", false)
	for _, j := range data {
		t.Run(j.name, func(t *testing.T) {
			res, err := Unpack(j.input)
			if j.err {
				assert.NotEqual(t, j.err, nil)
			} else {
				assert.Equal(t, err, nil)
				assert.Equal(t, j.output, res)
			}
		})
	}
}
