package errorutil

import (
	"fmt"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckErrorAsResponse(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
}
