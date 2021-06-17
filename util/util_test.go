package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tel = []string{"979939999", "979935685"}
var telEncrypt = []string{"***9399**", "***9356**"}

var telZero = []string{"0979939999", "0979935685"}
var telRemoveZero = []string{"979939999", "979935685"}

func TestEndCryptTel(t *testing.T) {
	for i, v := range tel {
		assert.Equal(t, telEncrypt[i], EndCryptTel(v), "expect ***xxxx** : %v", v)
	}
}

func TestRemoveZeroFirst(t *testing.T) {
	for i, v := range telZero {
		assert.Equal(t, telRemoveZero[i], RemoveZeroFirst(v), "expect not first zero : %v", v)
	}
}
