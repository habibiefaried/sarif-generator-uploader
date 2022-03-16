package uploader

import (
	"github.com/habibiefaried/sarif-generator-uploader/generator"
	"math/rand"
	"os"
	"testing"
	"time"
)

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestSARIFUpload(t *testing.T) {
	if os.Getenv("GithubToken") == "" {
		t.Log("Cannot proceed with test without github token")
	} else {
		m := generator.Init()
		driverName := "Test_" + randStringBytes(8)
		t.Log("Send result using name " + driverName)
		m.AddDriverName(driverName)

		err := m.AddRule("no-unused-vars", "disallow unused variables", "-", "Variables", "0.0")
		if err != nil {
			t.Fatal(err)
		}
		err = m.AddRule("xss-injection", "XSS Injection", "-", "Variables", "9.8")
		if err != nil {
			t.Fatal(err)
		}

		err = m.AddResults("no-unused-vars", "warning", "'x' is assigned a value but never used.", []string{"/http/api/user/2"})
		if err != nil {
			t.Fatal(err)
		}

		err = m.AddResults("xss-injection", "error", "XSS Injection in this URL", []string{"/http/api/xss/14"})
		if err != nil {
			t.Fatal(err)
		}

		err = m.AddResults("xss-injection", "error", "XSS Injection in this URL", []string{"/http/api/xssadvanced/14"})
		if err != nil {
			t.Fatal(err)
		}

		sarifContent, err := m.GetJSON()
		if err != nil {
			t.Fatal(err)
		}

		ui := Init(os.Getenv("GithubToken"), "habibiefaried", "Vulnerability-goapp", sarifContent, 1)
		err = ui.UploadSARIF()
		if err != nil {
			t.Fatal(err)
		}
	}
}
