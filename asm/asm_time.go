package asm

import (
	_ "unsafe"
)

var (
	vdsoClockgettimeSym uintptr
)

//go:linkname vdsoClockgettimeSym runtime.vdsoClockgettimeSym

//go:noescape
func GetCoarseTime() (sec int64, nsec int32)
