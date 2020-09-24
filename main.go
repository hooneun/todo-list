package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./views")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)

	// db, _ := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
	// defer db.Close()

	// var version string
	// db.QueryRow("SELECT VERSION()").Scan(&version)

	// fmt.Println("Connect to:", version)
	// dns := "root@tcp(127.0.0.1:3306)/todo_list?charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to db ")
	// }

	// // db.AutoMigrate(&User{})
	// db.AutoMigrate(&models.User{})

	// db.Create(&models.User{
	// 	Name: "HoonEun", Email: "kimhoon0316@gmail.com",
	// 	Password: "12341234",
	// })

	// var user models.User
	// db.First(&user, 1)
	// fmt.Println(user)
	// db.Delete(&user)
}
