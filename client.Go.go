package main

import "net"

type peerNode struct {
	addr net.Addr
	hash [20]byte
}
