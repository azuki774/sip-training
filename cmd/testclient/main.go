package main

import (
	"azuki774/sip-training/internal/model"
	"log/slog"
	"net"
	"os"
)

const udpAddr = "100.121.131.130:5060"
const sipSrvAddr = "100.121.131.130"
const sipSrvPort = 5060

var sipUser = "7001"

func main() {
	slog.Info("start")

	conn, err := net.Dial("udp", udpAddr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	slog.Info("source IP", localAddr)

	reqLine := model.RequestLine{
		Method:     model.SIPMethodRegister,
		RequestURI: "sip:100.121.131.130",
		Transport:  "UDP SIP/2.0",
	}

	reqHeader := model.MessageHeader{
		CallID: "1234567890",
		Contract: model.Contract{
			UserPart: sipUser,
			HostPart: "100.85.93.61", // TODO: 取得方法を考える
			HostPort: 33333,          // TODO:  取得方法を考える 違う
		},
		CSeq: model.CSeq{
			Seq:    1,
			Method: model.SIPMethodRegister,
		},
		From: model.From{
			UserPart: sipUser,
			HostPart: sipSrvAddr,
		},
		To: model.To{
			UserPart: sipUser,
			HostPart: sipSrvAddr,
		},
		Via: model.Via{
			Transport:     "UDP",
			SentByAddress: "100.85.93.61", // TODO: 取得方法を考える
			SentByPort:    33333,          // TODO:  取得方法を考える 違う
		},
		UserAgent: "YRP yabasugi Call Client",
	}

	req := append([]byte(reqLine.Build()), []byte(reqHeader.Build())...)
	_, err = conn.Write(req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}
