package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

mux := http.NewServeMux()
mux.HandleFunc("/", handlerr)

fmt.Println("starting server at 8080")
http.ListenAndServe(":8080", mux)

}
func handlerr(w http.ResponseWriter, r *http.Request) {
c, err := r.Cookie("lastvisit")

expiry := time.Now().AddDate(0,0,1) // a day later

cookie := &http.Cookie{
	Name: "lastvisit",
	Expires: expiry,
	Value: strconv.FormatInt(time.Now().Unix(), 10),
	HttpOnly: true,
}
http.SetCookie(w, cookie)
if err != nil {
	w.Write([]byte("welcome to the site!"))
  return
} else {
	lasttime, _ := strconv.ParseInt(c.Value, 10, 0)
	html := "welcome back! you last visited at"
	html = html + time.Unix(lasttime, 0).Format("15:04:05")
	w.Write([]byte(html))
}
// w.Write([]byte("hello"))


}
