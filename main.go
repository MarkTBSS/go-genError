package main

import "errors"

func genError() error {
	return errors.New("error")
}
func main() {
	if err := genError(); err != nil {
		panic(err)
	}
}
