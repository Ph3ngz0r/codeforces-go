// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, countAlternatingSubarrays, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-391/problems/count-alternating-subarrays/
// https://leetcode.cn/problems/count-alternating-subarrays/