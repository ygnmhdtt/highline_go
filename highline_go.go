package main

import (
	"io"
	"regexp"
	"log"
	"fmt"
)

type MaskWriter struct {
	masks []string
	Writer io.Writer
}

func NewMaskWriter(maskStrings []string) *MaskWriter {
	m := new(MaskWriter)
	m.masks = maskStrings
	return m
}

func (m *MaskWriter) Write(p []byte) (n int, err error) {
  str := string(p)
	return m.Writer.Write(m.mask(str))
}

func (m *MaskWriter) mask(str string) []byte {
	for maskString := range m.masks {
		fmt.Println(maskString)
		re := regexp.MustCompile(fmt.Sprintf(`"%v":\s.+,`, maskString))
		fmt.Println(fmt.Sprintf(`"%v":\s.+,`, maskString))
		str = re.ReplaceAllString(string(str), "\"password\": \"'FILTERED'\",")
	}
	return []byte(str)
}

func main() {
	m := NewMaskWriter([]string{"password", "initial_password"})
	log.SetOutput(m)

	str := `{
"organization_id": 7,
"password": "060920171850capybaratest",
"display_name": "カピバラテスト",
"project_code": "capybaratest",
"email": "060920171850capybara@example.com"
}`

	log.Print(str)
}
