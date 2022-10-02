package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kdl-dev/elecard-test-task/pkg/client"
)

var method, params string
var key, url string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	flag.StringVar(&method, "m", "AutoExec", "AutoExec | GetTasks | CheckResults\n")
	flag.StringVar(&params, "p", "", "params (1lbx,1lby,1rtx,1rty,2lbx,2lby..)")
	flag.Parse()

	key = os.Getenv("AUTH_KEY")
	url = os.Getenv("SERVER_ADDR") + os.Getenv("SERVER_API_PATH")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := client.Resolve(client.NewCLIClient(key), url, method, params); err != nil {
			log.Printf("%v\n", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
