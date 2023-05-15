package main

import (
	. "libs"
	"log"
	"net"
	"os"
	"sync"
)

func main() {

	args := os.Args[1:]
	var port string
	if len(args) != 1 {
		log.Println("No port specified, using default port 8888")
		port = "8888"
	} else {
		port = args[0]
	}

	listenAddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port: %s ", port)

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleClientRequest(conn)
	}
}

func handleClientRequest(client *net.TCPConn) {
	if client == nil {
		return
	}
	defer client.Close()

	buff := make([]byte, 255)

	// 认证协商
	var proto ProtocolVersion
	//n, err := auth.DecodeRead(client, buff) //解密
	n, err := client.Read(buff)
	resp, err := proto.HandleHandshake(buff[0:n])
	//auth.EncodeWrite(client, resp) //加密
	client.Write(resp)
	if err != nil {
		log.Print(client.RemoteAddr(), err)
		return
	}

	// 获取客户端代理的请求
	var request Socks5Resolution
	//n, err = auth.DecodeRead(client, buff)
	n, err = client.Read(buff)
	resp, err = request.LSTRequest(buff[0:n])
	//auth.EncodeWrite(client, resp)
	_, err = client.Write(resp)
	if err != nil {
		log.Print(client.RemoteAddr(), err)
		return
	}

	log.Println(client.RemoteAddr(), request.DSTDOMAIN, request.DSTADDR, request.DSTPORT)

	// 连接真正的远程服务
	dstServer, err := net.DialTCP("tcp", nil, request.RAWADDR)
	if err != nil {
		log.Print(client.RemoteAddr(), err)
		return
	}
	defer dstServer.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// 本地的内容copy到远程端
	go func() {
		defer wg.Done()
		Copy(client, dstServer)
	}()

	// 远程得到的内容copy到源地址
	go func() {
		defer wg.Done()
		Copy(dstServer, client)
	}()
	wg.Wait()

}
