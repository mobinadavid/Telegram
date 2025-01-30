package main

import "telegram/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
