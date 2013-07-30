package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
)

var port = flag.String("port", "8888", "port to connect to")

type Client struct {
	Conn net.Conn
	Req  map[string]interface{}
	Res  map[string]interface{}
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Listening on localhost:" + *port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			break
		}

		//	create a new go routine for each connection
		go handleConnection(conn)
	}
}

//	we have a connection, now listen for incoming JSON data
func handleConnection(conn net.Conn) {
	dec := json.NewDecoder(conn)
	for {
		client := Client{
			Conn: conn,
		}
		err := dec.Decode(&client.Req)
		if err != nil {
			client.Res = map[string]interface{}{
				"error": err.Error(),
			}
			go handleResponse(client)
			break
		}

		//	use reqId
		if _, ok := client.Req["reqId"]; !ok {

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
