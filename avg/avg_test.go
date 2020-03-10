/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 22:53
 *Task: Chapter 3 Lecture 3 HW 3.1 AVG
 */

package avg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test_AVG func
func Test_AVG(t *testing.T) {
	arr := [6]int{0, 0, 0, 0, 0, 0}
	res := Average(arr)

	assert.Equal(t, 0.0, res)

	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	res1 := Average(arr1)

	assert.Equal(t, 3.5, res1)
}
