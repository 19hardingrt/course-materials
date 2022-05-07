// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "/workspace/materals/lab/7/main/gmail-alleged.txt") // Currently function returns only number of open ports
	want := "Nickelback4life"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

// With go routines
func TestGenHashMapsRoutines(t *testing.T) {
	start := time.Now()
	GenHashMaps("/workspace/materals/lab/7/main/gmail-alleged.txt", true)
	duration := time.Since(start)
	fmt.Printf("Time to generate hash maps with go routines: %f\nTime per password: %f\n", duration.Seconds(), duration.Seconds() / 1667462)

}

// Without go routines
func TestGenHashMapsNoRoutines(t *testing.T) {
	start := time.Now()
	GenHashMaps("/workspace/materals/lab/7/main/gmail-alleged.txt", false)
	duration := time.Since(start)
	fmt.Printf("Time to generate hash maps without go routines: %f\nTime per password: %f\n", duration.Seconds(), duration.Seconds() / 1667462)

}


