package main

import (
	"github.com/unclerayray/cloudrun/bq"
	"github.com/unclerayray/cloudrun/tasks"
	"log"
	"os"
)

type EnvConfig struct {
	// Job-defined
	greeting string
	uid      string
}

func main() {
	config, err := getConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting Task")

	log.Println(config.greeting)

	log.Printf("Query and save employee: %s to BQ", config.uid)

	tasks.QueryAndSave(config.uid)

	log.Printf("Completed Task")

	bq.QueryBQ()

	log.Println("BQ query completed")
}

func getConfigFromEnv() (EnvConfig, error) {
	// [START cloudrun_jobs_env_vars]
	// Job-defined
	greeting := os.Getenv("GREETING")
	uid := os.Getenv("UID")

	config := EnvConfig{
		greeting: greeting,
		uid:      uid,
	}
	return config, nil
}
