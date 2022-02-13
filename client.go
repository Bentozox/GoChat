package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	con      net.Conn
	username string
	state    ClientRestriction
}

func (client *Client) run(cManager *ClientManager) {
	client.state = NONE
	con := client.con

	// Send message to client
	con.Write([]byte("Bienvenue !\n"))

	client.login()

	line, err2 := bufio.NewReader(con).ReadString('\n')

	for line != "" {
		if err2 != nil {
			fmt.Printf("Error while reading message : %s\n", err2)
			return
		}
		line = strings.ReplaceAll(line, "\n", "")
		message := fmt.Sprintf("[CLIENT : %s] %s", client.username, line)

		fmt.Print(message)
		cManager.sendExcept([]byte(message), Logged, client)

		line, err2 = bufio.NewReader(con).ReadString('\n')
	}
}

func (client *Client) login() {
	con := client.con

	client.writeStr("Entrez votre nom d'utilisateur : ")
	line, _, err2 := bufio.NewReader(con).ReadLine()

	for {
		if err2 != nil {
			fmt.Printf("[ERROR] Can't read message")
			return
		}

		if len(line) > 0 {
			client.username = string(line)
			break
		}

		client.writeStr("Entrez votre nom d'utilisateur : ")
		line, _, err2 = bufio.NewReader(con).ReadLine()
	}

	client.writeStr("Bravo ! tu peux désormais envoyer et recevoir des messages.\n")
	client.state = Logged
}

func (client *Client) write(data []byte) {
	client.con.Write(data)
}

func (client *Client) writeStr(str string) {
	client.con.Write([]byte(str))
}
