package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	dirList, e := ioutil.ReadDir(".")

	if e != nil {
		fmt.Println("read dir error")
		return
	}

	var dirInfo string
	dirInfo += "tree:\n"
	for i, v := range dirList {
		var iden string
		if v.IsDir() {
			iden = "文件夹"
		} else {
			iden = "文件"
		}
		dirInfo += strconv.Itoa(i) + ": " + v.Name() + ", " + iden + ", " + v.Mode().String() + ", " + string(v.Size()) + ", " + v.ModTime().String() + "\n"
	}
	dirInfo += "\n"

	data, err := ioutil.ReadFile("./Dockerfile")
	if err != nil {
		data, err = ioutil.ReadFile("./HowToDeployToDocker/Dockerfile")
		if err == nil {
			fmt.Fprintf(w, "hello docker\n\n" + dirInfo + "\nDockerfile:\n" + string(data))
		}
	} else {
		fmt.Fprintf(w, "hello docker\n\n" + dirInfo + "\nDockerfile:\n" + string(data))
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}