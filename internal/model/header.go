package model

import (
	"fmt"
	"strconv"
)

type MessageHeader struct {
	CallID        string   // Call-ID
	Contact       Contact  // Contact
	CSeq          CSeq     // CSeq
	From          From     // From
	To            To       // To
	Via           Via      // Via
	Allow         []string // Allow
	Supported     []string // Supported
	AllowEvents   []string // Allow-Events
	UserAgent     string   // User-Agent
	MaxForwards   int      // Max-Forwards
	Expires       int      // Expires
	ContentLength int      // Content-Length
}

func (m *MessageHeader) Build() string {
	var s string
	s += ColonFieldBuild("Call-ID", m.CallID)
	s += ColonFieldBuild("Contact", m.Contact.Build())
	s += ColonFieldBuild("CSeq", m.CSeq.Build())
	s += ColonFieldBuild("From", m.From.Build())
	s += ColonFieldBuild("To", m.To.Build())
	s += ColonFieldBuild("Via", m.Via.Build())
	s += ColonFieldBuild("Max-Forwards", strconv.Itoa(m.MaxForwards))
	s += ColonFieldBuild("User-Agent", m.UserAgent)
	s += ColonFieldBuild("Expires", strconv.Itoa(m.Expires))
	s += ColonFieldBuild("Content-Length", strconv.Itoa(0))
	s += "Content-Length: 0" + "\r\n"
	s += "\r\n" // ヘッダ終わりは空行
	return s
}

type Contact struct {
	// ContractURI
	UserPart  string
	HostPart  string
	HostPort  int
	Parameter []string
}

func (c *Contact) Build() string {
	// <sip:6001@100.121.131.130;transport=UDP>
	s := fmt.Sprintf("<sip:%s@%s:%d", c.UserPart, c.HostPart, c.HostPort)
	s += stringsOptBuild(c.Parameter) // Contract フィールドはURIオプショなので < > の間に入る
	s += ">"
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
