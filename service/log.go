package service

import "log"

func Log(err error) {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	log.Println(err)
}
