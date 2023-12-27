// each instance of this application will have its IP encrypted
// in a file in the assets dir, when sent a handshake takes place and you
// can dial and comm, idealy.

package pkg

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func encryptIP() string {

	plainText, err := getPublicIP()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("KEY")
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal("Error creating a cypher")
	}

	alloc := make([]byte, 169)
	c.Encrypt(alloc, []byte(plainText))

	return hex.EncodeToString(alloc)
}

func decrypt() string {
	file, err := os.OpenFile(".assets/access.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 169)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	encIp := data[:count]

	key := os.Getenv("KEY")
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	alloc := make([]byte, 169)
	c.Decrypt(alloc, encIp)
	return string(alloc[:])
}

func WriteIP() error {
	encryptedIP := encryptIP()

	createAssetsDir()
	f, err := os.OpenFile(".assets/access.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if _, err := f.Write([]byte(encryptedIP)); err != nil {
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
