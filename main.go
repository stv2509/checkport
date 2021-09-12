package main

import (
	p "checkport/parsehcl"
	rc "checkport/rawconnect"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for {
		client := &p.Host{}
		hosts, conntimeout, reptimeout := p.ShowConfig()
		for _, port := range hosts {
			client = port
//			fmt.Println(client.Name, conntimeout)
			rc.RawConnect(&wg, client.Name, client.Proto, client.Ports, conntimeout)
		}
		time.Sleep(time.Duration(reptimeout) * time.Second)
	}

}
