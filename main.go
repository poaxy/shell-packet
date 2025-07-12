package main

import (
	"fmt"
	"github.com/songgao/water"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func main() {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatalf("Failed to create TUN device: %s", err)
	}

	log.Printf("TUN device created: %s\n", iface.Name())

	if err := configureDevice(iface.Name()); err != nil {
		log.Fatalf("Failed to configure device: %s", err)
	}

	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <user> <host:port> <password>\n", os.Args[0])
		os.Exit(1)
	}

	user := os.Args[1]
	addr := os.Args[2]
	password := os.Args[3]

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Config: ssh.Config{
			Ciphers: []string{"chacha20-poly1305@openssh.com", "aes128-ctr", "aes192-ctr", "aes256-ctr"},
		},
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to get stdin pipe: %s", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout pipe: %s", err)
	}

	go func() {
		packet := make([]byte, 1500)
		for {
			n, err := iface.Read(packet)
			if err != nil {
				log.Printf("Error reading from TUN device: %s", err)
				continue
			}
			// This is where we would forward the packet over the SSH connection.
			_, err = stdin.Write(packet[:n])
			if err != nil {
				log.Printf("Error writing to SSH session: %s", err)
			}
		}
	}()

	go func() {
		packet := make([]byte, 1500)
		for {
			n, err := stdout.Read(packet)
			if err != nil {
				log.Printf("Error reading from SSH session: %s", err)
				continue
			}
			// This is where we would write the packet to the TUN device.
			_, err = iface.Write(packet[:n])
			if err != nil {
				log.Printf("Error writing to TUN device: %s", err)
			}
		}
	}()

	// Enable compression
	if err := session.RequestSubsystem("zlib@openssh.com"); err != nil {
		log.Printf("Failed to enable compression: %s", err)
	}

	fmt.Println("Successfully connected to SSH server!")

	// Keep the program running
	select {}
}
