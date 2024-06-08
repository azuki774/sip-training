package main

import (
	"azuki774/sip-training/internal/digest"
	"azuki774/sip-training/internal/model"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"
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
		RequestURI: fmt.Sprintf("sip:%v;transport=UDP", sipSrvAddr),
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

	fmt.Println("------ Req ------")
	fmt.Println(string(req))

	// catch response
	buf := make([]byte, 8192)
	n, _ := conn.Read(buf)
	resStr := string(buf[:n])
	fmt.Println("------ Res ------")
	fmt.Println(string(resStr))

	f := model.ParseColonField(string(resStr))

	wwwauthStr := f["WWW-Authenticate"]              // Digest realm="asterisk",nonce="1717835497/fc426b527a93603ba551f29ff77a5bc1",opaque="5566ebd26176cc4c",algorithm=MD5,qop="auth"
	wwwauthElms := strings.Split(wwwauthStr, " ")[1] // realm="asterisk",nonce="1717835497/fc426b527a93603ba551f29ff77a5bc1",opaque="5566ebd26176cc4c",algorithm=MD5,qop="auth"
	wwwauthElmsMap := model.ParseCommaEqualField(wwwauthElms)
	wwwAuthField := digest.WWWAuthenticate{
		AuthenticateScheme: "Digest",
		UserName:           sipUser,                                                 // username
		Realm:              strings.Replace(wwwauthElmsMap["realm"], "\"", "", -1),  // (unq) realm
		NonceValue:         strings.Replace(wwwauthElmsMap["nonce"], "\"", "", -1),  // (unq) nonce
		OpaqueValue:        strings.Replace(wwwauthElmsMap["opaque"], "\"", "", -1), // (unq) opaque
		Algorithm:          wwwauthElmsMap["algorithm"],                             // algorithm
		QOP:                strings.Replace(wwwauthElmsMap["qop"], "\"", "", -1),    // qop
		URI:                reqLine.RequestURI,                                      // uri
		NonceCount:         "00000001",                                              // nc
		//
		CNonse:   "1b043ccbff85cf2b8cac03898ebb6267", // cnonce // TODO: fix
		Response: "",                                 // response
	}
	wwwAuthField.ComputeResponse(string(model.SIPMethodRegister), "supersecret")

	reqHeader.CSeq.Seq += 1
	reqHeader.WWWAuthenticate = wwwAuthField
	fmt.Println(reqHeader.WWWAuthenticate)
	req = append([]byte(reqLine.Build()), []byte(reqHeader.Build())...)
	_, err = conn.Write(req) // Auth REGISTER
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println("------ Req ------")
	fmt.Println(string(req))

	// catch response
	buf = make([]byte, 8192)
	n, _ = conn.Read(buf)
	resStr = string(buf[:n])
	fmt.Println("------ Res ------")
	fmt.Println(string(resStr))

}
