package sshconsole

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHPW 账号密码连接
func SSHPW(user, password, ip string, port int) *ssh.Client {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Conf.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	Client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), &Conf)
	if err != nil {
		panic(err)
	}
	return Client
}

// SSHRsaFile rsa文件连接
func SSHRsaFile(ip string, port int, user, keyPath string) *ssh.Client {

	// 读取文件
	privateKey := ReadKey([]string{keyPath})

	var auths []ssh.AuthMethod
	auths = append(auths, privateKey...)
	config := formatRsaConfig(user, auths)
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), &config)
	if err != nil {
		panic(err)
	}
	return conn
}

// SSHRsaKey rsa密钥连接
func SSHRsaKey(ip string, port int, user string, key ssh.AuthMethod) *ssh.Client {
	var auths = []ssh.AuthMethod{key}
	config := formatRsaConfig(user, auths)
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
