package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.FormValue("x"))
		fmt.Println(request.FormValue("y"))
		fmt.Println(request.PostFormValue("y"))
		//if file, fileHeader, err := request.FormFile("z"); err == nil {
		//	fmt.Println(fileHeader.Filename, fileHeader.Size)
		//	io.Copy(os.Stdout, file)
		//}
	})
	http.ListenAndServe("0.0.0.0:8888", nil)
}
