package pkg

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

func CheckInternet() bool {
	timeout := 2 * time.Second
	_, err := net.DialTimeout("tcp", "www.google.com:80", timeout)
	return err == nil
}

// to read IP from our assets dir
func HandleConn() {
	listener, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return
	}

	defer listener.Close()

	var wg sync.WaitGroup

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Printf("Error accepting connection: %v\n", err)
				break
			}

			wg.Add(1)
			go listenAndWrite(conn, &wg)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Close the listener and wait for all goroutines to finish
	listener.Close()
	wg.Wait()
}

func listenAndWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}

		data := buf[:n]
		fmt.Printf("Received data: %s", data)

		// Echo the data back to the client
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("Error writing to connection: %v\n", err)
			return
		}
	}
}

// Write or ammend friendlist
// get IP
// listen on the connection
// write on the connection
