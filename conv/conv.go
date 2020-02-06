/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 17:30
 *Task: HW 5.1 Testing (unit testing)
 */

package conv

import (
	"strconv"
)

//MyStrToInt func
func MyStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
