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
//	"bytes"
	"io/ioutil"
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

//func (c *Client) Do(req *Request) (resp *Response, err error)
//func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
//func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
	

func main() {
	URL := "https://ilomxq5420b9c:443"
	fmt.Println("Connecting to URL:>", URL)
	v := &RibCl{Version: "2.0"}
	v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", UIDC{"True"}}})
	CustomTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: CustomTransport}
	output, err := xml.MarshalIndent(v,"  ","    ")
//	body := bytes.NewBuffer(output)
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(output)))
//  req, err := http.NewRequest("POST", URL, strings.NewReader("output"))
//	req, err := http.NewRequest("POST", URL, bytes.NewBufferString("output"))
//	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(output))
//	req, err := http.NewRequest("POST", URL, body)
//	req, err := http.NewRequest("POST", URL, xml.NewParser(v))
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

	if b, err := ioutil.ReadAll(resp.Body); err == nil {
    fmt.Println(string(b))
}
	fmt.Println(resp.Status)
//	fmt.Printf("%#v", req.Body)
	fmt.Println("request body: ", req.Body)
	fmt.Println("response body: ", resp.Body)
	fmt.Println("XML Structured Output: ")
	os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(output)
}
