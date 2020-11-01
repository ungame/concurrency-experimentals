package main

import (
	"concurrency-experimentals/assets"
	"concurrency-experimentals/configs"
	"concurrency-experimentals/db"
	"concurrency-experimentals/models"
	"concurrency-experimentals/utils"
	"fmt"
	"log"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {
	//CreateUserDataConcurrent()
	//CreateProductDataConcurrent()
	CreateOrdersDataConcurrent()
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

func CreateOrdersDataConcurrent() {
	fmt.Println("Loading users...")
	users := GetUsers()
	fmt.Printf("Users loaded: %v\n", len(users))

	fmt.Println("Loading products...")
	products := GetProducts()
	fmt.Printf("Products loaded: %v\n", len(products))

	mysqlConn := db.GetMysqlConnection()
	defer mysqlConn.Close()
	persistence := db.NewMySqlDbOrdersPersistence(mysqlConn)

	err := persistence.DeleteAll()
	if err != nil {
		log.Println("Error on delete all: ", err.Error())
	}

	start := time.Now()

	wg := sync.WaitGroup{}

	total := len(users)
	div := total / 20
	tasks := 1

	for i := 0; i < total; i += div {

		final := i + div
		if final >= total {
			final = total
		}
		data := users[i:final]

		mysqlConn := db.GetMysqlConnection()
		mysqlConn.SetMaxOpenConns(1000)
		mysqlConn.SetMaxIdleConns(1000)
		mysqlConn.SetConnMaxLifetime(time.Second * 1)
		defer mysqlConn.Close()
		persistence := db.NewMySqlDbOrdersPersistence(mysqlConn)

		wg.Add(1)
		go func(list []*models.User, task int, mysqlPersistence db.OrdersPersistence) {
			defer wg.Done()
			log.Printf("Start Task #%d: Users(%v)\n", task, len(list))
			totalOrders := 0
			for _, user := range list {
				maxOrders := utils.GetRandom().Intn(10)+1
				//fmt.Printf("User %v, Orders: %d\n", user.FirstName, maxOrders)
				for x := 0; x < maxOrders; x++ {
					random := utils.GetRandom()
					product := products[random.Intn(len(products))]
					order := assets.GenerateOrder(product, user, random)
					err = mysqlPersistence.Create(order)
					if err != nil {
						log.Println("Error on create order: ", err.Error())
					} else {
						totalOrders++
					}
				}
			}
			log.Printf("Finished Task #%d: Users(%d), Orders(%d)\n", task, len(list), totalOrders)
		}(data, tasks, persistence)

		tasks++
	}
	wg.Wait()

	log.Printf("Finished: %v\n", time.Since(start))
}

func GetUsers() []*models.User {
	mongoSess, err := mgo.Dial(configs.GetMongoDbDsn())
	if err != nil {
		panic(err)
	}
	defer mongoSess.Close()

	mongoConn := mongoSess.DB("mydb")
	persistence := db.NewMongoDbPersistence(mongoConn)
	items, err := persistence.GetAll()
	if err != nil {
		panic(err)
	}
	return items
}

func GetProducts() []*models.Product {
	postgresConn := db.GetPostgresConnection()
	defer postgresConn.Close()
	persistence := db.NewPostgresProductsPersistence(postgresConn)
	items, err := persistence.GetAll()
	if err != nil {
		panic(err)
	}
	return items
}
