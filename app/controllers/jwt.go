package controllers

import (
	"RevelREST/app/models"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	bcrypt "github.com/sensu/sensu-go/backend/authentication/bcrypt"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type JwtCtx struct {
	*revel.Controller
}

func Auth(c *revel.Controller) revel.Result {

	tokenString := c.Request.Header.Get("Authorization")
	if tokenString != "" {
		errResp := buildErrResponse(errors.New("Unauthorized!"), "401")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, fmt.Errorf("Unexpected signin method: %v", t.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		errResp := buildErrResponse(errors.New("Authorized !"), "200")
		c.Response.Status = 200
		return c.RenderJSON(errResp)
	} else {
		errResp := buildErrResponse(errors.New("Unauthorized!"), "401")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

}

func (c JwtCtx) GetToken() revel.Result {

	var err error
	var cred Credential
	var user models.User
	var token Token

	err = c.Params.BindJSON(&cred)
	if err != nil {
		errResp := buildErrResponse(errors.New("Bad Request !"), "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	//Validate Username
	user, err = models.ValidateUser(cred.Username)
	if err != nil {
		errResp := buildErrResponse(errors.New("User Not Found !"), "403")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	//Validate Password
	cek := bcrypt.CheckPassword(user.Password, cred.Password)
	fmt.Println(cek)
	if !cek {
		errResp := buildErrResponse(errors.New("Unauthorized ! "), "401")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Token, err = sign.SignedString([]byte("changehere"))
	if err != nil {
		errResp := buildErrResponse(errors.New("Internal Server Error !"), "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(token)

}
