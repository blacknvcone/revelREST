package controllers

import (
	"RevelREST/app/database"
	"RevelREST/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type UserControllerCtx struct {
	*revel.Controller
}

func (c UserControllerCtx) IndexUser() revel.Result {
	results := models.User{}
	if err := database.Users.Find(bson.M{}).All(&results); err != nil {
		log.Fatal(err)
	}

	return c.RenderJSON(results)
}

func (c UserControllerCtx) CreateUser() revel.Result {
	user := &models.User{}
	if body, err := ioutil.ReadAll(c.Request.GetBody()); err != nil {
		return c.RenderText("bad request")
	} else if err := json.Unmarshal(body, user); err != nil {
		var jsonData map[string]interface{}
		fmt.Println(c.Params.BindJSON(&jsonData))
		return c.RenderText("could not parse request")
	}

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(user)
}
