// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.

/*
$ cat exercises/ex1.11/top50.txt | xargs go run exercises/ex1.11/main.go
0.22s   146335  http://techtudo.com.br
0.25s   222247  http://uol.com.br
0.26s    91484  http://olx.com.br
0.26s      253  http://bb.com.br
0.28s   238770  http://folha.uol.com.br
0.30s    11137  http://google.com.br
0.30s   205521  http://bol.uol.com.br
0.30s    59997  http://itau.com.br
0.31s    11191  http://google.com
0.33s    38248  http://adf.ly
0.35s   681705  http://globo.com
0.39s    85411  http://correios.com.br
0.46s    82708  http://letras.mus.br
0.48s   134446  http://uptodown.com
0.51s   226169  http://otvfoco.com.br
0.67s    49423  http://msn.com
0.68s    39454  http://twitter.com
0.68s    59259  http://wordpress.com
0.69s     5939  http://reclameaqui.com.br
0.73s    13672  http://instagram.com
0.73s   164353  http://reddit.com
0.74s    84222  http://pinterest.com
0.75s   118964  http://fazenda.gov.br
0.76s    72472  http://tumblr.com
0.80s    52894  http://abril.com.br
0.81s    23240  http://twitch.tv
0.82s   209820  http://facebook.com
0.82s   112705  http://curapelanatureza.com.br
0.87s    35060  http://fatosdesconhecidos.com.br
1.05s     3295  http://onclkds.com
1.10s   113113  http://zipnoticias.com
1.11s    14038  http://popads.net
1.14s   156714  http://aliexpress.com
1.24s    43501  http://linkedin.com
1.32s    36530  http://whatsapp.com
1.62s   346724  http://americanas.com.br
1.64s   158617  http://microsoft.com
1.66s   218208  http://blastingnews.com
1.74s    15930  http://live.com
1.93s   302780  http://mercadolivre.com.br
1.97s    86413  http://wikipedia.org
1.99s    26423  http://bet365.com
2.21s   118200  http://onoticioso.com
2.33s   527104  http://youtube.com
2.34s   234805  http://metropoles.com
2.43s    98675  http://xvideos.com
2.55s   390832  http://yahoo.com
2.59s    55076  http://netflix.com
2.65s   117478  http://videodownloadconverter.com
Get http://caixa.gov.br/: stopped after 10 redirects
3.62s elapsed

Sites que não respondem resultam em tempo esgotado, erro de leitura ou algum outro erro.
A solução é definir explicitamente esses parâmetros.
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

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
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
