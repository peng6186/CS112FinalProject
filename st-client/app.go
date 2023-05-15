package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	. "libs"
	"log"
	"net"
	"sync"
)

var proxyRule = "direct"
var appCtx context.Context

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	appCtx = ctx
}

func UpdateConnList(text string) {
	runtime.EventsEmit(appCtx, "updateConnList", text)
}

func handleProxyRequest(localClient *net.TCPConn, proxyServerAddr *net.TCPAddr) {
	fmt.Println("handleProxyReq")

	request := ResolveClientRequest(localClient)

	// Update frontend
	UpdateConnList(request.DSTDOMAIN)
	log.Println(localClient.RemoteAddr(), request.DSTDOMAIN, request.DSTADDR, request.DSTPORT)

	// Connect to the proxy server
	proxyServerConn, err := net.DialTCP("tcp", nil, proxyServerAddr)
	if err != nil {
		log.Print(localClient.RemoteAddr(), err)
		return
	}
	defer proxyServerConn.Close()

	// Send Socks5 handshake to proxy server
	protocolVersion := ProtocolVersion{}
	err = protocolVersion.SendHandshake(proxyServerConn)
	if err != nil {
		return
	}

	// Send auth request to proxy server
	// Currently only support no auth
	buff := make([]byte, 255)
	_, err = proxyServerConn.Read(buff)
	if err != nil {
		return
	}
	err = HandleHandshakeResponse(buff)
	if err != nil {
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		Copy(localClient, proxyServerConn)
	}()

	go func() {
		defer wg.Done()
		Copy(proxyServerConn, localClient)
	}()
	wg.Wait()

}

func handleDirectRequest(localClient *net.TCPConn, serverAddr *net.TCPAddr) {
	fmt.Println("handleDirectReq")

	request := ResolveClientRequest(localClient)

	// Update frontend
	UpdateConnList(request.DSTDOMAIN)
	log.Println(localClient.RemoteAddr(), request.DSTDOMAIN, request.DSTADDR, request.DSTPORT)

	// Connect to the remote server
	dstServer, err := net.DialTCP("tcp", nil, request.RAWADDR)
	if err != nil {
		log.Print(localClient.RemoteAddr(), err)
		return
	}
	defer dstServer.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		Copy(localClient, dstServer)
	}()

	go func() {
		defer wg.Done()
		Copy(dstServer, localClient)
	}()
	wg.Wait()

}

func ResolveClientRequest(localClient *net.TCPConn) (request Socks5Resolution) {
	buff := make([]byte, 255)

	var proto ProtocolVersion
	n, err := localClient.Read(buff)
	resp, err := proto.HandleHandshake(buff[0:n])
	localClient.Write(resp)
	if err != nil {
		log.Print(localClient.RemoteAddr(), err)
		return
	}

	// Resolve remote address
	n, err = localClient.Read(buff)
	resp, err = request.LSTRequest(buff[0:n])
	_, err = localClient.Write(resp)
	if err != nil {
		log.Print(localClient.RemoteAddr(), err)
		return
	}

	return request
}

func (a *App) SetProxyMode(mode string) {
	if mode == "proxy" || mode == "direct" {
		proxyRule = mode
	}
}

var listener *net.TCPListener

func (a *App) Stop() {
	err := listener.Close()
	if err != nil {
		log.Print(err)
	}
}

func (a *App) Connect(listenAddrString string, serverAddrString string) {
	// st-proxy
	serverAddr, err := net.ResolveTCPAddr("tcp", serverAddrString)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connecting to proxy: %s ....", serverAddrString)

	// st-client
	listenAddr, err := net.ResolveTCPAddr("tcp", listenAddrString)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening: %s ", listenAddrString)

	listener, err = net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		localClient, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			return
		}
		if proxyRule == "proxy" {
			go handleProxyRequest(localClient, serverAddr)
		} else if proxyRule == "direct" {
			go handleDirectRequest(localClient, serverAddr)
		}
	}
}
