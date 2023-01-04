package model

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"net"
)

var DB *gorm.DB

type ViaSSHDialer struct {
	client *ssh.Client
}

func (s *ViaSSHDialer) Dial(context context.Context, addr string) (net.Conn, error) {
	return s.client.Dial("tcp", addr)
}
func InitDB() {
	var u, p string
	u = viper.GetString("db.username")
	p = viper.GetString("db.password")
	config := &ssh.ClientConfig{
		User: u,
		Auth: []ssh.AuthMethod{
			ssh.Password(p),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, _ := ssh.Dial("tcp", "43.138.61.49:22", config)
	mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client}).Dial)
	dsn := fmt.Sprintf("%v:%v@mysql+tcp(127.0.0.1:3306)/csc?charset=utf8&parseTime=True&loc=Local", u, p)
	DB, _ = gorm.Open("mysql", dsn)
}
