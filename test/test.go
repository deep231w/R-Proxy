package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // 1. Dial the TCP connection
    conn, err := net.Dial("tcp", "www.google.com:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // 2. Format and send the raw HTTP GET request
    // Note: Each header line must end with \r\n, and the request ends with a blank line.
    fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: www.google.com\r\nConnection: close\r\n\r\n")

    // 3. Read the response
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
