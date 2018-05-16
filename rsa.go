package sshconsole

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// ReadKey 传入文件地址取得key
func ReadKey(keypath []string) []ssh.AuthMethod {
	var privateKey []ssh.AuthMethod
	for _, v := range keypath {
		buf, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Printf("读取key文件%s失败:\n%s\n", v, err)
			os.Exit(1)
		}
		privateKey = append(privateKey, FormatPublicKey(buf))
	}
	return privateKey
}

// FormatPublicKey 格式化key
func FormatPublicKey(buf []byte) ssh.AuthMethod {
	signer, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		fmt.Printf("解析key%s失败:\n%s\n", string(buf), err)
		os.Exit(1)
	}
	return ssh.PublicKeys(signer)
}

// formatConfig
// 格式化配置
func formatConfig(user string, auths []ssh.AuthMethod) ssh.ClientConfig {
	config := ssh.ClientConfig{
		User: user,
		Auth: auths,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return config
}
