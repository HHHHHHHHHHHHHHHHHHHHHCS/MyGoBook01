package main

import (
	"fmt"
)

func Test2(arr *[]int)  {
	a:=make([]int,10)
	for i:=0;i<len(a);i++ {
		a[i]=i+100
	}
	*arr=append(*arr, a...)
}

func Test1(arr *[]int)  {
	a:=make([]int,10)
	for i:=0;i<len(a);i++ {
		a[i]=i
	}
	*arr=append(*arr, a...)
	Test2(arr)
}

func main() {
	arr := make([]int,0)
	Test1(&arr)
	for k,v:=range arr{
		fmt.Println(k,v)
	}

}
