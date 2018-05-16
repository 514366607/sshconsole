package sshconsole

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHPW 账号密码连接
func SSHPW(user, password, ip string, port int) *ssh.Client {
	config := formatConfig(user, []ssh.AuthMethod{ssh.Password(password)})
	Client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), &config)
	if err != nil {
		panic(err)
	}
	return Client
}

// SSHRsaFile rsa文件连接
func SSHRsaFile(ip string, port int, user, keyPath string) *ssh.Client {
	config := formatConfig(user, ReadKey([]string{keyPath}))
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), &config)
	if err != nil {
		panic(err)
	}
	return conn
}

// SSHRsaKey rsa密钥连接
func SSHRsaKey(ip string, port int, user string, key ssh.AuthMethod) *ssh.Client {
	config := formatConfig(user, []ssh.AuthMethod{key})
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), &config)
	if err != nil {
		panic(err)
	}
	return conn
}

// Command 命令行运行
func Command(Client *ssh.Client) {
	defer Client.Close()
	a := bufio.NewReader(os.Stdin)
	for {
		b, _, z := a.ReadLine()
		if z != nil {
			return
		}
		command := string(b)
		if session, err := Client.NewSession(); err == nil {
			defer session.Close()
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			session.Run(command)
		}
	}
}
