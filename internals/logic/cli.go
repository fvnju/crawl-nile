package logic

import "fmt"

func RunCli(username, password string) {
	session, err := getSessionToken()
	if err != nil {
		fmt.Println("Error getting session token: ", err)
	}
	err = loginToNileSIS(username, password, session)
	if err != nil {
		fmt.Println("Error logging in: ", err)
	}
	courses, err := scrapper(username, session)
	if err != nil {
		fmt.Println("Error getting courses: ", err)
	}
	err = logOut(username, session)
	if err != nil {
		fmt.Println("Error logging out: ", err)
	}

	for i, course := range courses {
		fmt.Printf("S/N: %d\nCode: %s\nName: %s\nGrade: %s\nCredit: %d\n\n", i+1, course.Code, course.Name, course.Grade, course.Credit)
	}
}
