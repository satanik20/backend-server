package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Launching server...")

	// Resolve TCP Address
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Printf("Unable to resolve IP")
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("Unable to start listener - %s", err)
	}

	for {
		// Listen for an incoming connection.
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Printf("Listener returned - %s", err)
			break
		}
		// Enable Keepalives
		err = conn.SetKeepAlive(true)
		if err != nil {
			fmt.Printf("Unable to set keepalive - %s", err)
		}
		go func(conn net.Conn) {
			i := 0
			for {
				b := make([]byte, 1024)
				_, err := conn.Read(b)
				if err != nil {
					if err != io.EOF {
						i++
						if i > 3 {
							log.Println("server unhealthy error:", err)
						}
					}
					break
				}
				fmt.Println(string(b))
			}
			fmt.Println("Stopping handle connection")
		}(conn)

		/* CPU Utilization
		mem := &runtime.MemStats{}
		for {
			cpu := runtime.NumCPU()
			log.Println("Server CPU Utilization:", cpu)

			// Byte
			runtime.ReadMemStats(mem)
			log.Println("Server Memory Utilization:", mem.Alloc)

			time.Sleep(20 * time.Second)
			log.Println("----------")
		}*/
	}
}
