package main

import (
	"context"
	"log"

	"github.com/Krystian19/quote_management/internal/bins/internal_bff"
	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/env"
	"github.com/enescakir/emoji"
)

func Run(ctx context.Context) error {
	ev := env.GetEnv()

	db, err := db.New(db.Opts{
		Url: ev.DB_URL,
	})
	if err != nil {
		return err
	}

	bf, err := internal_bff.New(internal_bff.Opts{
		Port: ev.INTERNAL_BFF_PORT,
		DB:   db,
	})
	if err != nil {
		return err
	}

	log.Printf("%v Clerk Api running", emoji.Eagle)
	return bf.Run()
}
