package model

import (
	"fmt"
	"strings"
)

type SIPMethod string // Ex. REGISTER

var SIPMethodRegister = SIPMethod("REGISTER")

type RequestLine struct {
	Method     SIPMethod
	RequestURI string
	Transport  string // UDP
}

func (r *RequestLine) Build() string {
	s := fmt.Sprintf("%s %s;transport=%s\n", string(r.Method), r.RequestURI, r.Transport)
	return s
}

func stringsOptBuild(ss []string) (s string) {
	return strings.Join(ss, ";") // "a;b;c"
}

// fieldName: value の形式で string 出力する。ただし、value が 空文字列なら、出力も空文字列。
func ColonFieldBuild(fieldName string, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s: %s\r\n", fieldName, value)
}
