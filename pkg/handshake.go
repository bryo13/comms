// each instance of this application will have its IP encrypted
// in a file in the assets dir, when sent a handshake takes place and you
// can dial and comm, idealy.

package pkg

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getPublicIP() (string, error) {
	checkInternt := CheckInternet()
	if checkInternt {
		resp, err := http.Get("https://api.ipify.org")
		if err != nil {
			log.Fatalln("couldnt find your public IP")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("couldnt read your public IP")
		}

		return string(body), nil
	}

	return "", errors.New("please check your internet")
}

func encryptIP() {}

func WriteIP() error {
	myIP, err := getPublicIP()
	if err != nil {
		return err
	}
	// format IP with a newline character
	formatedIP := fmt.Sprintf("%s\n", myIP)
	createAssetsDir()
	f, err := os.OpenFile(".assets/access.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if _, err := f.Write([]byte(formatedIP)); err != nil {
		f.Close()
		log.Fatal(err)
		return err
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// create the assest dir
func createAssetsDir() {
	err := os.Mkdir(".assets", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}
