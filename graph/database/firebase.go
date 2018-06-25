package database

import (
	"log"

	Firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
	"os"
	"google.golang.org/api/option"
)

var DB *db.Client
var Context context.Context

func InitializeFirebase() {
	// Export database Context
	Context = context.Background()

	// set configuration and options
	conf := &Firebase.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	opt := option.WithCredentialsFile("graph/database/serviceAccountKey.json")

	// Create firebase app and db client variable and export
	app, err := Firebase.NewApp(Context, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(Context)
	if err != nil {
		log.Fatalf("Error initializing database client:", err)
	}

	DB = client
}
