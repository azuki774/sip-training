package main

import (
	"azuki774/sip-training/internal/model"
	"log/slog"
	"net"
	"os"
)

func main() {
	slog.Info("start")
	udpAddr := "100.121.131.130:5060"

	conn, err := net.Dial("udp", udpAddr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	reqLine := model.RequestLine{
		Method:     model.SIPMethodRegister,
		RequestURI: "sip:100.121.131.130",
		Transport:  "UDP SIP/2.0",
	}
	_, err = conn.Write([]byte(reqLine.Build()))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}
