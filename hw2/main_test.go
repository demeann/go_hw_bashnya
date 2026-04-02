package main

import (
	"testing"
)

func TestOper(t *testing.T) {
	testTable := []struct {
		num    int
		expect int
	}{
		{num: 0, expect: 0},
		{num: 10, expect: 117},
		{num: 12306, expect: 479934},
		{num: -12307, expect: 12308},
		{num: 14, expect: 14859},
		{num: 18, expect: 6435},
		{num: 117, expect: 13726},
	}

	for _, testCase := range testTable {
		result := Oper(testCase.num)

		t.Logf("num(%d) result: %d", testCase.num, result)

		if result != testCase.expect {
			t.Errorf("Oper(%d): expect %d, but got %d", testCase.num, testCase.expect, result)
		}
	}
}
