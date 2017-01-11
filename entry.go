package main

import "net/http"


type Myhandler struct{}

func (m *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is working!!!\n"))
}




func main(){

	handler := &Myhandler{}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	http.ListenAndServe(":8090", mux)


}

