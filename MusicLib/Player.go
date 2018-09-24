package MusicLib

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source,fileType string){
	var p Player

	switch fileType {
	case "mp3":
		p=&MP3Player{}
	case "wav":
		p=&WavPlayer{}
	default:
		fmt.Println("Unknow type:",fileType)
		return
	}
	fmt.Println(p)
	p.Play(source)
}