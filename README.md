goprop
======

read and writre properties file.(for study)

**demo**

```go
// Teststrconv project main.go
package main

import (
	"fmt"
	"goprop"
)

func main() {

	var p goprop.Properties
	var errOpen error
	p, errOpen = goprop.LoadFile("C:\\temp\\file.txt") // read properties file

	if errOpen != nil {
		fmt.Println("read err", errOpen)
		return
	}

	fmt.Println(p.Get("not exists"))       // error
	fmt.Println(p.Get("linestring"))       // string
	fmt.Println(p.Atoi("lineatoi"))        // int
	fmt.Println(p.ParseBool("linebool"))   // bool
	fmt.Println(p.ParseFloat("linefloat")) // float
	fmt.Println(p.ParseInt("lineint"))     // int
	fmt.Println(p.ParseUint("lineuint"))   // unit

	errSave := goprop.SaveFile(p, "C:\\temp\\file1.txt") // write properties file

	if errSave != nil {
		fmt.Println("write err", errSave)
		return
	}

}

```

**install**

```
go get -u github.com/aotian16/goprop
go install github.com/aotian16/goprop
```
