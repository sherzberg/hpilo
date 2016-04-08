//hpilo-test2.3.go
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	URL := "ilomxq5420b9c:443"

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	fmt.Println("Connecting to URL:>", URL)
	conn, err := tls.Dial("tcp", URL, conf)
	defer conn.Close()
	if err != nil {
		log.Fatalf("client: dial: %s", err)
		return
	}

	HTTP_HEADER := "POST /ribcl HTTP/1.1\r\nHost: localhost\r\nContent-Length: %d\r\nConnection: Close\r\n\r\n"
	msg := "<RIBCL\r\n VERSION=\"2.0\"\r\n>\r\n<LOGIN\r\n PASSWORD=\"password\"\r\n USER_LOGIN=\"Administrator\"\r\n>\r\n<RIB_INFO\r\n MODE=\"read\"\r\n>\r\n<GET_FW_VERSION\r\n />\r\n</RIB_INFO>\r\n</LOGIN>\r\n</RIBCL>\r\n"
	xml_header := "<?xml version=\"1.0\"?>\r\n"

	body := []string{fmt.Sprintf(HTTP_HEADER, len(msg+xml_header)), xml_header, msg}

	for _, data := range body {
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Panic(err)
		}
		log.Printf("client: write %s", data)
	}

	response_data, _ := ioutil.ReadAll(conn)

	log.Printf("client response: %s", response_data)
}
