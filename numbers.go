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

const (
	BIG_ENDIAN = iota
	LITTLE_ENDIAN
)

var (
	// BytesOrder is the byte order of the matchine
	BytesOrder int
)

func intToBytes(v int64, nBytes int, bytesOrder int) []byte {
	b := make([]byte, nBytes)
	var s, inc int
	if bytesOrder == BIG_ENDIAN {
		s, inc = nBytes-1, -1
	} else {
		s, inc = 0, 1
	}

	for i := 0; i < nBytes; i++ {
		b[i] = byte((v >> uint(s*8)) & 0xff)
		s += inc
	}

	return b
}

func BytesToInt(b []byte, bytesOrder int) int64 {
	var v int64 = 0
	if bytesOrder == BIG_ENDIAN {
		for i := 0; i < len(b); i++ {
			v = (v << 8) | int64(b[i])
		}
	} else {
		for i := len(b) - 1; i >= 0; i-- {
			v = (v << 8) | int64(b[i])
		}
	}

	return v
}

// write int64 as little endian
func WriteFixedInt64(i int64, w io.Writer) error {
	o := intToBytes(i, 8, BytesOrder)

	if err := WriteBuffer(nil, w, o); err != nil {
		return err
	}

	return nil
}

func WriteFixedInt32(i int32, w io.Writer) error {
	o := intToBytes(int64(i), 4, BytesOrder)

	if err := WriteBuffer(nil, w, o); err != nil {
		return err
	}

	return nil
}

func WriteFixedInt16(i int16, w io.Writer) error {
	o := intToBytes(int64(i), 2, BytesOrder)

	if err := WriteBuffer(nil, w, o); err != nil {
		return err
	}

	return nil
}

func ReadFixedInt64(r io.Reader) (int64, error) {
	in := make([]byte, 8)

	if err := ReadBuffer(nil, r, in); err != nil {
		return -1, err
	}

	v := BytesToInt(in, BytesOrder)

	return v, nil
}

func ReadFixedInt32(r io.Reader) (int32, error) {
	in := make([]byte, 4)

	if err := ReadBuffer(nil, r, in); err != nil {
		return -1, err
	}

	v := BytesToInt(in, BytesOrder)

	return int32(v), nil
}

func ReadFixedInt16(r io.Reader) (int16, error) {
	in := make([]byte, 2)

	if err := ReadBuffer(nil, r, in); err != nil {
		return -1, err
	}

	v := BytesToInt(in, BytesOrder)

	return int16(v), nil
}

func DumpFixedInt64(i int64) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := WriteFixedInt64(i, b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func DumpFixedInt32(i int32) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := WriteFixedInt32(i, b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func DumpFixedInt16(i int16) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := WriteFixedInt16(i, b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func LoadFixedInt64(buf []byte) (int64, error) {
	b := bytes.NewBuffer(buf)

	v, err := ReadFixedInt64(b)
	if err != nil {
		return -1, err
	}

	return v, nil
}

func LoadFixedInt32(buf []byte) (int32, error) {
	b := bytes.NewBuffer(buf)

	v, err := ReadFixedInt32(b)
	if err != nil {
		return -1, err
	}

	return v, nil
}

func LoadFixedInt16(buf []byte) (int16, error) {
	b := bytes.NewBuffer(buf)

	v, err := ReadFixedInt16(b)
	if err != nil {
		return -1, err
	}

	return v, nil
}

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

func DumpInt64(val int64) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := WriteInt64(val, b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
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

func init() {
	n := 0x0102
	if n&0xff == 0x02 {
		BytesOrder = BIG_ENDIAN
	} else {
		BytesOrder = LITTLE_ENDIAN
	}
}
