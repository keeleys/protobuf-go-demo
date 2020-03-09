package main

import (
	b "hello/proto"

	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	var b1 = &b.Book{Name: "go语言开发", Author: &b.Book_Author{Name: "张三"}}
	protoBook, err := proto.Marshal(b1)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	log.Printf("序列号之后的字节长度 %d\n", len(protoBook))
	b2 := &b.Book{}
	if err := proto.Unmarshal(protoBook, b2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	log.Printf("book : %s, author: %s\n", b2.GetName(), b2.GetAuthor().GetName())
}
