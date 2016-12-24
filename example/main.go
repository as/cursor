package main

import(
	"github.com/as/cursor"
	"os"
	"strconv"
	"image"
	"fmt"
)

func atoi(a string) (i int){
	x, err := strconv.Atoi(a)
	if err != nil{
		panic(err)
	}
	return int(x)
}

func main(){
	x := atoi(os.Args[1])
	y := atoi(os.Args[2])
	if !cursor.MoveTo(image.Pt(x,y)){
		fmt.Println("failed to cursor")
	}
}
