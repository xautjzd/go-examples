package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

var fileName = "test.md"

func main() {
	// 常用写文件
	if err := ioutil.WriteFile(fileName, []byte("## title\n content"), 0644); err != nil {
		log.Fatal(err)
	}

	// 以读写方式打开文件
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("file: %s doesn't exist", fileName)
	}
	defer f.Close()

	// 无缓冲同步写，持久化存储后返回
	n, err := f.Write([]byte("# First Level\n## Second Level\n### Third Level\nmain content"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("write %d bytes to file: %s", n, fileName)

	// 带缓冲写, 写入缓冲区后即返回, 系统定时刷新缓冲区数据至磁盘
	writer := bufio.NewWriter(f)
	// 若未显式刷盘，程序异常退出时数据丢失; 若程序运行过程中被kill, defer 不会执行, 数据会丢失
	defer writer.Flush()
	if _, err := writer.WriteString("hello world"); err != nil {
		log.Fatalf("buffer write err: %v", err)

	}

	// 读文件
	if _, err := f.Seek(3, 0); err != nil {
		log.Fatalf("file seek err: %v", err)
	}
	content := make([]byte, 100)
	rn, err := f.Read(content)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("readed content: %s, total bytes: %d", content, rn)

	// 检查文件是否存在
	if _, err := os.Stat(fileName); err == nil {
		log.Printf("file: %s exists, delete it\n", fileName)
		if err := os.Remove(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
