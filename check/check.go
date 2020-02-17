package main

import (
	"encoding/json"
	resource "github.com/pidster/dummy-concourse-resource"
	"log"
	"os"
	"time"
)

func main() {
	var (
		request  resource.Request
		versions []resource.Version
	)

	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatal("Error processing initial request: " + err.Error())
	}

	t := time.Now()
	versions = make([]resource.Version, 1)
	versions[0] = resource.Version{
		Ref: t.Format("200601021504"),
		UpdatedAt: t,
	}

	if err := json.NewEncoder(os.Stdout).Encode(versions); err != nil {
		log.Fatalf("Error encoding JSON to stdout %s", err.Error())
	}
}
