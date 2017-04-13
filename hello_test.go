package hello
 
import "testing"

func TestHello(t *testing.T){
  expected := "hello, buddies"
  actual   := "hello, buddies"
  if actual != expected {
    t.Error("Test failed, expected: '%s', got:  '%s'", expected, actual)
  } 
}
