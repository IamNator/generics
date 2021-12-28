package model

import "testing"

func TestPassword_Hash(t *testing.T) {
	var p Password
	p = Password("IamNatorjknki")

	t.Log(len(p.Hash()))
}
