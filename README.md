# hpilo

Hewlett Packard iLO4 golang library for scripting bare metal without HP OneView

## Usage

```go
//main.go
package main

import "fmt"
import "github.com/sherzberg/hpilo"

func main() {
    client := hpilo.NewIloClient("hostname-0", "username", "password")

    fw_version, err := client.GetFwVersion()
    if err != nil {
        panic(err)
    }

    fmt.Println(fw_version)
}
```
