// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan []byte
}

func (c *Client) check(hub *Hub) {
	defer func() {
		switch hub.radio {
		case "rap":
			opsProcessedRap.Dec()
		case "pop":
			opsProcessedPop.Dec()
		case "rock":
			opsProcessedRock.Dec()
		case "slow":
			opsProcessedSlow.Dec()
		case "gen":
			opsProcessedGen.Dec()
		}
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	switch hub.radio {
	case "rap":
		opsProcessedRap.Inc()
	case "pop":
		opsProcessedPop.Inc()
	case "rock":
		opsProcessedRock.Inc()
	case "slow":
		opsProcessedSlow.Inc()
	case "gen":
		opsProcessedGen.Inc()
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 1000)}
	client.hub.register <- client
	client.check(hub)
}
