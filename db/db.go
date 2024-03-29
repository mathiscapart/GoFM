package db

import (
	"context"
	"database/sql"
	"gofm/models"
	"log"
	"time"
)

func connection() *sql.DB {
	conn, err := sql.Open("sqlite3", "database.db") // dsn
	if err != nil {
		log.Println(err)
	}

	return conn
}

func Database() {
	conn := connection()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	err := conn.PingContext(timeout)
	if err != nil {
		return
	}

	const createTable string = `CREATE TABLE IF NOT EXISTS GoFM (id INTEGER PRIMARY KEY NOT NULL, theme VARCHAR (255), artiste VARCHAR (255), song VARCHAR (255), image VARCHAR (255))`

	_, err = conn.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectTheme(theme string) *sql.Rows {
	conn := connection()
	const selectUser string = `SELECT * FROM GoFM WHERE theme = $1`
	query, err := conn.Query(selectUser, theme)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if !query.Next() {
		return nil
	}
	return query
}

func Row(theme string) []models.Song {
	selectMusic := SelectTheme(theme)
	var all []models.Song
	var song models.Song
	err := selectMusic.Scan(&song.Id, &song.Theme, &song.Artiste, &song.Song, &song.Image, &song.Radio, &song.Title)
	if err != nil {
		return nil
	}
	all = append(all, song)
	for selectMusic.Next() {
		if err := selectMusic.Scan(&song.Id, &song.Theme, &song.Artiste, &song.Song, &song.Image, &song.Radio, &song.Title); err != nil {
			log.Println(err)
		}
		all = append(all, song)
	}
	return all
}
