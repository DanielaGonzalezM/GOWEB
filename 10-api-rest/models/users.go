package models

import "apirest/db"

//Si se quere trabajar con xml se debe reemplazar `json:"id"` por `xml:"id"`

//si se quiere trabajar con yaml, este debe ser instalado y reemplazar `json:"id"` por `yaml:"id"`
//https://github.com/go-yaml/yaml
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

//Constructor usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

//Crear usuario e inserta db
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

//Insertar Registro
func (user *User) insert() {
	sql := "INSERT users SET username=?,password=?,email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

//Lista de usuarios
func ListUsers() (Users, error) {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, error := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, error
}

//Obtener un registro

func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	if rows, error := db.Query(sql, id); error != nil {
		return nil, error
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

		}
		return user, nil
	}
}

//Actualizar registro
func (user *User) update() {
	sql := "UPDATE users SET username=?,password=?,email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

//Guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

//Borrar registro
func (user *User) DeleteUser() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)

}
