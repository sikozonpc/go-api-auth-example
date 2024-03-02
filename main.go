package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sikozonpc/fullstackgo/config"
	"github.com/sikozonpc/fullstackgo/handlers"
	"github.com/sikozonpc/fullstackgo/services/auth"
	"github.com/sikozonpc/fullstackgo/store"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := store.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := store.NewStore(db)

	initStorage(db)

	sessionStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		Secure:     config.Envs.CookiesAuthIsSecure,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
	})
	authService := auth.NewAuthService(sessionStore)

	router := mux.NewRouter()

	handler := handlers.New(s, authService)

	// Cars
	router.HandleFunc("/", auth.RequireAuth(handler.HandleHome, authService)).Methods("GET")
	router.HandleFunc("/cars", auth.RequireAuth(handler.HandleListCars, authService)).Methods("GET")
	router.HandleFunc("/cars", auth.RequireAuth(handler.HandleAddCar, authService)).Methods("POST")
	router.HandleFunc("/cars/{id}", auth.RequireAuth(handler.HandleDeleteCar, authService)).Methods("DELETE")
	router.HandleFunc("/cars/search", auth.RequireAuth(handler.HandleSearchCar, authService)).Methods("GET")

	// Auth
	router.HandleFunc("/auth/{provider}", handler.HandleProviderLogin).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", handler.HandleAuthCallbackFunction).Methods("GET")
	router.HandleFunc("/auth/logout/{provider}", handler.HandleLogout).Methods("GET")
	router.HandleFunc("/login", handler.HandleLogin).Methods("GET")

	// Static Files
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("Server: Listening on %s:%s\n", config.Envs.PublicHost, config.Envs.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", config.Envs.Port), router))
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
