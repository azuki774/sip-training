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

type MessageHeader struct {
	CallID        string   // Call-ID
	Contract      Contract // Contract
	CSeq          CSeq     // CSeq
	From          From     // From
	To            To       // To
	Via           Via      // Via
	Allow         []string // Allow
	Supported     []string // Supported
	AllowEvents   []string // Allow-Events
	MaxForwards   int      // Max-Forwards
	UserAgent     string   // User-Agent
	Expires       int      // Expires
	ContentLength int      // Content-Length
}

func (m *MessageHeader) Build() string {
	var s string
	s += "Call-ID:" + m.CallID + "\n"
	s += "Contract:" + m.Contract.Build() + "\n"
	s += "CSeq:" + m.CSeq.Build() + "\n"
	s += "From:" + m.From.Build() + "\n"
	s += "To:" + m.To.Build() + "\n"
	s += "Via:" + m.Via.Build() + "\n"
	s += fmt.Sprintf("Max-Forwards:%d", m.MaxForwards) + "\n"
	s += "User-Agent:" + m.UserAgent + "\n"
	s += "Content-Length:0" + "\n"
	return s
}

type Via struct {
	Transport     string   // Transport
	SentByAddress string   // Sent-by Address
	SentByPort    int      // Sent-by port
	Parameter     []string // Branch, RPort
}

func (v *Via) Build() string {
	var s string
	s += "SIP/2.0/"
	s += v.Transport
	s += " " // SP
	s += fmt.Sprintf("%s:%d", v.SentByAddress, v.SentByPort)
	s += stringsOptBuild(v.Parameter)
	return s
}

type Contract struct {
	// ContractURI
	UserPart  string
	HostPart  string
	HostPort  int
	Parameter []string
}

func (c *Contract) Build() string {
	// <sip:6001@100.121.131.130;transport=UDP>
	s := fmt.Sprintf("<sip:%s@%s:%d>", c.UserPart, c.HostPart, c.HostPort)
	return s
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
