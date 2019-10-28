package entity

import (
	"encoding/json"
	// "fmt"
	"os"
	"log"
)

type UserData struct {
	Name string
	Password string
	Email string
	Telephone string
}

type Storage struct {
	users []UserData
	// meetings []MeetingData
}

type Response struct {
	IsSuccess bool
	Body string
}


var storage Storage
var data_path = "./"
var user_file = "users"
var meeting_file = "meetings"
var buffer_size = 65536

var cur_user_file = "cur_user"
var cur_user UserData

var empty_user = UserData {"", "", "", ""}

func init() {
	readCurUser()
	readStorage()
}

func myUnmarshal(input []byte, target interface{}) error {
	if len(input) == 0 {
		return nil
	}
	return json.Unmarshal(input, target)
}

func readFile(path string) (data []byte) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		// log.Fatal(err)
		return nil
	}
	data = make([]byte, buffer_size)
	n, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(data[:n]))
	return data[:n]
}

func writeFile(path string, data []byte) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func readCurUser() {
	user_data := readFile(data_path+cur_user_file)
	err := myUnmarshal(user_data, &cur_user)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(cur_user)
}

func writeCurUser() {
	user_data, err := json.Marshal(&cur_user)
	if err != nil {
		log.Fatal(err)
	}
	writeFile(data_path+cur_user_file, user_data)
	// fmt.Println(user_data)
}

func readStorage() {
	user_data := readFile(data_path+user_file)
	err := myUnmarshal(user_data, &(storage.users))
	if err != nil {
		log.Fatal(err)
	}
	// meeting_data := readFile(data_path+meeting_file)
	// err = json.Unmarshal(meeting_data, &storage.meetings)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(storage)
}

func writeStorage() {
	user_data, err := json.Marshal(&(storage.users))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(storage.users)
	writeFile(data_path+user_file, user_data)
	// meeting_data, err := json.Marshal(storage.meetings)
	// err := writeFile(data_path+meeting_file, meeting_data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(meeting_data)
}

func Register(username, password, email, telephone string) Response {
	if cur_user.Name != "" {
		return Response{false, "you have been login, please logout first."}
	}
	isValid := true
	for _, u := range storage.users {
		// fmt.Println(u)
		if u.Name == username {
			isValid = false
			break
		}
	}
	if !isValid {
		return Response{false, "username has already been register."}
	}
	storage.users = append(storage.users, UserData{username, password, email, telephone})
	writeStorage();
	return Response{true, "register success!"}
}

func Login(username, password string) Response {
	if cur_user.Name != "" {
		return Response{false, "you have been login as " + cur_user.Name}
	}
	isValid := false
	for _, u := range storage.users {
		if u.Name == username {
			if (password == u.Password) {
				isValid = true
				cur_user = u
			}
			break
		}
	}
	if !isValid {
		return Response{false, "username or password incorrect."}
	}

	writeCurUser();
	return Response{true, "login success!"}
}

func Logout() Response {
	cur_user = empty_user
	writeCurUser()
	return Response{true, "logout success!"}
}

func QueryAllUsers() Response {
	if cur_user.Name == "" {
		return Response{false, "forbidden, please login first."}
	}
	res := "username\temail\ttelephone\n"
	for _, u := range storage.users {
		res += u.Name + "\t" + u.Email + "\t" + u.Telephone + "\n"
	}
	return Response{true, res}
}