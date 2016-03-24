//hpilo-test2.2.go
package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", UIDC{"Yes"}}})
	CustomTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: CustomTransport}
	output, err := xml.MarshalIndent(v,"  ","    ")
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(output)))
	// text/xml or application/xml?
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	resp, err := client.Do(req)	
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to reach the server: ", err)
	} 
	//else { 
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println("body=", string(body))
	//}
	defer resp.Body.Close()
	fmt.Println("input data sent: ", v)
	fmt.Println("request body: ", req.Body)
	fmt.Println("response body: ", resp.Body)
	fmt.Println("XML Structured Output: ")
	os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(output)
}
