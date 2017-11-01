/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 15:56:36
 */

package goutils

import (
	"io"
)

func Write(l Logger, w io.Writer, b []byte) error {
	if n, err := w.Write(b); err != nil {
		l.Exception(err, "failed to write")
		return err
	} else if n != len(b) {
		l.Error("failed to write a whole buffer, buf: %d, write: %d", len(b), n)
		return NewErr("failed to write a whole buffer, buf: %d, write: %d", len(b), n)
	}

	return nil
}
