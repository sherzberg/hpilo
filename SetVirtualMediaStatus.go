//SetVirtualMediaStatus.go
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
  SetVMStatus SVMS `xml:"SET_VIRTUAL_MEDIA"`
}
//SetVirtuaMediaStatus
type SVMS struct {
  Device string `xml:"DEVICE,attr"`
  VmBootOption VMBO `xml:"VM_BOOT_OPTION"`
  VmWriteProtect VMWP `xml:"VM_WRITE_PROTECT"`
}
//VirtualMediaBootOption
type VMBO struct {
  Value string `xml:"value,attr"`

}
//VirtualMediaWriteProtect
type VMWP struct {
  Value string `xml:"value,attr"`
}
func main() {
  v := &RibCl{Version: "2.0"}
  v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", SVMS{"CDROM", VMBO{"BOOT_ONCE"}, VMWP{"Y"}}}})
  output, err := xml.MarshalIndent(v,"  ","    ")
  if err != nil {
    fmt.Printf("error: %v\n", err)
  }
  os.Stdout.Write([]byte(xml.Header))
  os.Stdout.Write(output)
}
//  BootOption string `xml:"value,attr"`
//  WriteProtect string `xml:"value,attr"`