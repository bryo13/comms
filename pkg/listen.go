package pkg

import (
	"fmt"
	"net"
	"os"
)

func ListenAndWrite() {
	reader, writer := net.Pipe()

	defer func() {
		reader.Close()
		writer.Close()
	}()

	dt := getData()

	go func() {
		fmt.Println("Writing to pipe")
		writer.Write([]byte(dt))
		writer.Close()

	}()

	// read from pipe, with blocking

	data := make([]byte, 32)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading from pipe:", err)
		return
	}

	fmt.Println("Read from pipe:", string(data[:n]))

}

func getData() string {
	var input string
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}

	return input
}
