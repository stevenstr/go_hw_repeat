/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 22:35
 *Task: Class Work L5 Testing (unit testing, )
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Unit test fr dummy, Tests passed
//Test_dummy func
func Test_dummy(t *testing.T) {
	//just check so that returns will be right
	k, _ := dummy(55)
	assert.Equal(t, 1, k)
	c, _ := dummy(5)
	assert.Equal(t, -1, c)
	v, _ := dummy(10)
	assert.Equal(t, 0, v)

	//just check so that err will be right
	_, err := dummy(55)
	assert.EqualError(t, err, "more")
	_, err = dummy(5)
	assert.EqualError(t, err, "less")
	_, err = dummy(10)
	assert.EqualError(t, err, "zero")
}

func Test_summ(t *testing.T) {
	//just check so that returns will be right
	x := []int{1, 2, 3}
	k := summ(x)
	assert.Equal(t, 6, k)
	y := []int{2, 2, 2, 2}
	c := summ(y)
	assert.Equal(t, 8, c)
	z := []int{5, 6}
	v := summ(z)
	assert.Equal(t, 11, v)
}
