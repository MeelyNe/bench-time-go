package time_bench

import (
	"testing"
	"time"

	"time_bench/asm"
	"time_bench/c"
)

func BenchmarkGetCoarseTime(b *testing.B) {
	for b.Loop() {
		_ = c.GetCoarseTime()
	}
}

func BenchmarkGetGoTimeNow(b *testing.B) {
	for b.Loop() {
		_ = time.Now()
	}
}

func BenchmarkGetCoarseTimeByAsm(b *testing.B) {
	for b.Loop() {
		_, _ = asm.GetCoarseTime()
	}
}
