package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func client() {

	var t = 1

	conn, errDial := net.Dial("tcp", "127.0.0.1:8080")

	if errDial != nil {
		log.Fatal("ERROR: ", errDial)
		return
	}

	for {

		if t > 1 {
			conn.Close()
			return
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")

		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("ERROR: ", err)
			return
		}

		fmt.Fprintf(conn, text+"\n")

		message, errMes := bufio.NewReader(conn).ReadString('\n')

		if errMes != nil {
			log.Fatal("ERROR: ", errMes)
			return
		}

		fmt.Print("Message from server: " + "Hello, " + message)
		t += 1
	}
}

func main() {

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("ERROR: ", err)
		return
	}

	fmt.Println("Server is starting...")

	go client()

	conn, err := ln.Accept()

	if err != nil {
		log.Fatal("ERROR: ", err)
		return
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal("ERROR: ", err)
		}

		conn.Write([]byte(message + "\n"))
	}

}
