package main

import "fmt"

type UserWarnet struct {
	FirstName string
	LastName  string
	Email     string
	isActive  bool
}

type GroupUser struct {
	GroupName   string
	Admin       UserWarnet
	Users       []UserWarnet
	isAvailable bool
}

type GroupUserMap struct {
	GroupName   string
	Admin       UserWarnet
	Users       map[string]UserWarnet
	isAvailable bool
}

func main() {
	user := UserWarnet{
		FirstName: "Tenri",
		LastName:  "Tend",
		Email:     "talkanda@gmail.com",
		isActive:  true,
	}

	user2 := UserWarnet{
		FirstName: "Angga",
		LastName:  "Budi",
		Email:     "awawawaw@gmail.com",
		isActive:  true,
	}

	display1 := user.displayGroup()
	display2 := user2.displayGroup()

	fmt.Println(display1)
	fmt.Println(display2)

	//group user
	group := GroupUser{
		GroupName:   "Warnet Jakarta",
		Admin:       user,
		Users:       []UserWarnet{user, user2},
		isAvailable: false,
	}

	fmt.Println(group)
	userMap := make(map[string]UserWarnet)
	userMap[user.FirstName] = user
	userMap[user2.FirstName] = user2

	groupMap := GroupUserMap{
		GroupName:   "Warnet Bandung",
		Admin:       user,
		Users:       userMap,
		isAvailable: false,
	}

	fmt.Println(groupMap)
}

func (userData UserWarnet) displayGroup() string {
	return fmt.Sprintf("Hello %s %s, please check your inbox, we have sent confirmation link to your email at %s", userData.FirstName, userData.LastName, userData.Email)
}
