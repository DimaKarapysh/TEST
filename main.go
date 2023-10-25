package main

import (
	"TEST2/agify"
	"TEST2/app"
	"TEST2/delivery"
	"TEST2/iteractors"
	"TEST2/repository"
	"TEST2/tools/config"
	"context"
	trmcontext "github.com/avito-tech/go-transaction-manager/trm/context"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	trmsql "github.com/avito-tech/go-transaction-manager/sql"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)

	}
}

func Run() error {
	// Env
	err := config.InitEnv()
	if err != nil {
		return errors.Wrap(err, "Config")
	}

	// Log
	logs, err := app.InitLogs()
	if err != nil {
		return errors.Wrap(err, "Log")
	}

	// Db
	db, err := app.InitDatabase(logs)
	if err != nil {
		return errors.Wrap(err, "DB")
	}

	// Migration
	err = app.RunMigrations()
	if err != nil {
		return errors.Wrap(err, "Migration")
	}

	//context
	ctx := context.Background()

	// Validator
	v := config.InitValidation()

	//transaction
	tr := trmsql.DefaultCtxGetter

	trManager := manager.Must(
		trmsql.NewDefaultFactory(db),
		manager.WithCtxManager(trmcontext.DefaultManager),
	)

	repos := repository.NewPersonRepos(logs, db, ctx, tr)

	agify := agify.NewAgify(logs)

	iter := iteractors.NewIterPerson(logs, agify, repos, ctx, trManager)

	delivery := delivery.NewPersonDeliveryService(logs, iter, v)

	// Router
	router := http.NewServeMux()
	router.HandleFunc("/get", delivery.GetPersons)
	router.HandleFunc("/getId/", delivery.GetPersonsById)
	router.HandleFunc("/getName/", delivery.GetPersonsByName)
	router.HandleFunc("/getLimit/", delivery.GetPersonsByLimit)
	router.HandleFunc("/getOffset/", delivery.GetPersonsByOffset)
	router.HandleFunc("/getLimitOffset/", delivery.GetPersonsByLimitOffset)
	router.HandleFunc("/insert", delivery.AddPerson)
	router.HandleFunc("/update", delivery.UpdatePerson)
	router.HandleFunc("/delete", delivery.DeletePerson)
	srv := &http.Server{
		Addr:    os.Getenv("REST_PORT"),
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", os.Getenv("REST_PORT"))

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Program ended!")

	return nil
}
