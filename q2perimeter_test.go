package main
import "testing"
func Test_perimeter(t *testing.T) {
     r1 := Rectangle{0,0,10,10}
        c1 := Circle{0,0,5}
        x:= totalPerimeter(&r1,&c1)
   if x != 71.41592653589794{
    t.Error("Expected 71.41592653589794, got ", x)
}
 }
func Test_area(t *testing.T) {
     r1 := Rectangle{0,0,10,10}
        c1 := Circle{0,0,5}
        x:= totalArea(&r1,&c1)
   if x != 178.539816339744854{
    t.Error("Expected 178.53981633974485, got ", x)
 }
}
