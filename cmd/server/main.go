package main

import (
	"fmt"
	"log"
	"net/http"
	test_service "test_server/cmd/server/services/test"
	"test_server/gen/go/test_server_interface/v1/test/testv1connect"

	"github.com/rs/cors"
)

func webServerRoutine(ipAddress string, port int, dirPath string) {
	log.Println("started go routine 'webServerRoutine'")
	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(
				http.Dir(dirPath),
			),
		),
	)
	log.Println("listening requests to the web server...")
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", ipAddress, port),
			nil,
		),
	)
}

func main() {
	go webServerRoutine(
		"127.0.0.1",
		8080,
		"/test-server/client",
	)

	mux := http.NewServeMux()

	testService := test_service.NewTestService()
	log.Println("done creating test service")

	testPath, testHandler := testv1connect.NewTestServiceHandler(testService)
	log.Println("done creating handlers of test service")

	mux.Handle(testPath, testHandler)
	log.Println("done registering handlers of test service")

	c := cors.AllowAll()
	corsHandler := c.Handler(mux)

	log.Println("listening requests to the Connect server...")
	log.Fatal(
		http.ListenAndServe(
			"127.0.0.1:8081",
			corsHandler,
		),
	)
}
