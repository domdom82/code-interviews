package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//echo1()
	echo2()
}

func echo1() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			buf, err := ioutil.ReadAll(r.Body)
			fmt.Printf("Read %d bytes\n", len(buf))

			nn, err := w.Write(buf)
			fmt.Printf("Wrote %d bytes\n", nn)

			if err != nil {
				fmt.Println("ERROR (write): ", err)
			}
		}
	})

	http.ListenAndServe(":8080", mux)
}

func echo2() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			buf := make([]byte, 4096)

			for {
				n, err := r.Body.Read(buf)
				fmt.Printf("Read %d bytes\n", n)

				if n > 0 {
					nn, err := w.Write(buf[:n])
					fmt.Printf("Wrote %d bytes\n", nn)

					if err != nil {
						fmt.Println("ERROR (write): ", err)
						break
					}
				}

				if err != nil {
					break
				}
			}
		}
	})

	http.ListenAndServe(":8080", mux)
}
