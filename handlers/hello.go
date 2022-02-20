package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DataHandler(rw http.ResponseWriter, r *http.Request) {
	thisLogger := log.New(os.Stdout, "product-api", log.LstdFlags)
	thisLogger.Println("Hello World")

	dataBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		/*responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Some error have come in."))
		return*/
		http.Error(rw, "Some error have come in.", http.StatusBadRequest)
		return
	}

	//log.Printf("Data %s\n", dataBody)
	thisLogger.Printf("Data %s\n", dataBody)

	fmt.Fprintf(rw, "Data thus received in Request is : %s\n", dataBody)

}
