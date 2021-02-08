package icap

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"
)

const (
	// const string to format time for HTTP call
	timeFormat = "Mon Jan 02 15:04:05 2006"
)

func sendReqmod(hostPort, service string, conn *net.TCPConn) {
	req := bytes.NewBuffer(nil)
	req.WriteString(fmt.Sprintf("REQMOD icap://%s/%s ICAP/1.0\r\n", hostPort, service))
	req.WriteString(fmt.Sprintf("Host: %s\r\n", hostPort))
	req.WriteString("User-Agent: C-ICAP-Client-Library/0.5.6\r\n")
	req.WriteString(fmt.Sprintf("Encapsulated: req-hdr=0, null-body=84\r\n"))
	req.WriteString("\r\n")

	reqLen := req.Len()

	var n int64
	n, err := io.Copy(conn, req)
	if err != nil {
		return
	}
	if n != int64(reqLen) {
		err = errors.New("partial write of reqmod request")
		return
	}
}

func sendHTTPReq(conn *net.TCPConn) {
	req := bytes.NewBuffer(nil)
	req.WriteString("GET any HTTP/1.0\r\n")
	req.WriteString(fmt.Sprintf("Date: %s\r\n", time.Now().UTC().Format(timeFormat)))
	req.WriteString("User-Agent: C-ICAP-Client/x.xx\r\n")
	req.WriteString("\r\n")

	reqLen := req.Len()

	var n int64
	n, err := io.Copy(conn, req)
	if err != nil {
		return
	}
	if n != int64(reqLen) {
		err = errors.New("partial write of http request")
		return
	}
}

func sendOptionsReq(hostPort string, conn *net.TCPConn) {
	req := bytes.NewBuffer(nil)
	req.WriteString("OPTIONS icap://%s/gw_rebuild ICAP/1.0\r\n")
	req.WriteString(fmt.Sprintf("Host: %s\r\n", hostPort))
	req.WriteString("User-Agent: C-ICAP-Client-Library/0.5.6\r\n")
	req.WriteString(fmt.Sprintf("Encapsulated: null-body=0\r\n"))
	req.WriteString("\r\n")

	reqLen := req.Len()

	var n int64
	n, err := io.Copy(conn, req)
	if err != nil {
		return
	}
	if n != int64(reqLen) {
		err = errors.New("partial write of options request")
		return
	}
}

func collectStatistics(host, port, service string) (res []byte, err error) {
	hostPort := net.JoinHostPort(host, port)
	var addr *net.TCPAddr
	if addr, err = net.ResolveTCPAddr("tcp", hostPort); err != nil {
		return
	}

	var conn *net.TCPConn
	if conn, err = net.DialTCP("tcp", nil, addr); err != nil {
		return
	}
	defer conn.Close()

	fmt.Println("Sending REQMOD request")
	sendReqmod(hostPort, service, conn)
	fmt.Println("Sent REQMOD request")

	fmt.Println("Sending HTTP request")
	sendHTTPReq(conn)
	fmt.Println("Sent HTTP request")

	fmt.Println("Closing connection write")
	err = conn.CloseWrite()
	if err != nil {
		return
	}
	fmt.Println("Closed connection write")

	fmt.Println("Reading response from the connection")
	res, err = ioutil.ReadAll(conn)
	fmt.Println("Read response from the connection")

	return
}

func CheckHealth(host, port string) (alive bool, err error) {
	alive = true

	hostPort := net.JoinHostPort(host, port)
	var addr *net.TCPAddr
	if addr, err = net.ResolveTCPAddr("tcp", hostPort); err != nil {
		return
	}

	var conn *net.TCPConn
	if conn, err = net.DialTCP("tcp", nil, addr); err != nil {
		return
	}
	defer conn.Close()

	fmt.Println("Sending OPTIONS request")
	sendOptionsReq(hostPort, conn)
	fmt.Println("Sent OPTIONS request")

	err = conn.CloseWrite()
	if err != nil {
		return
	}

	res, err := ioutil.ReadAll(conn)
	if err != nil || res == nil {
		alive = false
	}

	icapCode := ParseIcapHeader(res)
	if icapCode != 200 {
		alive = false
	}

	return
}
