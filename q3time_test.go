package main

import ("testing"
        "time")
func Test_sleeptime(t *testing.T) {

    start := time.Now()
    sleeptime(int(5))
    sec := time.Since(start).Seconds()

    if sec < 5 || sec > 5.30 {
        t.Error("error in sleep function")
    }
}
