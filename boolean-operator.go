package main

import "fmt"

func main() {
	var examScore = 80
	var attendanceScore = 80

	// var successExam = examScore >= 80
	// var successAttendance = attendanceScore >= 80
	// var isSuccess = successExam && successAttendance
	fmt.Println(examScore >= 80 && attendanceScore >= 80)
}
