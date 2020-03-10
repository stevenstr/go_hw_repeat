/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 23:59
 *Task: Chapter 3 Lecture 3 HW 3.2 Max
 */

package maxp

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test_MAX func
func Test_MAX(t *testing.T) {
	sl := []string{"a", "b", "cccp", "rus"}
	res := Max(sl)

	assert.Equal(t, "cccp", res)

	sl1 := []string{"ananas", "b", "cccp", "russians"}
	res1 := Max(sl1)

	assert.Equal(t, "russians", res1)

	sl2 := []string{"Lesha", "Kolya", "Ivan", "Kukumberbetch"}
	res2 := Max(sl2)

	assert.Equal(t, "Kukumberbetch", res2)
}

func Benchmark_MAX(b *testing.B) {
	sl := []string{"a"}
	for i := 0; i < b.N; i++ {
		st := strconv.Itoa(i)
		sl = append(sl, st)
		_ = Max(sl)
	}
}
