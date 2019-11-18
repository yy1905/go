package publicutils

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"net"
)

/*

example:
	client, session, e := SSHConnect("root", "111111", "192.168.1.11", 22)
	bytes, _ := session.Output("df -lh")
or
	bytes, _ := session.CombinedOutput("df -lh")

fmt.Println(string(bytes))

*/

func SSHConnect(user, password, host string, port int) (*ssh.Client, *ssh.Session, error) {

	var (
		authMethods  []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// 创建密码校验方法
	authMethods = make([]ssh.AuthMethod, 0)
	authMethods = append(authMethods, ssh.Password(password))

	// 创建一个格式合法的回调函数
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	/*创建SSH客户端配置*/
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: authMethods,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// 连接地址
	addr = fmt.Sprintf("%s:%d", host, port)

	// 拨号并获取SSH客户端
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, err
	}

	// 创建新的会话
	if session, err = client.NewSession(); err != nil {
		return nil, nil, err
	}

	return client, session, nil
}



/*
func main() {

	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	ip := flag.String("i", "", "IP")
	cmd := flag.String("e", "", "command line")
	flag.Parse()

	session, err := connect(*username, *password, *ip, 22)

	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run(*cmd)
}

*/