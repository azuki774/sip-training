package main

import (
	"azuki774/sip-training/internal/model"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
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

	localAddrPort := conn.LocalAddr().(*net.UDPAddr).String()
	slog.Info("get source info", "source", localAddrPort)
	localAddr, localPortStr, err := net.SplitHostPort(localAddrPort)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	localPort, err := strconv.Atoi(localPortStr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	reqLine := model.RequestLine{
		Method:     model.SIPMethodRegister,
		RequestURI: fmt.Sprintf("sip:%v", localAddr),
		Transport:  "UDP SIP/2.0",
	}

	reqHeader := model.MessageHeader{
		CallID: "1234567890",
		Contact: model.Contact{
			UserPart:  sipUser,
			HostPart:  localAddr,
			HostPort:  localPort,
			Parameter: []string{"", "transport=UDP", "rinstance=1111111111111111"}, // 先頭に;をつけるため "" がいる
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
			SentByAddress: localAddr,
			SentByPort:    localPort,
			Parameter:     []string{"", "branch=zzzzzzzzzzzzz", "rport"}, // 先頭に;をつけるため "" がいる
		},
		MaxForwards: 70,
		Expires:     60,
		UserAgent:   "YRP yabasugi Call Client",
	}

	req := append([]byte(reqLine.Build()), []byte(reqHeader.Build())...)
	_, err = conn.Write(req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// catch response
	// buf := make([]byte, 8192)
	// n, _ := conn.Read(buf)
	// resStr := string(buf[:n])

}
