package main

import (
	"data-service/internal/entity"
	"data-service/internal/handlers"
	"data-service/internal/infrastructure/env"
	"data-service/internal/service"
	"data-service/internal/storage"
	"data-service/internal/storage/postgres"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type ConsumerMessage struct {
	RoutingKey string `json:"routingKey"`
}

type FilmMessage struct {
	Name     string `json:"name"`
	Director string `json:"director"`
}

type ReviewMessage struct {
	Rating float64 `json:"rating"`
	Text   string  `json:"text"`
}

func main() {

	// Set up configuration for the consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//Create new consumer
	consumer, err := sarama.NewConsumer([]string{"broker:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	//Subscribe to the topic
	topic := "asd"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			panic(err)
		}
	}()

	db, err := postgres.NewPostgresDB(fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=%s password=%s sslrootcert=%s",
		env.DbHost, env.DbPort, env.DbName, env.DbUser, env.DbSslMode, env.DbPassword, env.DbSslCertPath))
	if err != nil {
		return
	}
	storage := storage.NewStorage(db)
	services := service.NewService(storage)
	handlers := handlers.NewHandler(services)

	srv := new(Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while runnning http server: %s", err.Error())
		}
	}()

	// Start consuming messages
	for message := range partitionConsumer.Messages() {
		var consumerMessage ConsumerMessage
		err = json.Unmarshal(message.Value, &consumerMessage)
		if err != nil {
			fmt.Printf("Catalog consumer handler error: %s", err)
			return
		}

		if consumerMessage.RoutingKey == "film" {
			var filmMessage entity.Film
			err = json.Unmarshal(message.Value, &filmMessage)
			services.FilmService.Create(filmMessage)
		}

		if consumerMessage.RoutingKey == "review" {
			var reviewMessage entity.Review
			err = json.Unmarshal(message.Value, &reviewMessage)
			services.ReviewService.Create(reviewMessage)
		}
		fmt.Printf("Message value: %s\n", string(message.Value))
	}
}

//package main
//
//import (
//	_ "github.com/lib/pq"
//)
//
//type Film struct {
//	ID       string `faker:"uuid_hyphenated"`
//	Name     string `faker:"sentence"`
//	Director string `faker:"name"`
//}
//
//type Review struct {
//	ID     string  `faker:"uuid_hyphenated"`
//	Text   string  `faker:"sentence"`
//	Rating float32 `faker:"boundary_start=1.0, boundary_end=10.0"`
//	FilmID string
//}
//
//func main() {
//	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=%s password=%s sslrootcert=%s",
//		env.DbHost, env.DbPort, env.DbName, env.DbUser, env.DbSslMode, env.DbPassword, env.DbSslCertPath)
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Generate films data
//	var filmsData []Film
//	for i := 0; i < 50; i++ {
//		var film Film
//		err := faker.FakeData(&film)
//		if err != nil {
//			log.Fatal(err)
//		}
//		filmsData = append(filmsData, film)
//	}
//
//	filmsQuery := "INSERT INTO films (id, name, director) VALUES ($1, $2, $3)"
//	for _, film := range filmsData {
//		_, err = db.Exec(filmsQuery, film.ID, film.Name, film.Director)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	// Generate reviews data
//	var reviewsData []Review
//	for i := 0; i < 200; i++ {
//		var review Review
//		err := faker.FakeData(&review)
//		if err != nil {
//			log.Fatal(err)
//		}
//		review.FilmID = filmsData[i%50].ID
//		reviewsData = append(reviewsData, review)
//	}
//	reviewsQuery := "INSERT INTO reviews (id, text, rating, film_id) VALUES ($1, $2, $3, $4)"
//	for _, review := range reviewsData {
//		_, err = db.Exec(reviewsQuery, review.ID, review.Text, review.Rating, review.FilmID)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("Data inserted successfully!")
//}
