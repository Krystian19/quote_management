package main

import (
	"context"
	"log"

	"github.com/Krystian19/quote_management/internal/bins/external_bff"
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

	bf := external_bff.New(external_bff.Opts{
		Port: ev.EXTERNAL_BFF_PORT,
		DB:   db,
	})

	log.Printf("%v External BFF Api running", emoji.TRex)
	return bf.Run()
}
