package db

import (
	"concurrency-experimentals/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type mongoDbPersistence struct {
	conn *mgo.Database
}

func NewMongoDbPersistence(conn *mgo.Database) UserPersistence {
	return &mongoDbPersistence{conn}
}

func (p *mongoDbPersistence) Create(user *models.User) error {
	return p.conn.C(UsersCollection).Insert(user.ToMongoUserModel())
}

func (p *mongoDbPersistence) Get(id string) (*models.User, error) {
	model := new(models.MongoUserModel)
	err := p.conn.C(UsersCollection).FindId(bson.ObjectIdHex(id)).One(model)
	if err != nil {
		return nil, err
	}
	return model.ToUser(), nil
}

func (p *mongoDbPersistence) GetAll() ([]*models.User, error) {
	var items []*models.MongoUserModel
	err := p.conn.C(UsersCollection).Find(bson.M{}).All(&items)
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(items))
	for _, item := range items {
		users = append(users, item.ToUser())
	}
	return users, nil
}

func (p *mongoDbPersistence) DeleteAll() error {
	return p.conn.C(UsersCollection).DropCollection()
}
