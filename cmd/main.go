package main

import (
	"fmt"
	"net/http"
	"test/api/server"
	"test/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	
	app, err := app.New()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize application: %v", err))
	}

	srv := server.New(app)
	app.Log.WithFields(logrus.Fields{
		"Port": app.Config.Port,
	}).Info("Server Started")
	err = srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		app.Log.Error("Server Shutdown")
	}
}
