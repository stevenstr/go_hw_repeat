/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 23:59
 *Task: Chapter 3 Lecture 3 HW 3.2 Max
 */

package maxp

import "unicode/utf8"

//Max func
func Max(sl []string) (res string) {
	ln := 0
	for _, v := range sl {
		if ln < utf8.RuneCountInString(v) {
			ln = utf8.RuneCountInString(v)
			res = v
		}
	}
	return res
}
