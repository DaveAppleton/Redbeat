package gasnow

import (
	"fmt"
	"testing"
)

func TestGN(t *testing.T) {
	reply, err := GetLatest()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reply)
	t.Fail()
}
