package digest

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// WWWAuthenticate: この構造体には " " をつけない
type WWWAuthenticate struct {
	AuthenticateScheme string
	UserName           string // username
	Realm              string // realm
	NonceValue         string // nonce
	OpaqueValue        string // opaque
	Algorithm          string // algorithm
	QOP                string // qop
	URI                string // uri
	NonceCount         string // nc
	//
	CNonse   string // cnonce
	Response string // response
}

// MD5 and qop = auth only.
// This function does not refresh NC,CNonse.
func (w *WWWAuthenticate) ComputeResponse(method string, password string) {
	A1 := strings.Join([]string{w.UserName, w.Realm, password}, ":")
	A2 := strings.Join([]string{method, w.URI}, ":")
	HA1 := fmt.Sprintf("%x", md5.Sum([]byte(A1)))
	HA2 := fmt.Sprintf("%x", md5.Sum([]byte(A2)))
	responseTmp := strings.Join([]string{HA1, w.NonceValue, w.NonceCount, w.CNonse, w.QOP, HA2}, ":")
	w.Response = fmt.Sprintf("%x", md5.Sum([]byte(responseTmp)))
}

func (w *WWWAuthenticate) Build() (s string) {
	if w.AuthenticateScheme == "" {
		return ""
	}
	s = w.AuthenticateScheme + " "
	s += fmt.Sprintf("username=\"%s\",", w.UserName)
	s += fmt.Sprintf("realm=\"%s\",", w.Realm)
	s += fmt.Sprintf("nonce=\"%s\",", w.NonceValue)
	s += fmt.Sprintf("uri=\"%s\",", w.URI)
	s += fmt.Sprintf("response=\"%s\",", w.Response)
	s += fmt.Sprintf("cnonce=\"%s\",", w.CNonse)
	s += fmt.Sprintf("nc=%s,", w.NonceCount)
	s += fmt.Sprintf("qop=%s,", w.QOP)
	s += fmt.Sprintf("algorithm=%s,", w.Algorithm)
	s += fmt.Sprintf("opaque=\"%s\"", w.OpaqueValue)
	return s
}
