package model

import "fmt"

type SIPMethod string // Ex. REGISTER

var SIPMethodRegister = SIPMethod("REGISTER")

type RequestLine struct {
	Method     SIPMethod
	RequestURI string
	Transport  string // UDP
}

func (r *RequestLine) Build() string {
	s := fmt.Sprintf("%s %s;transport=%s", string(r.Method), r.RequestURI, r.Transport)
	return s
}
