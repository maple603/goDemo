package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func runClient(port int, file string) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Printf("无法建立连接", err)
	}
	defer conn.Close()
	log.Println("建立连接成功！")

	//打开
	f, err := os.Open(file)
	if err != nil {
		log.Printf("无法打开文件：%v", err)
		return
	}

	defer f.Close()

	//写入头文件
	conn.Write([]byte(file))

	p := make([]byte, 2)
	_, err = conn.Read(p)
	if err != nil {
		log.Printf("无法获取服务端信息%s", err)
		return
	} else if string(p) != "ok" {
		log.Printf("无法获取服务端相应%s", string(p))
		return
	}
	log.Println("头信息发送成功")

	_, err = io.Copy(conn, f)
	if err != nil {
		log.Printf("复制文件失败(%s):(%s)", file, err)
		return
	}
	log.Println("文件发送成功")
}
