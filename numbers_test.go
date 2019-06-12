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
		buf, dumpErr := DumpInt64(n)
		if dumpErr != nil {
			t.Errorf("failed to dump num, num: %d, err: %v", n, dumpErr)
		}
		loadVal, loadErr := LoadInt64(buf)
		if loadErr != nil {
			t.Errorf("failed to load num, err: %v", loadErr)
		}
		if n != loadVal {
			t.Errorf("%d failed, load val: %d\n", n, loadVal)
		}
	}
}

func TestWriteFixedNum_ok(t *testing.T) {
	inputs64 := []int64{0, 100, 0x01020304050607, -100000, -0x0102030405}
	inputs32 := []int32{-1, 0, 1000, 0x010205, -0xff91ea}
	inputs16 := []int16{-1, 0, 1000, 0x0102, -0x4fff}

	for _, n := range inputs64 {
		buf, dumpErr := DumpFixedInt64(n)
		if dumpErr != nil {
			t.Errorf("failed to dump int64, int64: %d, err: %v", n, dumpErr)
		}
		loadVal, loadErr := LoadFixedInt64(buf)
		if loadErr != nil {
			t.Errorf("failed to load num, err: %v", loadErr)
		}
		if n != loadVal {
			t.Errorf("%d failed, load val: %d\n", n, loadVal)
		}
	}

	for _, n := range inputs32 {
		buf, dumpErr := DumpFixedInt32(n)
		if dumpErr != nil {
			t.Errorf("failed to dump int32, int32: %d, err: %v", n, dumpErr)
		}
		loadVal, loadErr := LoadFixedInt32(buf)
		if loadErr != nil {
			t.Errorf("failed to load num, err: %v", loadErr)
		}
		if n != loadVal {
			t.Errorf("%d failed, load val: %d\n", n, loadVal)
		}
	}

	for _, n := range inputs16 {
		buf, dumpErr := DumpFixedInt16(n)
		if dumpErr != nil {
			t.Errorf("failed to dump int16, int16: %d, err: %v", n, dumpErr)
		}
		loadVal, loadErr := LoadFixedInt16(buf)
		if loadErr != nil {
			t.Errorf("failed to load num, err: %v", loadErr)
		}
		if n != loadVal {
			t.Errorf("%d failed, load val: %d\n", n, loadVal)
		}
	}
}
