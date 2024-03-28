package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Song struct {
	Id      int    `json:"id"`
	Theme   string `json:"theme"`
	Artiste string `json:"artiste"`
	Song    string `json:"song"`
	Image   string `json:"image"`
	Radio   string `json:"radio"`
	Title   string `json:"title"`
}

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
	conn.PingContext(timeout)

	const createTable string = `CREATE TABLE IF NOT EXISTS GoFM (id INTEGER PRIMARY KEY NOT NULL, theme VARCHAR (255), artiste VARCHAR (255), song VARCHAR (255), image VARCHAR (255))`

	_, err := conn.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

/*func InsertUser(value Song) {
	conn := connection()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	conn.PingContext(timeout)

	const insertUser string = `INSERT INTO user (name, email, password) VALUES ($1, $2, $3)`

	_, err := conn.Exec(insertUser, value.Name, value.Mail, value.Password)
	if err != nil {
		log.Fatal(err)
	}
}*/

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

func Row(theme string) []Song {
	selectRap := SelectTheme(theme)
	var all []Song
	var song Song
	selectRap.Scan(&song.Id, &song.Theme, &song.Artiste, &song.Song, &song.Image, &song.Radio, &song.Title)
	all = append(all, song)
	for selectRap.Next() {
		if err := selectRap.Scan(&song.Id, &song.Theme, &song.Artiste, &song.Song, &song.Image, &song.Radio, &song.Title); err != nil {
			log.Println(err)
		}
		all = append(all, song)
	}
	return all
}
