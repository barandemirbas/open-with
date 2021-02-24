package openwith

import (
	"testing"
)

func Test(t *testing.T) {
	out, err := Browser("golang.org")
	if err != nil {
		t.Error(out)
	}
}
