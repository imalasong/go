package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"log"
	"net/http"
	"sync"
)

func main() {
	//testChan()
	//go1()
	//go2()
}

func go2() {

	src := []byte(`package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
`)

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}

func go1() {
	src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s  %-8s   %q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}

func testChan() {

	task := make(chan int, 10)
	collect := make(chan int)

	// init task
	go func() {
		for i := 0; i < 10; i++ {
			task <- i
		}
		close(task)
	}()

	var wg sync.WaitGroup

	for t := range task {
		wg.Add(1)
		go func(t int) {
			fmt.Println("写：", t)
			defer wg.Done()
			collect <- t * 10
		}(t)
	}

	go func() {
		wg.Wait()
		close(collect)
	}()

	for c := range collect {
		fmt.Printf("消费:%d\n", c)
	}

}

func testRecover() {
	defer func() {
		recover()
	}()
	fmt.Println("hello go complier")

	var wx sync.WaitGroup

	wx.Add(1)
	go func() {
		fmt.Println("go function exit 1")
		server()
		defer wx.Done()
		fmt.Println("go function exit 2")
		fmt.Println("go function exit 3")
	}()

	wx.Wait()
	fmt.Println("main exit")
}

func server() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

	})
	err := http.ListenAndServe("localhost:8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}))
	if err != nil {
		fmt.Println("server start error :", err)
	}

}
