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
	return strings.Join(ss, ",") // "a,b,c"
}

type MessageHeader struct {
	CallID        string   // Call-ID
	Contract      string   // Contract
	CSeq          CSeq     // CSeq
	From          From     // From
	To            To       // To
	Via           From     // Via
	Allow         []string // Allow
	Supported     []string // Supported
	AllowEvents   []string // Allow-Events
	UserAgent     string   // User-Agent
	Expires       int      // Expires
	ContentLength int      // Content-Length
}

func (m *MessageHeader) Build() string {
	var s string
	s += "Call-ID:" + m.CallID + "\n"
	s += "Contract:" + m.Contract + "\n"
	s += "CSeq:" + m.CSeq.Build() + "\n"
	s += "From:" + m.From.Build() + "\n"
	s += "To:" + m.To.Build() + "\n"
	s += "User-Agent:" + m.UserAgent + "\n"
	s += "Content-Length:0" + "\n"
	return s
}

type Via struct {
	Transport     string   // Transport
	SentByAddress string   // Sent-by Address
	SentByPort    int      // Sent-by port
	Parameter     []string // TODO: not yet implement, Branch, RPort
}

type Contract struct {
	// ContractURI
	UserPart  string
	HostPart  string
	HostPort  int
	Parameter []string
}

type CSeq struct {
	Seq    int       // Sequence Number
	Method SIPMethod // Method
}

func (c *CSeq) Build() string {
	return fmt.Sprintf("%d %s", c.Seq, string(c.Method))
}

type From struct {
	UserPart  string
	HostPart  string
	Parameter []string // TODO: not yet implement, transport=UDP
	Tag       string   // tag=
}

func (f *From) Build() string {
	// <sip:6001@100.121.131.130;transport=UDP>
	var s string
	s = fmt.Sprintf("<sip:%s@%s>", f.UserPart, f.HostPart)
	if f.Tag != "" {
		s += fmt.Sprintf("tag=%s", f.Tag)
	}
	return s
}

type To struct {
	UserPart  string
	HostPart  string
	Parameter []string // TODO: not yet implement, transport=UDP
	Tag       string   // tag=
}

func (t *To) Build() string {
	// <sip:6001@100.121.131.130;transport=UDP>
	var s string
	s = fmt.Sprintf("<sip:%s@%s>", t.UserPart, t.HostPart)
	if t.Tag != "" {
		s += fmt.Sprintf("tag=%s", t.Tag)
	}
	return s
}
