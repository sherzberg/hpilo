//SetOneTimeBoot.go
// sets one time boot device on hardware; cdrom, floppy, network, etc. 
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
  ServerInfo SInfo `xml:"SERVER_INFO"`
}
//ServerInfo
type SInfo struct {
  Mode string `xml:"MODE,attr"`
  SetOneTimeBoot SOTB `xml:"SET_ONE_TIME_BOOT"`
}
//SetPendingBootMode
type SOTB struct {
  Value string `xml:"VALUE,attr"`
//options are "CDROM, "NETWORK", "HDD", "USB", or "RBSU"
}
func main() {
  v := &RibCl{Version: "2.0"}
  v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", SInfo{"write", SOTB{"CDROM"}}})
  output, err := xml.MarshalIndent(v,"  ","    ")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
  os.Stdout.Write([]byte(xml.Header))
  os.Stdout.Write(output)
}
