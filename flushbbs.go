package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/jinzhu/gorm"
)

func index(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(20 << 20)
	file, handler, err:= r.FormFile("dd")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	new_file, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	expireotion := time.Now()
	expireotion = expireotion.AddDate(1, 0, 0)
	cookie := http.Cookie{Name:"name", Value:"cty", Expires:expireotion}
	http.SetCookie(w, &cookie)

	fmt.Println(r.Cookie("name"))

	fmt.Fprintf(w, "%v", handler.Header)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer  new_file.Close()

	io.Copy(new_file, file)



	for a, b := range r.Form {
		fmt.Println(a, b)
	}

	fmt.Fprint(w, "hello")
}

type man struct {
	uid int
	name string
}
func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8");

	a := db.AutoMigrate(&man{})

	fmt.Println(a.Error.Error())
	http.HandleFunc("/", index)
	err = http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func print_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
