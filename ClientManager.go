package main

type ClientManager struct {
	clients []*Client
}

func (cm *ClientManager) addClient(client *Client) {
	cm.clients = append(cm.clients, client)
}

func (cm *ClientManager) removeClient() {
	// TODO
}

func (cm *ClientManager) send(data []byte, restrict ClientRestriction) {
	data = append(data, '\n')

	for _, client := range cm.clients {
		if restrict == NONE || client.state == restrict {
			client.write(data)
		}
	}
}

func (cm *ClientManager) sendExcept(data []byte, restrict ClientRestriction, exceptClient *Client) {
	data = append(data, '\n')

	for _, client := range cm.clients {
		if restrict == NONE || client.state == restrict && exceptClient != client {
			client.write(data)
		}
	}
}
