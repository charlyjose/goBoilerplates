/*
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	var netTransport = &http.Transport{
		DialTLS: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
		TLSHandshakeTimeout: time.Second * 5,
	}

	var netClient = &http.Client{
		Transport: netTransport,
		Timeout:   time.Second * 10,
	}

	response, err := netClient.Get("https://microsoft.com")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Print(response)
}
*/

// Go contains rich function for grab web contents. _net/http_ is the major
// library.
// Ref: [golang.org](http://golang.org/pkg/net/http/#pkg-examples).
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// keep first n lines
func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {
	// We can use GET form to get result.
	resp, err := http.Get("http://g.cn/robots.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get:\n", keepLines(string(body), 3))

	fmt.Println("\n\n\n\n")

	// We can use POST form to get result, too.
	resp, err = http.Get("https://duckduckgo.com/?q=wow&ia=definition")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.OpenFile("duckSearch.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body, err = ioutil.ReadAll(resp.Body)
	l, err := file.WriteString(string(body))
	if err != nil {
		log.Fatal(err)
		file.Close()
	}
	fmt.Println("Bytes Written: ", l, "\n\n")
	fmt.Println("post:\n", string(body))

}
