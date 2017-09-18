// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.

/*
$ go run exercises/ex1.10/main.go https://tools.ietf.org/html/rfc2616 https://time.is/
0.69s    25866  https://time.is/
1.68s   537381  https://tools.ietf.org/html/rfc2616
1.68s elapsed

results: C:\Users\Rafael\AppData\Local\Temp

$ go run exercises/ex1.10/main.go https://tools.ietf.org/html/rfc2616 https://time.is/
0.67s    25866  https://time.is/
1.64s   537381  https://tools.ietf.org/html/rfc2616
1.64s elapsed

results: C:\Users\Rafael\AppData\Local\Temp

Executando sucessivamente, o tempo de execução não sofre nenhuma alteração significante.
O conteúdo do site também é sempre o mesmo, exceto em casos que o site gere conteúdo dinâmico.
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

var reURL = regexp.MustCompile("[^[:word:]]")

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("\nresults: %s\n", os.TempDir())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// create a temporary filename to store results
	f, err := ioutil.TempFile("", "fetchall-")
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s: %v", url, err)
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
