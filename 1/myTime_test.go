package main

import (
	"github.com/go-playground/assert/v2"
	_ "github.com/go-playground/assert/v2"
	"testing"
)

type data struct {
	name   string
	server string
	err    bool
}

func newData(name string, server string, err bool) *data {
	return &data{
		name:   name,
		server: server,
		err:    err,
	}
}

func runTest(t *testing.T, test *data) {
	t.Run(test.name, func(t *testing.T) {
		time, err := GetTime(test.server)
		if test.err {
			assert.Equal(t, nil, time)
			assert.NotEqual(t, nil, err)
		} else {
			assert.Equal(t, nil, err)
			assert.NotEqual(t, nil, time)
		}
	})
}

func TestMain_GetTime(t *testing.T) {
	test1 := newData("OK", "0.beevik-ntp.pool.ntp.org", false)
	test2 := newData("Error: bad ntp server address", "NOT NTP SERVER", true)
	runTest(t, test1)
	runTest(t, test2)
}
