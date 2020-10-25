package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url"`
	Genger    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GenerateID() string {
	u.ID = bson.NewObjectId().Hex()
	return u.ID
}

func (u *User) GenerateEmail() string {
	u.Email = fmt.Sprintf("%s@email.com", u.Username)
	return u.Email
}

func (u *User) GenerateUsername(r *rand.Rand) string {
	u.Username = fmt.Sprintf("%s_%d", u.FirstName, r.Int31n(99999))
	return u.Username
}

func (u *User) GenerateAvatarUrl() string {
	u.AvatarUrl = fmt.Sprintf("https://storage.com/avatars/%s-%d%02d%02d.jpg",
		u.FirstName, u.Birthday.Year(), u.Birthday.Month(), u.Birthday.Day())
	return u.AvatarUrl
}

func (u *User) GeneratePassword() string {
	b, _ := bcrypt.GenerateFromPassword([]byte(uuid.New().String()), bcrypt.DefaultCost)
	u.Password = string(b)
	return u.Password
}

func (u *User) GenerateBirthday(r *rand.Rand) time.Time {
	u.Birthday = time.Now().AddDate(-r.Intn(50), -r.Intn(12), -r.Intn(28))
	return u.Birthday
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) ToMongoUserModel() *MongoUserModel {
	return &MongoUserModel{
		ID:        bson.ObjectIdHex(u.ID),
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Genger:    u.Genger,
		AvatarUrl: u.AvatarUrl,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type MongoUserModel struct {
	ID        bson.ObjectId `bson:"id"`
	Username  string        `bson:"username"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	AvatarUrl string        `bson:"avatar_url"`
	Genger    string        `bson:"gender"`
	Birthday  time.Time     `bson:"birthday"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func (m *MongoUserModel) ToUser() *User {
	user := new(User)
	user.ID = m.ID.Hex()
	user.Username = m.Username
	user.FirstName = m.FirstName
	user.LastName = m.LastName
	user.Email = m.Email
	user.Password = m.Password
	user.Genger = m.Genger
	user.Birthday = m.Birthday
	user.CreatedAt = m.CreatedAt
	user.UpdatedAt = m.UpdatedAt
	return user
}
