package highline

import (
	"bytes"
	"log"
	"testing"
	"regexp"
)

func TestWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	m := NewMaskWriter(buf)
	m.SetFilter([]string{"password"})
	log.SetOutput(m)

	str := `{
"id": 1,
"password": "password",
"display_name": "test",
"email": "test@example.com"
 }
`

	// test
	log.Print(str)
	actual := buf.String()
	rep := regexp.MustCompile(`^.{4}/.{2}/.{2} .{2}:.{2}:.{2} `)
	actual = rep.ReplaceAllString(actual, "")

	expected := `{
"id": 1,
"password": "[FILTERED]",
"display_name": "test",
"email": "test@example.com"
 }
`

	if actual != expected {
		t.Log("|" + expected + "|")
		t.Log("|" + actual + "|")
		t.Errorf("expected is %v, but actual was %v", expected, actual)
	}
}
