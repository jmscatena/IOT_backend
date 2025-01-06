package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"backend/adapter/postgre"
	"backend/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgre.GetConnection(ctx)
	defer conn.Close()

	postgre.RunMigrations()
	sensorService := di.ConfigSensorDI(conn)

	router := mux.NewRouter()
	router.Handle("/sensor", http.HandlerFunc(sensorService.Create)).Methods("POST")
	router.Handle("/sensor", http.HandlerFunc(sensorService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
