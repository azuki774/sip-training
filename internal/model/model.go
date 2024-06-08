package model

import (
	"fmt"
	"regexp"
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

// コンマ区切り key = value をマップに変換
func ParseCommaEqualField(str string) map[string]string {
	// a = b, c = d ... { (a,b), (c,d)} というマップに変換
	f := make(map[string]string)
	commaSplit := strings.Split(str, ",") // a = b, c = d --> [a = b, c = d]

	for _, es := range commaSplit {
		// ex. CSeq: 1 REGISTER -> "CSeq", "1 REGISTER"
		e := strings.SplitN(es, "=", 2)
		if len(e) >= 2 {
			// = が1つ以上含まれるものだけが対象
			key := strings.TrimSpace(e[0])
			value := strings.TrimSpace(e[1]) // 2つ以上 : がある場合はすべて value に入る
			f[key] = value
		}
	}

	return f
}

// fieldName: value の形式の string 入力をマップする。
func ParseColonField(str string) map[string]string {
	// a: bbbbb
	// b: ccccc --> {(a, bbbb), (b, cccc)}
	f := make(map[string]string)
	reg := "\r\n|\n"
	rows := regexp.MustCompile(reg).Split(str, -1) // \r\n ごとに分けて配列に

	for _, row := range rows {
		// ex. CSeq: 1 REGISTER -> "CSeq", "1 REGISTER"
		splitRow := strings.SplitN(row, ":", 2)
		if len(splitRow) >= 2 {
			// : が1つ以上含まれるものだけが対象
			key := strings.TrimSpace(splitRow[0])
			value := strings.TrimSpace(splitRow[1]) // 2つ以上 : がある場合はすべて value に入る
			f[key] = value
		}
	}
	return f
}
