package main

import (
	"fmt"
	"log"
    "os"
    "strings"
)
func main (){
    input, err := os.ReadFile("input.txt");
    if( err != nil){
        log.Fatal(err)
    }
    lines  := strings.Split(string(input[:]),"\n");
    fmt.Println(lines[1])
    results := [] uint8 {};
    for i := 0; i< len(lines); i++ {
        for j:=0; j < len(lines[i]); j++ {
           results = append(results, lines[i][j]) 
        }
    }
}
