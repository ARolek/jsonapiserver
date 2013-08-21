### Concurrent JSON API Server base
A simple concurrent JSON API server base in GO. This is not a webserver.

### Summary 
Concurrently handle JSON requests which can then be routed to application logic.

### Main features 

- Each client connection gets its own go routine
- Each client request gets its own go routine
- JSON validation incoming and outgoing
- Each request requires a "reqId" which will be passed in the response. This is how requests and responses are paired up on the client.

### Demo
By default, the server echos all incoming valid JSON requests. To start the server, open up terminal and issue the command:

```
go run main.go
```

Use telnet to connect to the server on port 8888 (or the port set using the --port flag). Send JSON data to the server and it will echo the request. You must provied a "reqId" in the request. For example: 

```
{"reqId": 12, "foo": "bar"}
```

### Command Line Parameters

--port : local port to bind to. defaults to 8888

### Why
I was working with reading net.Conn incoming data into buffers, but I did't like the fixed buffer size I was finding in most examples. To accommodate various sized requests, the buffer had to be rather large and required trimming before being passed to json.UnMarshal().

In order to know that a request is complete we need a [frame](http://en.wikipedia.org/wiki/Frame_(networking\)). JSON provides framing. Initially I was passing net.Conn into json.Decoder, but if the JSON is malformed there was a problem with the data sticking in the buffer. Now I use a bufio.Scan() to read the incoming byte data and use json.Unmarshal to validate the JSON data. 

Leveraging go routines, we get a clean and concurrent JSON API server base to work from.

### Questions / Feedback
[@alexrolek](https://twitter.com/alexrolek)
