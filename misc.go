/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-11-15 16:06:39
 */

package goutils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	_UNKNOWN_CALLER = "unknown caller information"
)

// CallerInfo returns the detail information in "file:lineno:funcname" format of the caller.
func CallerInfo(skip int) string {
	buf := &bytes.Buffer{}

	ptr, fname, lno, ok := runtime.Caller(skip)
	if !ok {
		return _UNKNOWN_CALLER
	}

	fn := runtime.FuncForPC(ptr)
	if fn == nil {
		return _UNKNOWN_CALLER
	}

	if _, err := fmt.Fprintf(buf, "%s:%d:%s", filepath.Base(fname), lno, fn.Name()); err != nil {
		return _UNKNOWN_CALLER
	} else {
		return string(buf.Bytes())
	}
}
