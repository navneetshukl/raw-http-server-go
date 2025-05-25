package main

import "os"

func loadHTML() ([]byte, error) {
	htmlData, err := os.ReadFile("./www/index.html")
	if err!=nil{
		return nil,err
	}
	return htmlData,err
}