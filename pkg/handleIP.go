// each instance of this application will have its IP encrypted
// in a file in the assets dir, when sent a handshake takes place and you
// can dial and comm, idealy.

// aes block counter

package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type enc struct {
	nonce []byte
}

var non []byte

func init() {
	e := new(enc)
	non = e.generateNonce()
}

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
	checkErr("reading IP", err)

	err = godotenv.Load()
	checkErr("error opening .env", err)

	key := os.Getenv("KEY")
	c, err := aes.NewCipher([]byte(key))
	checkErr("error creating a cypher", err)

	aesgcm, err := cipher.NewGCM(c)
	checkErr("error creating gcm", err)

	cipherText := aesgcm.Seal(nil, non, []byte(plainText), nil)
	return hex.EncodeToString(cipherText)
}

func Decrypt() string {
	encData := encryptIP()

	err := godotenv.Load()
	checkErr("error opening .env", err)

	key := os.Getenv("KEY")
	c, err := aes.NewCipher([]byte(key))
	checkErr("error creating a cypher", err)

	aesgcm, err := cipher.NewGCM(c)
	checkErr("error creating gcm", err)

	hexEncData, _ := hex.DecodeString(encData)
	decData, err := aesgcm.Open(nil, non, hexEncData, nil)
	checkErr("Error decrypting", err)
	return string(decData)
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

// error handling helper func
func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("Error in %s. The err is %v", msg, err.Error())
		os.Exit(2)
	}
}

// helper nonce gen
func (e *enc) generateNonce() []byte {
	e.nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, e.nonce); err != nil {
		panic(err.Error())
	}
	return e.nonce
}
