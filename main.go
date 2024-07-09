package main

func main() {
	torrentLink := "/Users/fernando/Downloads/debian-12.6.0-amd64-DVD-1.iso.torrent"

	res, err := parseTorrent(torrentLink)

	print("he")
	print(res, err)

}
