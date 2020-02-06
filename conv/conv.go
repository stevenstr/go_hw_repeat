/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 17:35
 *Task: HW 5.2 Testing (unit testing)
 */

package conv

import (
	"strconv"
)

//MyStrToInt func
func MyStrToIntAtoi(s string) (int, error) {
	return strconv.Atoi(s)
}