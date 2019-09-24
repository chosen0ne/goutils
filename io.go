/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 15:56:36
 */

package goutils

import (
	"io"
)

// WriteBuffer writes the whole buffer to the writer 'w'.
// It returns an error when failed to write or a partial write.
func WriteBuffer(l Logger, w io.Writer, b []byte) error {
	if w == nil || b == nil {
		return NewErr("invalid params")
	}

	if n, err := w.Write(b); err != nil {
		if l != nil {
			l.Exception(err, "failed to write")
		}
		return err
	} else if n != len(b) {
		if l != nil {
			l.Error("failed to write a whole buffer, buf len: %d, write: %d", len(b), n)
		}
		return NewErr("failed to write a whole buffer, buf len: %d, write: %d", len(b), n)
	}

	return nil
}

// ReadBuffer read a whole buffer from the Reader 'r'.
// It reaturns an error when failed to read or a partial read.
func ReadBuffer(l Logger, r io.Reader, b []byte) error {
	if r == nil || b == nil {
		return NewErr("invalid params")
	}

	if n, err := io.ReadFull(r, b); err != nil {
		if l != nil {
			l.Exception(err, "failed to read")
		}
		return err
	} else if n != len(b) {
		if l != nil {
			l.Error("failed to read a whole buffer, buf len: %d, read: %d", len(b), n)
		}
		return NewErr("failed to read a whole buffer, buf len: %d, read: %d", len(b), n)
	}

	return nil
}
