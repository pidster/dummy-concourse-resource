package out

import (
	"encoding/json"
	resource "github.com/pidster/dummy-concourse-resource"
	"log"
	"os"
)

func main() {
	var (
		request  resource.Request
		versions []resource.Version
	)

	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatal("Error processing initial request: " + err.Error())
	}

	versions = make([]resource.Version, 0)

	if err := json.NewEncoder(os.Stdout).Encode(versions); err != nil {
		log.Fatalf("Error encoding JSON to stdout %s", err.Error())
	}
}
