### Concurrent JSON API Server
A simple concurrent JSON API server base in GO. This is not a webserver.

### Summary 
Concurrently handle JSON requests which can then be routed to application logic.

### Main features 

- Each client connection gets its own go routine
- Each client request gets its own go routine
- JSON validation incoming and outgoing

### Demo
By default, the server echos all incoming valid JSON requests. To start the server, open up terminal and issue the command:

```
go run main.go
```

Use telnet to connect to the server on port 8888 (or the port set using the --port flag). Send JSON data to the server and it will echo the request.

### Command Line Parameters

--port : local port to bind to

### Why
I was working with reading net.Conn incoming data into buffers, but I did't like the fixed buffer size requirement. In order to accommodate varios sized requests, the buffer had to be large and required trimming before being passed to json.UnMarshal().

In order to know that a network request is complete we need to [frame](http://en.wikipedia.org/wiki/Frame_(networking) it. JSON provides framing, allowing we can pass net.Conn into a json.Decoder and simply call Decode(). Leveraging go routines, we get a clean and concurrent JSON API server base to work from.

### Questions / Feedback
[@alexrolek](https://twitter.com/alexrolek)
