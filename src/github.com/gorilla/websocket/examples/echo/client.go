// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	//"os"
	//"os/signal"
	//"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.146.128:7070", "http service address")

//var addr = flag.String("addr", "10.19.85.175:7070", "http service address")

// var addr = flag.String("addr", "localhost:7070", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/network/netInfo/sysruntimeinfolist"}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/network/netInfo/cpuinfolist"}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/network/netInfo/memoryinfolist"}
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/network/netInfo/interfaceflowlist"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	mss := `{"ip":"192.168.146.128","ethn":"ens33","timeinterval":"5"}`
	//mss := `{"ip":"localhost","ethn":"enp4s0f1","timeinterval":"5"}`
	//mss := `{"ip":"localhost","ethn":"eno3","timeinterval":"5"}`
	//err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	err1 := c.WriteMessage(websocket.TextMessage, []byte(mss))
	// err1 := c.WriteJSON(mss)
	if err1 != nil {
		log.Println("write:", err1)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv from server: %s", message)

		log.Println("interrupt")
		// Cleanly close the connection by sending a close message and then
		// waiting (with timeout) for the server to close the connection.
		erro := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if erro != nil {
			log.Println("write close:", erro)
			return
		}
	}
}
