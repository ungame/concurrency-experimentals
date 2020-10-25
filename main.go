package main

import (
	"concurrency-experimentals/assets"
	"concurrency-experimentals/configs"
	"concurrency-experimentals/db"
	"concurrency-experimentals/utils"
	"log"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {
	// CreateUserDataConcurrent()
	CreateProductDataConcurrent()
}

func CreateUserDataConcurrent() {

	mongoSess, err := mgo.Dial(configs.GetMongoDbDsn())
	if err != nil {
		panic(err)
	}
	defer mongoSess.Close()

	mongoConn := mongoSess.DB("mydb")
	persistence := db.NewMongoDbPersistence(mongoConn)

	err = persistence.DeleteAll()
	if err != nil {
		log.Println("Error on delete all: ", err.Error())
	}

	start := time.Now()

	items := assets.LoadUserData()
	total := len(items)
	log.Printf("Create Data: %d Items\n", total)

	wg := sync.WaitGroup{}

	div := total / 50
	tasks := 1

	for i := 0; i < total; i += div {

		final := i + div
		if final >= total {
			final = total
		}
		data := items[i:final]

		wg.Add(1)
		go func(d []assets.UserData, task int) {
			defer wg.Done()
			log.Printf("Start Task #%d: Items(%v)\n", task, len(d))
			for _, item := range d {
				random := utils.GetRandom()
				user := item.GenerateUser(random)
				err := persistence.Create(user)
				if err != nil {
					log.Println("Error on create user: ", err.Error())
				}
			}
			log.Printf("Finished Task #%d: Items(%d)\n", task, len(d))
		}(data, tasks)

		tasks++
	}
	wg.Wait()

	log.Printf("Finished: Items(%d), Time(%v)\n", len(items), time.Since(start))
}

func CreateProductDataConcurrent() {

	postgresConn := db.GetPostgresConnection()
	defer postgresConn.Close()
	persistence := db.NewPostgresProductsPersistence(postgresConn)

	err := persistence.DeleteAll()
	if err != nil {
		log.Println("Error on delete all: ", err.Error())
	}

	start := time.Now()

	items := assets.LoadProductData()
	total := len(items)
	log.Printf("Create Data: %d Items\n", total)

	wg := sync.WaitGroup{}

	div := total / 20
	tasks := 1

	for i := 0; i < total; i += div {

		final := i + div
		if final >= total {
			final = total
		}
		data := items[i:final]

		postgresConn := db.GetPostgresConnection()
		defer postgresConn.Close()
		persistence := db.NewPostgresProductsPersistence(postgresConn)

		wg.Add(1)
		go func(d []assets.ProductData, task int) {
			defer wg.Done()
			log.Printf("Start Task #%d: Items(%v)\n", task, len(d))
			for _, item := range d {
				random := utils.GetRandom()
				product := item.GenerateProduct(random)
				err := persistence.Create(product)
				if err != nil {
					log.Println("Error on create product: ", err.Error())
				}
			}
			log.Printf("Finished Task #%d: Items(%d)\n", task, len(d))
		}(data, tasks)

		tasks++
	}
	wg.Wait()

	log.Printf("Finished: Items(%d), Time(%v)\n", len(items), time.Since(start))
}
