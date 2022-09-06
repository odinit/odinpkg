package main

import (
	"fmt"
	"github.com/odinit/global/gtype"
)

func main() {
	a := map[string]string{"a": "b", "c": "d"}
	b := gtype.MSS(a)
	fmt.Println(J(b, "=", "&"))
}

func J(a gtype.GMap, sep1, sep2 string) string {
	return a.Join(sep1, sep2)
}
