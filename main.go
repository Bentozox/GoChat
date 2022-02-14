package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/**
Try him by using (linux !, windows ?) : telnet localhost 58000
*/
func main() {
	cManager := ClientManager{}
	listen, err := net.Listen("tcp", "0.0.0.0:58000")

	if err != nil {
		fmt.Printf("Can't start server : %s\n", err)
		return
	}

	go serverConsole(&cManager)

	for {
		con, err := listen.Accept()

		if err != nil {
			fmt.Printf("Error while accepting client : %s\n", err)
			return
		} else {
			client := Client{con: con}
			cManager.addClient(&client)
			go client.run(&cManager) // Async
		}

	}
}

func serverConsole(cManager *ClientManager) {
	message, _, err := bufio.NewReader(os.Stdin).ReadLine()

	for {

		if err != nil {
			fmt.Printf("[SERVER] Message not submited. Error : %s\n", err)
			return
		}

		if len(message) > 0 {
			cManager.send(append([]byte("[SERVER] "), message...), Logged)
		}

		println("[SERVER] Message submited")
		message, _, err = bufio.NewReader(os.Stdin).ReadLine()
	}

}
