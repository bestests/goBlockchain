package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func main () {
	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("Call!!!")
		w.Write([]byte("<!DOCTYPE html><html lang='ko'><head><meta charset='UTF-8' /><title>Hello, World</title></head><body><h1>Hello, World!!!</h1><hr /></body></html>"))
	})

	http.HandleFunc("/hello", func (w http.ResponseWriter, req *http.Request) {
		fi, err := os.Open("./views/html/test.html")

		if err != nil {
			panic(err)
		}

		defer fi.Close()

		buff := make([]byte, 1024)

		for {
			cnt, err := fi.Read(buff)

			if err != nil && err != io.EOF {
				panic(err)
			}

			if cnt == 0 {
				break
			}

			w.Write(buff[:cnt])
		}
	})

	http.ListenAndServe(":5000", nil)
}
