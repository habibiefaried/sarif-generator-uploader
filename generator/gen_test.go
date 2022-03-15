package generator

import (
	"fmt"
	"testing"
)

func TestStructure(t *testing.T) {
	m := Init()
	m.AddDriverName("BurpSuiteBisa9")

	err := m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables")
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddResults("no-unused-vars", "warning", "'x' is assigned a value but never used.", []string{"/http/api/user/2"})
	if err != nil {
		t.Fatal(err)
	}

	j, err := m.GetJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(j)
}
