//eject-virtual-media.go
package main
import (
	"encoding/xml"
	"fmt"
	"os"
)
type RibCl struct {
	XMLName xml.Name `xml:"RIBCL"`
	Version string	`xml:"VERSION,attr"`
	RibLogin []Login `xml:"LOGIN"`
}
type Login struct {
	UserLogin string `xml:"USER_LOGIN,attr"`
	UserPass string `xml:"PASSWORD,attr"`
	RibInfo Info `xml:"RIB_INFO"`
}
type Info struct {
	Mode string `xml:"MODE,attr"`
	EjectVirtualMedia EVM `xml:"EJECT_VIRTUAL_MEDIA"`
}
type EVM struct {
	Device string `xml:"DEVICE,attr"`
}
func main() {
    v := &RibCl{Version: "2.0"}
	v.RibLogin = append(v.RibLogin, Login{"Administrator", "password123", Info{"write", EVM{"CDROM"}}})
	output, err := xml.MarshalIndent(v,"  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
