package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range data {
		fmt.Println(k, v)
	}
}
