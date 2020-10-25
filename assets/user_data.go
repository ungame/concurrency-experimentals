package assets

import (
	"concurrency-experimentals/models"
	"concurrency-experimentals/utils"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type UserData struct {
	Gender string
	Origin string
	Name   string
}

func (d *UserData) GenerateUser(r *rand.Rand) *models.User {
	user := new(models.User)
	user.FirstName = strings.ToLower(d.Name)
	user.LastName = strings.ToLower(d.Origin)
	user.Genger = d.Gender
	user.GenerateID()
	user.GenerateUsername(r)
	user.GenerateEmail()
	user.GenerateBirthday(r)
	user.GenerateAvatarUrl()
	user.GeneratePassword()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	return user
}

type RawUserData struct {
	Data [][]string `json:"data"`
}

func LoadUserData() []UserData {

	dirName, _ := os.Getwd()
	b := utils.LoadFromFile(dirName + "/assets/names.json")

	raw := RawUserData{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		log.Println("Unmarshal Error: ", err.Error())
		return nil
	}

	data := make([]UserData, 0, len(raw.Data))
	for _, item := range raw.Data {
		d := UserData{
			Gender: item[0],
			Origin: item[1],
			Name:   item[2],
		}
		data = append(data, d)
	}

	return data
}
