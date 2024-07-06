package main

type bencodeBuffer struct {
	byteBuffer []byte
	index      int64
}

func (buf *bencodeBuffer) readByte() byte {
	res := buf.byteBuffer[buf.index]
	buf.index += 1
	return res
}

func (buf *bencodeBuffer) unreadByte() {
	buf.index -= 1
}

func (buf *bencodeBuffer) readUntil() (result []byte, err error) {

}
