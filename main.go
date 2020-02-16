package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	//filecrawl
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	//md5sum
	//hash, err := md5sum(os.Args[1])
	hash, err := md5sum(visit)
	if err == nil {
		fmt.Println(hash)
	}

}
