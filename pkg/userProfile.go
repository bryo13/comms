package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

// check if user has a username
// read user name
// transcate username to message
func createUser() string {
	var username string
	fmt.Println("Enter your username")
	fmt.Fscanf(os.Stdin, "%s: \n", &username)
	return username
}

// returns the path our folder lives
func MakeDir() string {
	home, _ := homedir.Dir()
	path := fmt.Sprintf("%s/%s", home, ".terminal_text")
	err := os.Mkdir(path, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	return path
}

func MakeUserFIle(username string) {
	dirPath := MakeDir()

	filePath := fmt.Sprintf("%s/%s", dirPath, "t_t.txt")
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// format username = username + newline
	formated_user := fmt.Sprintf("%s\n", username)
	if _, err := f.Write([]byte(formated_user)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func checklIfUserExists() bool {
	dirpath := MakeDir()
	if err := os.Mkdir(dirpath, 0750); err != nil {
		return true
	}
	return false
}

func getUserName(filepath string) string {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scanner: %v\n", err)
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastLine string
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scanner: %v\n", err)
	}

	return lastLine
}

func InitUser() {
	dirPath := MakeDir()
	filePath := fmt.Sprintf("%s/%s", dirPath, "t_t.txt")
	userAvailable := checklIfUserExists()
	if userAvailable {
		user := getUserName(filePath)
		fmt.Println(user, "welcome back")
		return
	}

	username := createUser()
	MakeUserFIle(username)
}
