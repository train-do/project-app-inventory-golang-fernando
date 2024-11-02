package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func DecodeFromJSON(req interface{}) {
	// fmt.Printf("%s\n%+v\n", reflect.TypeOf(req).Name(), req)
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// fmt.Printf("%v\n%v\n", reflect.TypeOf(req).Name(), req)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(req); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// fmt.Printf("After Decode :%s\n%+v\n", reflect.TypeOf(req).Name(), req)
}
