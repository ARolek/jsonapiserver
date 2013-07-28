### JSON API Server
A simple JSON server in GO. This is not a webserver.

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

Use telnet to connect to the server on port 8888. Send JSON data to the server and it will echo the request.

### Config
Edit main.go

- PORT : local port to bind to

### Why
I was working with writting net.Conn.Read() into buffers, but I did not like the fixed size buffers that were required. In order to accommodate varios sized requests, I had to set the buffer to be rather large, and then trim it before it was passed to json.UnMarshal().

In order to know that a request is complete we need a [frame](http://en.wikipedia.org/wiki/Frame_(networking)). JSON provides framing, thus we can pass net.Conn into a json.Decoder and simply call Decode(). Leveraging go routines, we get a clean and concurrent JSON API server base to work from.

### Questions / Feedback
[@alexrolek](https://twitter.com/alexrolek)