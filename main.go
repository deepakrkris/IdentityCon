package main

import "IdentityCon/lib"


func main() {
    params := map[string]string {
		"name" : "foobar",
		"email" : "foobar@gmail.com",
		"area" : "625009",
	}
	lib.GenerateIdenticon(params, 256, 256)
}
