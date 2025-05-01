```
goos: linux
goarch: amd64
pkg: time_bench
cpu: AMD Ryzen 5 7400F 6-Core Processor             
BenchmarkGetCoarseTime
BenchmarkGetCoarseTime-12         	33037270	        33.09 ns/op
BenchmarkGetGoTimeNow
BenchmarkGetGoTimeNow-12          	31677024	        36.45 ns/op
BenchmarkGetCoarseTimeByAsm
BenchmarkGetCoarseTimeByAsm-12    	307548382	         3.965 ns/op
PASS

Process finished with the exit code 0
```