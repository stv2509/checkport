package checkport

import (
	"fmt"
	"net"
	"sync"
	"time"
	"os"
)

func RawConnect(wg *sync.WaitGroup, host, proto string, ports []string, sec int) {
	wg.Add(len(ports))
	for _, port := range ports {
		go func(port string) {
			defer wg.Done()
			timeout := time.Duration(sec) * time.Second
			conn, err := net.DialTimeout(proto, net.JoinHostPort(host, port), timeout)
			if err != nil {
				node_name := os.Getenv("NODE_NAME")
				fmt.Println(time.Now().Format("02.01.2006 15:04:05"), node_name, "Connecting error:", net.JoinHostPort(host, port), err.Error())
			}
			if conn != nil {
				defer conn.Close()
//				fmt.Println("Opened", net.JoinHostPort(host, port))
			}
		}(port)
	}
	wg.Wait()
}


