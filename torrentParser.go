package main

import (
	"bufio"
	"github.com/jackpal/bencode-go"
	"os"
)

type torrentInfo struct {
	length    int64
	name      string
	piece_len string

	pieces string
}

type torrentFile struct {
	announce     string
	comment      string
	createdBy    string
	creationDate int64
	info         torrentInfo
}

func parseTorrent(filepath string) (res interface{}, err error) {

	f, err := os.Open(filepath)
	fi, err := f.Stat()

	if err != nil {
		return nil, err
	}

	filereader := bufio.NewReader(f)
	filereader = bufio.NewReaderSize(filereader, int(fi.Size()))
	bencode.Unmarshal(filereader, &res)

	return res, nil
}
