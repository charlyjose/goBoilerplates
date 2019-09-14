package main

import (
	// "bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sort := flag.String("s", "s", "")
	view := flag.String("v", "s", "")
	help := flag.Bool("h", false, "")

	flag.Parse()

	fmt.Println("Sort: ", *sort)
	fmt.Println("View: ", *view)
	fmt.Println("Help: ", *help)
	fmt.Println("Tail: ", flag.Args())
	if *sort == "" {
		flag.PrintDefaults()
	}
	if *view == "" {
		flag.PrintDefaults()
	}
	if *help {
		file, err := os.OpenFile("help.txt", os.O_RDONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// scanner := bufio.NewScanner(file)
		// for scanner.Scan() {
		// 	fmt.Println(scanner.Text())
		// }

		b := make([]byte, 1)
		for {
			n, err := file.Read(b)

			if n > 0 {
				fmt.Print(b[:n], " ", string(b[:n]))
			}

			if err == io.EOF {
				break
			}

			fmt.Printf("\tRead %d bytes: %v \n", n, err)
			if err != nil {
				log.Printf("Read %d bytes: %v \n", n, err)
				break
			}
		}
		// b, err := ioutil.ReadAll(file)
		// fmt.Print(b)
	}
}
