/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-12-12 15:38:11
 */

package goutils

import (
	"bytes"
)

// DumpInt64 can dump a int64 value to a bytes buffer.
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
func DumpInt64(val int64) []byte {
	b := &bytes.Buffer{}

	var signBit byte = 0x0
	if val < 0 {
		signBit = 0x40
		val = 0 - val
	}

	isFirstByte := true
	for val != 0 {
		var val byte
		if isFirstByte {
			// First two bits is special:
			// The first one is used to indicate wheather the next byte is included or not.
			// The second one is the sign bit uesed to indicate the number is negative or not.
			val = byte(val&0x3f) | signBit
			val >>= 6
			isFirstByte = false
		} else {
			val = byte(val & 0x7f)
			val >>= 7
		}

		// bit for next byte
		if val != 0 {
			val |= 0x80
		}
		b.WriteByte(val)
	}

	return b.Bytes()
}

func LoadInt64(buf []byte) int64 {
	bytesCount := 1
	for i := 0; buf[i]&0x80 == 0x80; i++ {
		bytesCount++
	}

	var v int64 = 0
	var signBit byte
	for i := 0; i < bytesCount; i++ {
		if i == 0 {
			signBit = byte(0x40 & buf[i])
			v |= int64(buf[i] & 0x3f)
		} else {
			bitsMove := 6 + (7 * (i - 1))
			v |= int64(buf[i]&0x7f) << uint(bitsMove)
		}
	}

	if signBit == 0x40 {
		v = -1 * v
	}

	return v
}
