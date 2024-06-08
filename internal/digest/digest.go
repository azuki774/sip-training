package digest

import (
	"crypto/md5"
	"fmt"
	"strings"
)

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
func (w *WWWAuthenticate) ComputeResponse(method string, password string) (response string) {
	A1 := strings.Join([]string{w.UserName, w.Realm, password}, ":")
	A2 := strings.Join([]string{method, w.URI}, ":")
	HA1 := fmt.Sprintf("%x", md5.Sum([]byte(A1)))
	HA2 := fmt.Sprintf("%x", md5.Sum([]byte(A2)))
	responseTmp := strings.Join([]string{HA1, w.NonceValue, w.NonceCount, w.CNonse, w.QOP, HA2}, ":")
	return fmt.Sprintf("%x", md5.Sum([]byte(responseTmp)))
}
