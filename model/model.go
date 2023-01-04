package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//type ViaSSHDialer struct {
//	client *ssh.Client
//}

//func (s *ViaSSHDialer) Dial(context context.Context, addr string) (net.Conn, error) {
//	return s.client.Dial("tcp", addr)
//}

func InitDB() {
	var u, p string
	u = viper.GetString("db.username")
	p = viper.GetString("db.password")
	//config := &ssh.ClientConfig{
	//	User: u,
	//	Auth: []ssh.AuthMethod{
	//		ssh.Password(p),
	//	},
	//	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	//}
	//client, _ := ssh.Dial("tcp", "43.138.61.49:22", config)
	//mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client}).Dial)
	//"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s"
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/csc?charset=utf8&parseTime=True&loc=Local", u, p)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
