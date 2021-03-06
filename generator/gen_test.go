package generator

import (
	"testing"
)

func TestStructure1(t *testing.T) {
	m := Init()
	m.AddDriverName("BurpSuiteBisa15")

	err := m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables", "0.0")
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddRule("sql-injection", "SQL Injection", "-", "Variables", "9.8")
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

	_, err = m.GetJSON()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStructure2(t *testing.T) {
	m := Init()
	m.AddDriverName("BurpSuiteBisa15")

	err := m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables", "0.0")
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddResults("no-unused-vars", "warning", "'x' is assigned a value but never used.", []string{"/http/api/user/2"})
	if err != nil {
		t.Fatal(err)
	}

	err = m.AddResults("xss-injection", "error", "SQL Injection in this URL", []string{"/http/api/sqli/14"})
	if err == nil {
		t.Fatal("Must be error here!")
	} else {
		t.Log(err)
	}

	if len(m.Runs[0].Results) != 1 {
		t.Fatal("The rules should be 1")
	}

	_, err = m.GetJSON()
	if err != nil {
		t.Fatal(err)
	}
}
