// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "../main/gmail-alleged.txt") // Currently function returns only number of open ports
	want := "p@ssword"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestGenHashMaps(t *testing.T) {
	GenHashMaps("../main/gmail-alleged.txt")
}

func DrBPassword(t *testing.T) {
	GenHashMaps("../main/gmail-alleged.txt")
	got1, _ := GetMD5("90f2c9c53f66540e67349e0ab83d8cd0")
	want1 := "p@ssword"
	got2, _ := GetSHA("1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032")
	want2 := "letmein"
	if got1 != want1 || got2 != want2 {
		t.Errorf("got1 %s, wanted1 %s\ngot2 %s, wanted2 %s", got1, want1, got2, want2)
	}

}


