// Date: 18-2-9

package main

import (
	"fmt"

	"github.com/admpub/marmot/miner"
)

func PostFile(filename string, targetUrl string) {
	worker, _ := miner.New(nil)
	result, err := worker.SetURL(targetUrl).SetBin([]byte("dddd")).SetFileInfo(filename+".xxxx", "uploadfile").PostFile()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(result))
	}
}

// sample usage
func main() {
	target_url := "http://127.0.0.1:1789/upload"
	filename := "./doc.go"
	PostFile(filename, target_url)
}
