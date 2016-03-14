//hpilo-test-connect.go
package main

import (
	"net/http"
	"encoding/xml"
	"bytes"
	"fmt"
	"io/ioutil"
)

type RibCl struct {
  XMLName xml.Name `xml:"RIBCL"`
  Version string  `xml:"VERSION,attr"`
  RibLogin []Login `xml:"LOGIN"`
}
type Login struct {
  UserLogin string `xml:"USER_LOGIN,attr"`
  UserPass string `xml:"PASSWORD,attr"`
  RibInfo Info `xml:"RIB_INFO"`
}
//RibInfo
type Info struct {
  Mode string `xml:"MODE,attr"`
  UidControl UIDC `xml:"UID_CONTROL"`
}
//UIDControl
type UIDC struct {
  value string `xml:"UID_CONTROL"`
}

func main() {
   v := &RibCl{Version: "2.0"}
   v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", UIDC{"Yes"}}})
	uid := &UIDC
	buf, _ := xml.Marshal(uid)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://10.28.100.10:443", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
//	resp, err := http.Get(url)
//resp.Header.Add("Accept", "application/xml")
//resp.Header.Add("Content-Type","application/xml; charset=utf-8")
//if err != nil {
//     return "", err
}