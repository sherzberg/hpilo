//hpilo-test2.3.go
package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	//	"net/http"
	//	"os"
	//	"strings"
	//	"bytes"
	//	"io/ioutil"
)

type RibCl struct {
	XMLName  xml.Name `xml:"RIBCL"`
	Version  string   `xml:"VERSION,attr"`
	RibLogin []Login  `xml:"LOGIN"`
}
type Login struct {
	UserLogin string `xml:"USER_LOGIN,attr"`
	UserPass  string `xml:"PASSWORD,attr"`
	RibInfo   Info   `xml:"SERVER_INFO"`
}

//RibInfo
type Info struct {
	Mode       string `xml:"MODE,attr"`
	UidControl UIDC   `xml:"UID_CONTROL"`
}

//UIDControl
type UIDC struct {
	UID string `xml:"UID,attr"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	v := &RibCl{Version: "2.0"}
	v.RibLogin = append(v.RibLogin, Login{"Administrator", "password", Info{"write", UIDC{"True"}}})
	login_ribcl, err := xml.MarshalIndent(v, "", "    ")

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	URL := "ilomxq5420b9c:443"
	fmt.Println("Connecting to URL:>", URL)
	conn, err := tls.Dial("tcp", URL, conf)
	defer conn.Close()
	if err != nil {
		log.Fatalf("client: dial: %s", err)
		return
	}

	log.Println("client: connected to: ", conn.RemoteAddr())

	HTTP_HEADER := "POST /ribcl HTTP/1.1\r\nHost: localhost\r\nContent-Length: %d\r\nConnection: Close\r\n\r\n"
	xml_header := "<?xml version=\"1.0\"?>\r\n"
	msg_length := len(string(login_ribcl) + xml_header)

	msg := fmt.Sprintf(HTTP_HEADER, msg_length)

	log.Printf("headers: %s", msg)
	log.Printf("headers: %s", msg)

	body := []string{fmt.Sprintf(HTTP_HEADER, msg_length), xml_header, string(login_ribcl)}

	for _, data := range body {
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Panic(err)
		}
		log.Printf("client: write: %#v", data)
	}

	response_body, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("client response: %s", response_body)
}
