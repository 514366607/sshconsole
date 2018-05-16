package main

import (
	"sshConsole"
)

func main() {
	client := sshconsole.SSHPW("root", "", "192.168.0.10", 22)
	defer client.Close()
	sshconsole.Command(client)
}
