//WarmBootServer.go
//The WARM_BOOT_SERVER command will force a warm boot of the server, if the server is currently on.
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
  WarmBootServer WBS `xml:"WARM_BOOT_SERVER"`
}
//WarmBootServer
type WBS struct {
  value string `xml:"WARM_BOOT_SERVER"`
}
func main() {
  v := &RibCl{Version: "2.0"}
  v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", WBS{""}}})
  output, err := xml.MarshalIndent(v,"  ","    ")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
  os.Stdout.Write([]byte(xml.Header))
  os.Stdout.Write(output)
}
