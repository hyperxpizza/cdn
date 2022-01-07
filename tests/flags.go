package tests

import "flag"

var configPath = flag.String("path", "", "path to config.json file")
var secure = flag.Bool("secure", false, "secure grpc connection witr tls")
