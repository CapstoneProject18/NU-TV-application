package main

import (
	"fmt"
	"update"
)

func main() {
	t, reqHead, err := update.UpdateNUtv()
	if err != nil {
		_ = fmt.Errorf("this is something unwanted %v", err)
	}
	fmt.Println(t, reqHead)

}
