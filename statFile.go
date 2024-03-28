package main

import (
	"fmt"
	"github.com/tcolgate/mp3"
	"io"
	"log"
	"os"
)

func sizeFile(filepath string) int64 {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
	}

	return fi.Size()
}

func longueurFile(filepath string) float64 {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := 0.0
	d := mp3.NewDecoder(file)
	var f mp3.Frame
	skipped := 0

	for {
		err := d.Decode(&f, &skipped)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return 0
		}
		t = t + f.Duration().Seconds()
	}
	return t
}
