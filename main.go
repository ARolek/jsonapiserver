package main

import (
	"encoding/json"
	"log"
	"net"
)

const (
	PORT = "8888"
)

type Client struct {
	Conn net.Conn
	Req  map[string]interface{}
	Res  map[string]interface{}
}

func main() {
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Listening on localhost:" + PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			break
		}
		client := Client{
			Conn: conn,
		}
		//	create a new go routine for each connection
		go handleConnection(client)
	}
}

//	we have a connection, now listen for incoming json data
func handleConnection(client Client) {
	dec := json.NewDecoder(client.Conn)
	for {
		err := dec.Decode(&client.Req)
		if err != nil {
			client.Res = map[string]interface{}{
				"error": err.Error(),
			}
			go handleResponse(client)
			break
		}

		//	create a new go routine for each request on this connection
		go handleRequest(client)
	}
}

func handleRequest(client Client) {
	//	application logic routing goes here.
	//	when the appliction logic is done, save the response to client.Res

	//	Example: copy req to res to demo an echo server
	client.Res = client.Req

	handleResponse(client)
}

func handleResponse(client Client) {
	//	encode our response to JSON
	json, err := json.Marshal(client.Res)
	if err != nil {
		log.Println("json marshal error. response data: ", client.Res)
	}

	//	send a response to the client
	_, err = client.Conn.Write(json)
}
