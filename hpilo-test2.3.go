//hpilo-test2.3.go
package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
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
	RibInfo   Info   `xml:"RIB_INFO"`
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
	URL := "https://ilomxq5420b9c:443"
	fmt.Println("Connecting to URL:>", URL)
	v := &RibCl{Version: "2.0"}
	v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", UIDC{"True"}}})
	output, err := xml.MarshalIndent(v,"  ","    ")
	log.SetFlags(log.Lshortfile)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "ilomxq5420b9c:443", conf)
	if err != nil {
		log.Println(err)
		log.Fatalf("client: dial: %s", err)
		return
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
//	state := conn.ConnectionState()
	n, err := conn.Write(output)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("client: write: %s", err)	
	reply := make([]byte, 256)
	n, err =conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}
