package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		client := Client{
			Conn: conn,
		}

		var reqData map[string]interface{}

		err := json.Unmarshal(scanner.Bytes(), &reqData)
		if err != nil {
			log.Println(err)
			client.Res = map[string]interface{}{
				"success": false,
				"error":   "JSON decode fail - " + err.Error(),
			}
			go handleResponse(client)
			continue
		}

		//	add our decoded JSON data to our client struct
		client.Req = reqData

		//	check for a reqId. this is how the client knows what response
		//	goes with what request
		if client.Req["reqId"] == "" || client.Req["reqId"] == nil {
			client.Res = map[string]interface{}{
				"success": false,
				"error":   "reqId not provided",
			}
			go handleResponse(client)
			continue
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
