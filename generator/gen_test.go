package generator

import (
	"fmt"
	"testing"
)

func TestStructure(t *testing.T) {
	m := Init()
	m.AddDriverName("BurpSuiteBisa15")

	err := m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables")
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddRule("sql-injection", "SQL Injection", "-", "Variables")
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddResults("no-unused-vars", "warning", "'x' is assigned a value but never used.", []string{"/http/api/user/2"})
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddResults("sql-injection", "error", "SQL Injection in this URL", []string{"/http/api/sqli/14"})
	if err != nil {
		t.Fatal(err)
	}

	if len(m.Runs) != 1 {
		t.Fatal("This tool only generates maximum 1 run, hardcoded just like that")
	}

	if len(m.Runs[0].Tool.Driver.Rules) != 2 {
		t.Fatal("The rules should be 2")
	}

	if len(m.Runs[0].Results) != 2 {
		t.Fatal("The rules should be 2")
	}

	j, err := m.GetJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(j)
}
