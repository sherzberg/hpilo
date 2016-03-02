//BootMode.go
/*Current possible BUG
DocumentID: c04867696
Advisory: ProLiant Gen9 Servers - SET_PERSISTENT_BOOT (UEFI) XML Script Will Generate an "Invalid Device Choice" Error Message for Devices in UEFI Boot Mode
DESCRIPTION
On ProLiant Gen9 servers with Unified Extensible Firmware Interface (UEFI)boot mode, the SET_PERSISTENT_BOOT command takes one or more UEFI boot parameters and sets the normal boot order. However, if the server is running HP Integrated Lights-Out 4 (iLO 4) Firmware Version 2.22, the execution of this script results in a failure with the error message "Invalid Device choice" and the error status "0x0001."
SCOPE
Any HP ProLiant Gen9 server running HP Integrated Lights-Out 4 (iLO 4) Firmware Version 2.22.
RESOLUTION
A future version of iLO 4 firmware will correct this issue.
As a workaround, set the boot order for the UEFI devices through either the HP RESTful interface, the SMASH CLP, or the iLO GUI.
*/

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
  SetPendingBootMode SPBM `xml:"SET_PENDING_BOOT_MODE"`
}
//SetPendingBootMode
type SPBM struct {
  Value string `xml:"VALUE,attr"`
//options are "UEFI" or "LEGACY"
}
func main() {
  v := &RibCl{Version: "2.0"}
  v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", SInfo{"write", SPBM{"UEFI"}}})
  output, err := xml.MarshalIndent(v,"  ","    ")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
  os.Stdout.Write([]byte(xml.Header))
  os.Stdout.Write(output)
}
