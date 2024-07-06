package main

import (
	"bufio"
	"log"
	"net/http"
)



func main() {
	torrentLink := "https://cdimage.debian.org/debian-cd/current/amd64/bt-dvd/debian-12.6.0-amd64-DVD-1.iso.torrent"
	resp, err := http.Get(torrentLink)

	if err != nil {
		log.Fatalln(err)
	}

	var reader = bufio.NewReader(resp.Body)
	bufio.NewReaderSize(reader, 1 << )
	var decodedObj, er2 = bDecode(reader)
	if er2 != nil {
		log.Fatalln(er2)
	}
	print(decodedObj)

}
