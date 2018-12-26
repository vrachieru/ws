package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/http"
    "os"
    "strings"

    "github.com/gorilla/websocket"
)

type flagArray []string

func (i *flagArray) String() string {
    return strings.Join(*i, ", ")
}

func (i *flagArray) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func makeHeader(flags flagArray) http.Header {
    headers := make(http.Header)
    for _, f := range flags {
        kv := strings.Split(f, ":")
        if len(kv) != 2 {
            continue
        }
        key, val := kv[0], kv[1]
        if len(key) == 0 || len(val) == 0 {
            continue
        }
        key, val = strings.TrimSpace(key), strings.TrimSpace(val)
        headers.Add(key, val)
    }
    return headers
}

func connect(url string, headers http.Header) *websocket.Conn {
    fmt.Printf("Trying %s...\n", url)
    ws, _, err := websocket.DefaultDialer.Dial(url, headers)
    if err != nil {
        fmt.Printf("Could not connect to %s: %s", url, err)
        os.Exit(1)
    }
    fmt.Printf("Connected to %s.\n", url)
    fmt.Printf("Exit with CTRL+C.\n\n")
    return ws
}

func send(ws *websocket.Conn) {
    fmt.Printf("> ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text := scanner.Text()
        ws.WriteMessage(websocket.TextMessage, []byte(text))
    }
}

func receive(ws *websocket.Conn) {
    for {
        _, message, err := ws.ReadMessage()
        if err != nil {
            fmt.Printf("Could not read message from socket: %s", err)
            os.Exit(1)
        }
        fmt.Printf("< %s\n> ", message)
    }
}

var (
    url string
    headers flagArray
)

func main() {
    flag.Usage = func() {
        fmt.Printf("Usage: %s [options] url\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
    }
    flag.Var(&headers, "H", "Extra header to use in the WS request.\n" +
                            "You can specify as many as needed by repeating the flag.\n" +
                            "Example: -H \"Foo-Header: foo\" -H \"Bar-Header: bar\"")
    flag.Parse()

    url = flag.Arg(0)
    if url == "" {
        flag.Usage()
    }

    ws := connect(url, makeHeader(headers))
    go send(ws)
    receive(ws)
}
