package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html>hello</html>"))
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xhtml+xml")
	w.Write([]byte(`<note>
	<to>Tove</to>
	<from>Jani</from>
	<heading>Reminder</heading>
	<body>Don't forget me this weekend!</body>
	</note>`))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/xml", xmlHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Printf("Failed to start http server: %s\n", err)
	}
}
