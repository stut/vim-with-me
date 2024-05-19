package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/theprimeagen/vim-with-me/pkg/v2/assert"
)

func main() {
    var addr string = ""
    flag.StringVar(&addr, "addr", "localhost:42069", "the address for the relay driver to send messages to")

    u := url.URL{Scheme: "ws", Host: addr, Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    assert.NoError(err, "unable to connect to the websocket server")

    for {
        mt, msg, err := c.ReadMessage()
        assert.NoError(err, "read message failed")
        fmt.Printf("mt=%d msg=%d\n", mt, len(msg))
    }
}

