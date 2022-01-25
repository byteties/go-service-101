package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
  "github.com/gin-gonic/gin"
)

func getMovies(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	coll := client.Database("go-service-101").Collection("movies")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	var datasMongo []bson.M
	if err = cursor.All(context.TODO(), &datasMongo); err != nil {
			log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, datasMongo)
}

func postMovies(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	coll := client.Database("go-service-101").Collection("movies")

	doc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
	c.IndentedJSON(http.StatusCreated, result)
}

func main() {
	router := gin.Default()
  router.GET("/movies", getMovies)
  router.POST("/movies", postMovies)
  router.Run("localhost:8080")
}