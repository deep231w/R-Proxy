package testserver

import (
	"fmt"
	"net/http"
)


func main (){
	mux := http.NewServeMux();

	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the home page!")
	})


	fmt.Println("server running on port 9000");

	err:=http.ListenAndServe("9000", mux)

	if err!=nil{
		fmt.Println("error appeared during start of server:" ,err);
	}
}