package main

import (
	"../MusicLib"
	"fmt"
)

func main() {
	a := MusicLib.MusicEntry{"1", "Will", "AAA", "www.baidu.com", "MP3"}
	fmt.Printf("%+v\n", a)
}
