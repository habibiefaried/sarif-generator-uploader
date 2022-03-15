package generator

import (
	"fmt"
	"testing"
)

func TestStructure(t *testing.T) {
	m := Init()
	m.AddDriverName("BurpSuiteBisa9")
	m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables")
	j, err := m.GetJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(j)
}
