package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func handler(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("获取IP %s", remoteAddr)

	//获取文件头信息
	p := make([]byte, 1024)
	//p2 := make([]byte,0,1024)
	n, err := conn.Read(p)
	if err != nil {
		log.Printf("获取文件头信息失败(%s):(%s)", remoteAddr, err)
		return
	} else if n == 0 {
		log.Printf("空文件头(%s)", remoteAddr)
		return
	}

	filename := string(p[:n])
	log.Printf("file: %s", filename)
	conn.Write([]byte("ok"))

	//打开一个本地文件
	os.MkdirAll("receice",os.ModePerm)
	//创建一个本地文件
	f, err := os.Create("receice/" + filename)
	if err != nil {
		log.Printf("无法创建文件[%s]:[%s]", remoteAddr, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, conn)
	for {
		buffer := make([]byte, 1024*200)
		_, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Printf("获取失败(%s):(%s)", remoteAddr, err)
		} else if err == io.EOF {
			break
		}
	}
	if err != nil {
		log.Printf("文件接收失败(%s)", remoteAddr, err)
		return
	}
	log.Printf("文件接收成功(%s):(%s)", remoteAddr, filename)
}

func runServer(port int) {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("服务监听失败: %s", err)
	}

	log.Println("服务已启动")
	for {
		conn, err := l.Accept()
		if err != nil {
			if ni, ok := err.(net.Error); !ok || !ni.Temporary() {
				log.Printf("接受请求失败: %v", err)
			}
			continue
		}
		log.Println("打印成功%s", conn)
		go handler(conn)
	}
}
