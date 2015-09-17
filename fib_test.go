package main

import "testing"

func Test_fibseries(t *testing.T) {

var x uint
x=fibseries(0)
if x != 0 {
t.Error("Expected 0, got ", x)
}
}
