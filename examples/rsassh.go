package main

import (
	"sshConsole"
)

func main() {
	client := sshconsole.SSHRsaFile("test.wondergm.com", 22, "root", "/Users/hible/.ssh/id_rsa")
	defer client.Close()
	sshconsole.Command(client)
}
