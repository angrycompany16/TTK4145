package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

var (
	backupFlag = "--backup"
)

func main() {
	args := os.Args

	fmt.Println()
	fmt.Println()
	fmt.Println()

	if len(args) <= 1 {
		fmt.Println("Make primary")
		makeBackup()
		runAsPrimary(0)
	} else if args[1] == backupFlag {
		fmt.Println("Make backup")
		runAsBackup()
	}

	fmt.Println("Heisann")
}

func runAsPrimary(start uint32) {
	var i uint32 = start

	address := net.UDPAddr{
		IP:   nil,
		Port: 45535,
	}

	conn, err := net.DialUDP("udp", nil, &address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		i++
		fmt.Println("The magic number is: ", i)

		message := make([]byte, 4)
		binary.LittleEndian.PutUint32(message, i)
		conn.Write(message)

		time.Sleep(time.Second)
	}

}

func makeBackup() {
	exec.Command("gnome-terminal", "--", "go", "run", "primary.go", backupFlag).Run()
}

func runAsBackup() {
	address := net.UDPAddr{
		IP:   nil,
		Port: 45535,
	}

	var conn *net.UDPConn
	var err error

	for {
		conn, err = net.ListenUDP("udp", &address)
		if err == nil {
			break
		}
		fmt.Println("Error encountered when connecting to primary via UDP, retrying...")
		time.Sleep(time.Second * 1)
	}

	var memory uint32
	buffer := make([]byte, 4)

	for {
		// Read + store last view of number
		// If message times out, switch to primary
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Print(err)
			continue
		}

		n, _, err := conn.ReadFromUDP(buffer) // n is length of data
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				log.Print(err)
				continue
			}
			fmt.Println("Primary timed out, switching to backup...")
			break
		}

		memory = binary.LittleEndian.Uint32(buffer[:n])
		_, err = fmt.Println("Remembering value: ", memory)
		if err != nil {
			log.Print(err)
			continue
		}
	}

	conn.Close()
	makeBackup()
	runAsPrimary(memory)
}
