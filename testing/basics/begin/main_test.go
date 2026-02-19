// testing/basics/begin/main_test.go
package main
import "testing"

// write a test for sum
func TestSum(t *testing.T) {
	// test sum with integers
	ints := []int{1, 2, 3, 4, 5}
	expectedIntSum := 15
	if sum(ints...) != expectedIntSum {
		t.Errorf("Expected sum of %v to be %d, got %d", ints, expectedIntSum, sum(ints...))
	}
}
// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	// setup code
	println("Setting up tests...")
	// run tests
	m.Run()
}