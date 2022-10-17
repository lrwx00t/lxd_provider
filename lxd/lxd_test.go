// LXD provider test package
package lxd_test

import (
	"testing"

	"github.com/lrwx00t/lxd_provider/lxd"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Name     string
	Expected bool
}

// struct for testing multiple cases that contain errors
type testErrorCase struct {
	Name     string
	Expected string
}

func TestContainerName(t *testing.T) {
	cases := []testCase{
		{"alpine-c1", true},
		{"alpine-c2", false},
		{"alpine-c3", false},
	}
	for _, c := range cases {
		got := lxd.GetContainerWithName(c.Name)
		if got != c.Expected {
			t.Errorf("container: [%s] expected '%t', but got '%t'", c.Name, c.Expected, got)
		}
	}
}

func TestStopContainerWithName(t *testing.T) {
	errorMsg := "the provided name doesn't match any existing container"
	cases := []testErrorCase{
		{"alpine-c1", ""},
		{"alpine-c2", errorMsg},
		{"alpine-c3", errorMsg},
	}
	for _, c := range cases {
		err := lxd.StopContainerWithName(c.Name)
		if err != nil {
			assert.Containsf(t, err.Error(), c.Expected, "expected error containing %s, got %s", c.Expected, err)
		}
	}
}

func TestStartContainerName(t *testing.T) {
	errorMsg := "the provided name doesn't match any existing container"
	cases := []testErrorCase{
		{"alpine-c1", ""},
		{"alpine-c2", errorMsg},
		{"alpine-c3", errorMsg},
	}
	for _, c := range cases {
		err := lxd.StartContainerWithName(c.Name)
		if err != nil {
			assert.Containsf(t, err.Error(), c.Expected, "expected error containing %s, got %s", c.Expected, err)
		}
	}
}
