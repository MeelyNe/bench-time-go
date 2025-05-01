package asm

import (
	"testing"
)

func TestGetCoarseTime(t *testing.T) {
	sec, nsec := GetCoarseTime()
	t.Log(sec, nsec)
}
