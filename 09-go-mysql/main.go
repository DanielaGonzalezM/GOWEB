package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistTable("users"))
	db.CreateTable(models.UserSchema, "users")
	db.Ping()
	db.TruncateTable("users")
	user := models.CreateUser("Daniela5", "123", "daniela@gmail.com")
	users := models.ListUsers()
	fmt.Println(users)
	fmt.Println(user)

	user = models.GetUser(4)
	fmt.Println(user)
	user.Username = "Dany"
	user.Save()
	user.DeleteUser()
	db.Close()

}
