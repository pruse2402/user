package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"user/users/dbcon"
	"user/users/internals"
	"user/users/models"

	"github.com/julienschmidt/httprouter"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	if !parseJSON(w, r.Body, &user) {
		return
	}
	db := dbcon.Get()

	user.CreatedDate = time.Now().UTC()

	if _, err := db.Model(&user).Insert(); err != nil {
		e := &models.ErrorData{}
		e.Message = "Error in save user"
		e.Err = err
		e.IsDbErr = true
		renderERROR(w, e)
		return
	}

	res := struct {
		Message string       `json:"message"`
		User    models.Users `json:"user"`
	}{}

	res.Message = "User saved successfully"
	res.User = user
	renderJSON(w, http.StatusOK, res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, isErr := getIDFromParams(w, r, "id")

	if !isErr {
		return
	}
	user := models.Users{LegalEntityID: id}
	db := dbcon.Get()

	if err := db.Select(&user); err != nil {
		e := &models.ErrorData{}
		e.Message = "User not found"
		e.Err = err
		renderERROR(w, e)
		return
	}

	if !parseJSON(w, r.Body, &user) {
		return
	}

	if err := internals.UpdateUser(db, &user); err != nil {
		e := &models.ErrorData{}
		e.Message = "Error while updating user"
		e.Err = err
		e.IsDbErr = true
		renderERROR(w, e)
		return
	}

	res := struct {
		Message string       `json:"message"`
		User    models.Users `json:"user"`
	}{}
	res.Message = "Updated successfully"
	res.User = user

	renderJSON(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value("params").(httprouter.Params)

	idStr := params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	db := dbcon.Get()

	user := models.Users{}

	if err := db.Model(&user).Where("users.id=?", id).Select(); err != nil {
		e := &models.ErrorData{}
		e.Message = "User not found"
		e.Err = err
		renderERROR(w, e)
		return
	}

	renderJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params, _ := r.Context().Value("params").(httprouter.Params)

	res := make(map[string]interface{})

	idStr := params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	db := dbcon.Get()
	user := models.Users{}

	if _, err := db.Model(&user).Where("legalEntityId=?", id).Delete(); err != nil {
		e := &models.ErrorData{}
		e.Message = "Unable to Delete"
		e.Err = err
		renderERROR(w, e)
		return
	}

	res["message"] = "User deleted successfully"

	renderJSON(w, http.StatusOK, user)

}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	id, isErr := getIDFromParams(w, r, "id")
	if !isErr {
		return
	}

	res := make(map[string]interface{})

	params, _ := ListQueryParams(r)

	db := dbcon.Get()

	user := models.Users{}

	queryStr := fmt.Sprintf(`SELECT * FROM users as u WHERE u.legal_entity_id = %v`, id)

	if params.CompanyName != "" {
		companyNamequeryStr := fmt.Sprintf("AND u.company_name = '%v' ", params.CompanyName)
		queryStr = queryStr + companyNamequeryStr

	}

	if params.FirstName != "" {
		firstNamequeryStr := fmt.Sprintf("AND u.first_name = '%v' ", params.FirstName)
		queryStr = queryStr + firstNamequeryStr

	}

	if params.LastName != "" {
		lastNamequeryStr := fmt.Sprintf("AND u.last_name = '%v' ", params.LastName)
		queryStr = queryStr + lastNamequeryStr

	}

	if params.LegalEntityID > 0 {
		legalEntityIDqueryStr := fmt.Sprintf("AND u.legal_entity_id = '%v' ", params.LegalEntityID)
		queryStr = queryStr + legalEntityIDqueryStr
	}

	if _, err := db.Query(&user, queryStr); err != nil {
		e := &models.ErrorData{}
		e.Message = "Unable to get the list"
		e.Err = err
		renderERROR(w, e)
		return
	}

	res["user"] = user
	res["message"] = "fetched the user details successfully"

	renderJSON(w, http.StatusOK, res)
}
