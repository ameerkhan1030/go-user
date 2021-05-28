package app

import (
	"test/internal/database"
	"test/pkg/config"
	"test/pkg/logger"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type App struct {
	Config           config.Config
	Log              *logrus.Logger
	DataStore        *database.DataStore
}

func New() (App, error) {
	cfg, err := config.New()
	if err != nil {
		return App{}, errors.Wrap(err, "config new failure")
	}

	log:= logger.New(cfg)

	db, err := database.New(cfg)
	if err != nil {
		return App{}, errors.Wrap(err, "datastore new failure")
	}

	err = database.MigrateDB(db.DB)
	if err != nil {
		return App{}, errors.Wrap(err, "migration failure")
	}

	return App{cfg, log, db}, nil
}
