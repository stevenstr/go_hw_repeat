/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 17:35
 *Task: HW 5.2 Testing (unit testing)
 */

package convstr

import "fmt"

//MyStrToIntSscanf func
func MyStrToIntSscanf(s string) (int, error) {
	var a int
	_, err := fmt.Sscanf(s, "%d", &a)
	return a, err
}
