package stripego_test

import (
	"testing"

	"github.com/pilinux/stripego"
)

func TestEnv(t *testing.T) {
	err := stripego.Env()
	if err != nil {
		t.Errorf(
			"failed to load .env: %v", err,
		)
	}
}
