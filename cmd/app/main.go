package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func HandleSubscribePage(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		file, err := os.Open("./../../public/index.html")
		if err != nil {
			writer.Write([]byte("Can't get main page"))
			return
		}
		io.Copy(writer, file)
	case http.MethodPost:
		request.ParseMultipartForm(10000)
		_ = request.MultipartForm.Value["name"][0]
		_ = request.MultipartForm.Value["mail"][0]

	}
}

func main() {
	http.HandleFunc("/page", HandleSubscribePage)
	fsHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/", fsHandler)
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
