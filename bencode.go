package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

func bDecode(reader *bufio.Reader) (res interface{}, err error) {
	res, err = decode(reader)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func decode(reader *bufio.Reader) (result interface{}, err error) {
	token, err := reader.ReadByte()

	if err != nil {
		return nil, err
	}

	switch token {
	case 'i':
		intBuffer, err := readUntil(reader, 'e')
		if err != nil {
			return nil, err
		}
		intBuffer = intBuffer[:len(intBuffer)-1] //removing the last character

		integer, err := strconv.ParseInt(string(intBuffer), 10, 64) //bencode uses 64-bit ints

		if err != nil {
			return nil, err
		}

		return integer, nil
	case 'l':
		//parsing list
		blist := []interface{}{}

		for {
			ch, err := reader.ReadByte()

			if err == nil {
				if ch == 'e' {
					return blist, nil
				} else {
					reader.UnreadByte()
				}
			}

			val, err := decode(reader)
			if err != nil {
				return nil, err
			}

			blist = append(blist, val)
		}
	case 'd':

		dict := map[string]interface{}{}
		for {
			c, err := reader.ReadByte()
			if err == nil {
				if c == 'e' {
					return dict, nil
				} else {
					reader.UnreadByte()
				}
			}

			val, err := decode(reader)

			if err != nil {
				return nil, err
			}

			key, ok := val.(string)
			if !ok {
				return nil, errors.New("bencode: key is not a string.")
			}

			val, err = decode(reader)

			if err != nil {
				return nil, err
			}

			dict[key] = val
		}
	default:
		reader.UnreadByte()
		strlen, err := readUntil(reader, ':')

		strlen = strlen[:len(strlen)-1]

		strl, err := strconv.ParseInt(string(strlen), 10, 64)

		buffer := make([]byte, strl)

		_, err = readAtLeast(reader, buffer, int(strl))

		return string(buffer), err
	}
}

func readUntil(reader *bufio.Reader, delimiter byte) (result []byte, err error) {
	bufferedReader := reader.Buffered()

	var buf []byte

	var er error

	buf, er = reader.Peek(bufferedReader)
	if er != nil {
		return nil, er
	}

	if i := bytes.IndexByte(buf, delimiter); i >= 0 {
		return reader.ReadSlice(delimiter)
	}
	return reader.ReadBytes(delimiter)

}

func readAtLeast(reader *bufio.Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, io.ErrShortBuffer
	}
	for n < min && err == nil {
		var bytesRead int
		bytesRead, err = reader.Read(buf[bytesRead:])
		n += bytesRead
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return
}
