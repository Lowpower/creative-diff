package main

import (
	"flag"
	"os/exec"
)

func GetResponse(host string, filename string) []byte {
	server_info := "-server=" + host
	request_file := "-request_file=./diff-data/" + filename
	cmd := exec.Command("./client_creative_search", server_info, request_file)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	srcHost := flag.String("srchost", "127.0.0.1", "请输入src地址")
	dstHost := flag.String("dsthost", "127.0.0.1", "请输入dst地址")
	filename := "test1"
	GetResponse(*srcHost, filename)
	GetResponse(*dstHost, filename)
}
