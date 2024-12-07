package app

import (
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vishalvivekm/pyqqserver/db"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/vishalvivekm/pyqqserver/handler"
	"net/http"
	"github.com/rs/cors"
	"fmt"
)

type App struct {
	Router *mux.Router
	MongoClient *mongo.Client
	Config Config
	handler *handler.Handler

}
type Config struct {
	MongoURI string
	Port string
}

func (a *App) Init() {
	
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	a.Config.MongoURI = os.Getenv("MONGO_URI")
	if a.Config.MongoURI == "" {
		log.Fatal("mongoURI env variable is required.")
	}
	if a.Config.Port == "" {
		a.Config.Port = "8080"
	}

	a.Router = mux.NewRouter()
	a.MongoClient = db.InitMongoDB(a.Config.MongoURI)
	a.handler = handler.NewHandler(a.MongoClient)
	a.InitRoutes()
}
func (a *App) GetDB() *mongo.Client {
    if a.MongoClient == nil {
        a.MongoClient = db.InitMongoDB(a.Config.MongoURI)
    }
    return a.MongoClient
}
func (a *App) InitRoutes() {
    a.Router.HandleFunc("/drive/{type}/{subject}", a.handler.GetResources).Methods("GET")
    a.Router.HandleFunc("/{course}/{semester}/{branch}", a.handler.GetSubjects).Methods("GET")
    a.Router.HandleFunc("/{course}/{semester}/{branch}/{subject}", a.handler.GetSubjectDetails).Methods("GET")
}
func (a *App) Run() {

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
        AllowCredentials: true,
    })
    handlerr := c.Handler(a.Router)
    fmt.Printf("server live on port %s\n", a.Config.Port)
    log.Fatal(http.ListenAndServe(":"+a.Config.Port, handlerr))
}