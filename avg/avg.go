/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 22:53
 *Task: Chapter 3 Lecture 3 HW 3.1 AVG
 */

package avg

//average func
func Average(arr [6]int) (avg float64) {
	for i := 0; i < len(arr); i++ {
		avg += float64(arr[i])
	}
	return avg / float64(len(arr))
}
