package c

/*
#include <time.h>
#include <stdint.h>

int64_t get_coarse_time() {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME_COARSE, &ts);
    return (int64_t)ts.tv_sec * 1000000000LL + (int64_t)ts.tv_nsec;
}
*/
import "C"

func GetCoarseTime() int64 {
	return int64(C.get_coarse_time())
}
