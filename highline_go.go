package main

import (
	"fmt"
	"regexp"
)

type MaskStrings struct {
	maskStrings []string
}

func NewMaskStrings(s []string) *MaskStrings {
	m := new(MaskStrings)
	m.maskStrings = s
	fmt.Println(m.maskStrings)
	return m
}

func (m *MaskStrings) Mask(str string) string {
	for maskString := range m.maskStrings {
		// TODO: maskStringを使う
		// TODO: maskStringが正規表現の場合に対応する
		fmt.Println(maskString)
		re := regexp.MustCompile(fmt.Sprintf(`"%v":\s.+,`, maskString))
		fmt.Println(fmt.Sprintf(`"%v":\s.+,`, maskString))
		str = re.ReplaceAllString(str, `"password": "[FILTERED]",`)
	}
	return str
}

func main() {
	m := NewMaskStrings([]string{"password", "initial_password"})

	str := `{
"organization_id": 7,
"password": "060920171850capybaratest",
"display_name": "カピバラテスト",
"project_code": "capybaratest",
"email": "060920171850capybara@example.com"
}`

	fmt.Println(m.Mask(str))
}
