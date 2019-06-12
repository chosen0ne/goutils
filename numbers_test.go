/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-12-13 17:15:53
 */

package goutils

import (
	"testing"
)

func TestDump_ok(t *testing.T) {
	inputs := []int64{10000, 0, -1, -1002, 999999999, -9999999999, 123455666}
	for _, n := range inputs {
		loadVal, err := LoadInt64(DumpInt64(n))
		if err != nil {
			t.Errorf("failed to load num, err: %v", err)
		}
		if n != loadVal {
			t.Errorf("%d failed, load val: %d\n", n, loadVal)
		}
	}
}
