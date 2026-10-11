package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func OneRepMax(weight, reps int) (float32, error) {
	// Epley
	if weight < 1 {
		return 0, errors.New("received non-positive weight value")
	} else if reps < 1 {
		return 0, errors.New("received non-positive rep count")
	}
	return float32(weight) * (1 + float32(reps)/30), nil
}

func StrengthWeightRatio(orm float32, weight int) (float32, error) {
	if weight < 1 {
		return 0, errors.New("received non-positive weight value")
	} else if orm <= 0 {
		return 0, errors.New("received non-positive ORM")
	}
	return orm / float32(weight), nil
}

func CreateSchema() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	conn, err := pgx.Connect(context.Background(), connectionString)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
}

func main() {
	CreateSchema()
}
