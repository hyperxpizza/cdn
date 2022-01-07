package tests

import (
	"flag"
	"fmt"
)

var configPath = flag.String("path", "", "path to config.json file")
var secure = flag.Bool("secure", false, "secure grpc connection witr tls")
var filePath = flag.String("filePath", "", "path to file for upload")
var bucket = flag.String("bucket", "", "bucket inside the filebrowser")

func printFlags() {
	fmt.Printf("configPath = %s\n", *configPath)
	fmt.Printf("secure = %v\n", *secure)
	fmt.Printf("filePath = %s\n", *filePath)
	fmt.Printf("bucket = %s\n", *bucket)
}
