package main

import (
	"io"
	"regexp"
	"log"
	"fmt"
	"os"
)

type MaskWriter struct {
	masks []string
	Writer io.Writer
}

func NewMaskWriter(writer io.Writer) *MaskWriter {
	m := new(MaskWriter)
	m.Writer = writer
	return m
}

func (m *MaskWriter) SetFilter(str []string) {
	m.masks = str
}

func (m *MaskWriter) Write(p []byte) (n int, err error) {
	return m.Writer.Write(m.mask(p))
}

func (m *MaskWriter) mask(p []byte) []byte {
	str := string(p)
	for _, maskString := range m.masks {
		re := regexp.MustCompile(fmt.Sprintf(`"%v":\s.+,`, maskString))
		str = re.ReplaceAllString(string(str), fmt.Sprintf(`"%v": "[FILTERED]",`, maskString))
	}
	return []byte(str)
}

func main() {
	m := NewMaskWriter(os.Stdout)
	m.SetFilter([]string{"password","initial_password"})
	log.SetOutput(m)

	str := `{
"id": 1,
"password": "password",
"display_name": "test",
"email": "test@example.com"
 }`

	log.Print(str)
}
