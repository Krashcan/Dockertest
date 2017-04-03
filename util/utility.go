package util

import(
	"time"
	"net/http"
	"log"
)

func Logger( inner http.Handler, name string ) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf("%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		name,
		time.Since(start))
	})
}


func ErrorExists(err error,code int,w http.ResponseWriter) bool{
	if err!=nil{
		log.Print(err)
		http.Error(w,err.Error(),code)
		return true	
	}
	
	return false
}