/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-12-12 15:38:11
 */

package goutils

import (
	"bufio"
	"bytes"
	"io"
)

// WriteInt64 can write a int64 value to io.Writer.
// Variable length encoding is used, and the rule is as follows:
//	1) First bit of each byte is special to indicate wheather the following byte
//	   is included or not.
//	2) Second bit of the first byte is the sign bit used to specify the number is
//	   negative or not.
// The order of byte sequence is little endian. For example, '-1234' will be dumped
// as follows:
//	binary sequence for '-1234' is '-100 11010010'.
//	binary sequence dumped consisted of two bytes:
//		- '11 010010'
//		- '0 0010011'
func WriteInt64(val int64, w io.Writer) error {
	var signBit byte = 0x0
	if val < 0 {
		signBit = 0x40
		val = 0 - val
	}

	out := bufio.NewWriter(w)

	// '0' is a special case
	if val == 0 {
		if err := out.WriteByte(byte(0x00)); err != nil {
			return err
		}
		if err := out.Flush(); err != nil {
			return err
		}

		return nil
	}

	// not '0'
	isFirstByte := true
	for val != 0 {
		var b byte
		if isFirstByte {
			// First two bits is special:
			// The first one is used to indicate wheather the next byte is included or not.
			// The second one is the sign bit uesed to indicate the number is negative or not.
			b = byte(val&0x3f) | signBit
			val >>= 6
			isFirstByte = false
		} else {
			b = byte(val & 0x7f)
			val >>= 7
		}

		// bit for next byte
		if val != 0 {
			b |= 0x80
		}

		if err := out.WriteByte(b); err != nil {
			return err
		}
	}

	if err := out.Flush(); err != nil {
		return err
	}

	return nil
}

func DumpInt64(val int64) []byte {
	b := &bytes.Buffer{}

	WriteInt64(val, b)

	return b.Bytes()
}

func ReadInt64(r io.Reader) (int64, error) {
	var v int64 = 0
	var signBit byte
	var b byte = 0x80
	var err error

	in := bufio.NewReader(r)
	for idx := 0; b&0x80 != 0; idx++ {
		b, err = in.ReadByte()
		if err != nil && err != io.EOF {
			return -1, err
		} else if err == io.EOF {
			return -1, NewErr("failed to read end byte, no enough bytes")
		}

		if idx == 0 {
			signBit = byte(0x40 & b)
			v |= int64(b & 0x3f)
		} else {
			bitsMove := 6 + (7 * (idx - 1))
			v |= int64(b&0x7f) << uint(bitsMove)
		}
	}

	if signBit == 0x40 {
		v = -1 * v
	}

	return v, nil
}

func LoadInt64(buf []byte) (int64, error) {
	b := bytes.NewBuffer(buf)

	val, err := ReadInt64(b)
	if err != nil {
		return -1, err
	}

	return val, nil
}
