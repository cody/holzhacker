package holzhacker

import "testing"

func TestHolzhacker(t *testing.T) {
	myLog :=  Create("filename")

	myLog.Print("Hello")
	myLog.Printf("1 + 2 = %d", 1+2)
}
