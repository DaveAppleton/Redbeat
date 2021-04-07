package exdb

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInsert(t *testing.T) {
	n := NewExchange("TEST1", "TEST2")
	fmt.Println(n.Price)
	v, _ := decimal.NewFromString("3.14159")
	id, err := n.updateExchange(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
	id, err = n.UpdateExchange(v)
	if err != nil {
		t.Fatal(err)
	}
	if id != 0 {
		t.Fatal("should return zero", id)
	}
}
