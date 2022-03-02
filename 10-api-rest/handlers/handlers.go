package handlers

import (
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	/*	//Para trabajar con xml
		//rw.Header().Set("Content-Type", "text/xml")
		//Para trabajar con yaml no existe content-type especifico por lo que no se coloca
		rw.Header().Set("Content-Type", "application/json")

		users, _ := models.ListUsers()

		//Para trabajar con xml
		//output, _ := xml.Marshal(users)
		//Para trabajar con yaml
		//output, _ := yaml.Marshal(users)
		output, _ := json.Marshal(users)
		fmt.Fprintln(rw, string(output))
	*/
	//*********************************

	if users, error := models.ListUsers(); error != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	/*
		rw.Header().Set("Content-Type", "application/json")

		//Obtener ID
		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["id"])

		user, _ := models.GetUser(userId)

		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))
	*/
	//********************

	if user, error := getUserByRequest(r); error != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}

}
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	/*
		rw.Header().Set("Content-Type", "application/json")

		//Obtener Registro
		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {

			user.Save()

		}

		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))
	*/
	//***************************
	//Obtener Registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}

}
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	/*
		rw.Header().Set("Content-Type", "application/json")

		//Obtener Registro
		user := models.User{}

		//Obtener ID
		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["id"])
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {

			fmt.Println(user.Id)
			if user.Id == 0 {
				user.Id = int64(userId)
			}

			fmt.Println(user.Id)
			user.Save()

		}

		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))
	*/
	//****************************
	var userId int64
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		userId = user.Id
	}
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)

	}
}
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	/*
		rw.Header().Set("Content-Type", "application/json")

		//Obtener ID
		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["id"])

		user, _ := models.GetUser(userId)
		user.DeleteUser()

		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))
	*/
	//***********************
	if user, error := getUserByRequest(r); error != nil {
		models.SendNotFound(rw)
	} else {
		user.DeleteUser()
		models.SendData(rw, user)
	}

}
func getUserByRequest(r *http.Request) (models.User, error) {
	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}

}
