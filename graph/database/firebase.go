package database

import (
	"log"

	Firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var DB *db.Client
var Context context.Context

func InitializeFirebase() {
	Context = context.Background()

	conf := &Firebase.Config{
		DatabaseURL: "https://dreamcatcher-34304.firebaseio.com",
	}
	opt := option.WithCredentialsFile("")

	app, err := Firebase.NewApp(Context, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(Context)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	DB = client
}

func main1() {
	// var result map[string]Job
	// if err := ref.Get(ctx, &result); err != nil {
	// 	log.Fatalln("Error reading from database:", err)
	// }
	// log.Println(121212, result)
	// for key, acc := range result {
	// 	log.Printf("%s => %v\n", key, acc.Name)
	// }
}
