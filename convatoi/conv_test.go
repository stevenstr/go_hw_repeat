/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 17:35
 *Task: 5.2 Testing
 */

package convatoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestDone func
func TestDone(t *testing.T) {
	k, err := MyStrToIntAtoi("3737")
	assert.Equal(t, k, 3737, "equal")
	assert.Nil(t, err, nil)
}

//TestNull func
func TestNull(t *testing.T) {
	k, err := MyStrToIntAtoi("")
	assert.Equal(t, k, 0, "empty")
	assert.Error(t, err, "error empty")
}

//TestMix func
func TestMix(t *testing.T) {
	k, err := MyStrToIntAtoi(" m2k4h3")
	assert.Equal(t, k, 0, "mixed")
	assert.Error(t, err, "error mixed")
}

//TestOver func
func TestOver(t *testing.T) {
	k, err := MyStrToIntAtoi("929292992292992929292992929292992929292292929")
	assert.Equal(t, k, 9223372036854775807, "overbuffered")
	assert.Error(t, err, "overbuffered")
}
