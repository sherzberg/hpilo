//HoldPwrBtn.go
// This HOLD_PWR_BTN command is used to simulate a physical press and hold of the server power button.
package main
import (
  "encoding/xml"
  "fmt"
  "os"
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
  HoldPwrBtn HPB `xml:"HOLD_PWR_BTN"`
}
//HoldPowerButton
type HPB struct {
  value string `xml:"HOLD_PWR_BTN"`
}
func main() {
  v := &RibCl{Version: "2.0"}
  v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", HPB{""}}})
  output, err := xml.MarshalIndent(v,"  ","    ")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
  os.Stdout.Write([]byte(xml.Header))
  os.Stdout.Write(output)
}