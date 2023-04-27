package handlers
import (
	"net/http"
	"fmt"
)

func Hola(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Hola Mundo")
}