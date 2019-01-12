package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	s, _ := GenerateToken("1", "sujian", "121116111@qq.com")
	fmt.Println(s)

}

func TestParseToken(t *testing.T) {
	//setting.SetUp()
	claims, e := ParseToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoiMSIsInVzZXJuYW1lIjoic3VqaWFuIiwiZXhwIjoxNTQ3NjI4NjM0LCJlbWFpbCI6IjEyMTExNjExMTFAcXEuY29tIn0.T9wLoExAw5rbH4LtedO4rIIWn0ESMFIqUb1VQtvzuYY")
	//claims, e := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsInVzZXJuYW1lIjoic3VqaWFuIiwiZW1haWwiOiIxMjExMTYxMTFAcXEuY29tIiwiZXhwIjoxNTQ3MDM0MDI1LCJpc3MiOiJnaW4tYmxvZyJ9.pQBpGeEVgSN2orCCGTW6WukkxhUaE9edn7P52pYGRYA")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(claims)
}

func TestBase64(t *testing.T) {
	var header = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
	}
	fmt.Println(header)
	jsonValue, _ := json.Marshal(header)

	fmt.Println(jsonValue)
	fmt.Println(string(jsonValue) == `{"alg": "HS256", "typ": "JWT"}`)

	//fmt.Println(base64.URLEncoding.EncodeToString(jsonValue))

	fmt.Println(string(jsonValue))

	right := strings.TrimRight(base64.URLEncoding.EncodeToString(jsonValue), "=")
	fmt.Println(right)
}

func TestBase642(t *testing.T) {

	right := strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(`{"typ": "JWT", "alg": "HS256"}`)), "=")
	fmt.Println(right)
}
