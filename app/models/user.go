package models

import (
	"RevelREST/app/models/mongodb"
	"time"

	bcrypt "github.com/sensu/sensu-go/backend/authentication/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

const (
	bcryptCost = 14
)

func newUserCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("users")
}

// AddUser insert a new User into database and returns
// last inserted user on success.
func AddUser(m User) (user User, err error) {
	c := newUserCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	encryptedPass, err := bcrypt.HashPassword(m.Password)
	m.Password = string(encryptedPass)
	return m, c.Session.Insert(m)
}

// UpdateUser update a User into database and returns
// last nil on success.
func (m User) UpdateUser() error {
	c := newUserCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"name": m.Name, "email": m.Email, "password": m.Password, "updatedAt": time.Now()},
	})
	return err
}

// DeleteUser Delete User from database and returns
// last nil on success.
func (m User) DeleteUser() error {
	c := newUserCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetUsers Get all User from database and returns
// list of User on success
func GetUsers() ([]User, error) {
	var (
		users []User
		err   error
	)

	c := newUserCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&users)
	return users, err
}

// GetUser Get a User from database and returns
// a User on success
func GetUser(id bson.ObjectId) (User, error) {
	var (
		user User
		err  error
	)

	c := newUserCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&user)
	return user, err
}

// GetUser by username Get a User from database and returns
// a User on success
func ValidateUser(username string) (User, error) {
	var (
		err  error
		user User
	)

	c := newUserCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"username": username}).One(&user)
	return user, err
}
