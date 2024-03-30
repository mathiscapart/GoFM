package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gofm/db"
	"gofm/models"
	"io"
	"log"
	"os"
	"time"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	radio string

	register chan *Client

	unregister chan *Client
}

func newHub(radio string) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		radio:      radio,
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		}
	}
}

func (h *Hub) sendMusique() {
	if h.radio == "rap" || h.radio == "pop" || h.radio == "rock" || h.radio == "slow" || h.radio == "gen" {
		for {
			h.sendJSON()
		}
	} else {
		log.Println("Error theme song")
	}

}

func (h *Hub) sendJSON() {
	var songs []models.Song
	clock := time.Now().Hour()
	clockMinute := time.Now().Minute()
	fmt.Println("Start Radio", h.radio)
	if clock == 7 && clockMinute >= 0 && clockMinute < 2 {
		songs = db.Row("radio")
	} else if h.radio == "gen" {
		if 0 <= clock && clock < 6 {
			songs = db.Row("rap")
		} else if 6 <= clock && clock < 10 {
			songs = db.Row("rock")
		} else if 10 <= clock && clock < 18 {
			songs = db.Row("pop")
		} else if 18 <= clock && clock < 24 {
			songs = db.Row("slow")
		}
	} else {
		songs = db.Row(h.radio)
	}
	for _, song := range songs {
		clock = time.Now().Hour()
		clockMinute = time.Now().Minute()
		fmt.Println("La Radio en cours : ", song.Radio, " à "+"Heure : ", clock, ", Minute : ", clockMinute)
		if clock == 7 && clockMinute >= 0 && clockMinute < 2 && song.Radio != "GoFM" {
			fmt.Println(song.Radio)
			fmt.Println("break radio for Horoscope")
			break
		}
		file, err := os.Open("mp3/" + song.Song)
		if err != nil {
			log.Printf("error opening file: %v", err)
			return
		}

		if err != nil {

			log.Println(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(file)

		secondBit := sizeFile("mp3/"+song.Song) / int64(longueurFile("mp3/"+song.Song))

		buffer := make([]byte, secondBit*2+300)
		for {
			n, err := file.Read(buffer)
			if err != nil {
				if errors.Is(io.EOF, err) {
					log.Println("End file")
					break
				}
				log.Printf("error reading file: %v", err)
				break
			}
			if n == 0 {
				break
			}
			str := base64.StdEncoding.EncodeToString(buffer[:n])

			musique := map[string]interface{}{
				"id":        song.Id,
				"partition": str,
				"image":     song.Image,
				"theme":     song.Theme,
				"artiste":   song.Artiste,
				"radio":     song.Radio,
				"title":     song.Title,
			}

			for c := range h.clients {
				go func(c *Client) {
					err = c.conn.WriteJSON(musique)
				}(c)
			}

			if err != nil {
				return
			}
			time.Sleep((2 * time.Second) - 10)
		}
	}
}
