package main


import (
  "fmt"
  "net"
  "strings"
)

func main() {
    // Create a listener on port 8080.
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Accept connections on the listener.
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }

        // Read the connection.
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Check for known intrusion signatures in the data.
        for _, signature := range []string{"root", "exploit", "password"} {
            if strings.Contains(string(buf[:n]), signature) {
                fmt.Println("Intrusion detected!")
                return
            }
        }
    }
}
