package out

import (
	"encoding/json"
	resource "github.com/pidster/dummy-concourse-resource"
	"log"
	"os"
)

func main() {

	var config resource.PutConfig
	if err := json.NewDecoder(os.Stdin).Decode(&config); err != nil {
		log.Fatal("Error decoding payload: " + err.Error())
	}

	response := resource.Response{
		Version: resource.Version{Ref: "0"},
		Metadata: resource.Metadata{
			resource.NameValue{Name:"ConcourseUrl", Value:os.Getenv("ATC_EXTERNAL_URL")},
			resource.NameValue{Name:"BuildId", Value:os.Getenv("BUILD_ID")},
			resource.NameValue{Name:"BuildName", Value:os.Getenv("BUILD_NAME")},
			resource.NameValue{Name:"BuildJobName", Value:os.Getenv("BUILD_JOB_NAME")},
			resource.NameValue{Name:"BuildPipelineName", Value:os.Getenv("BUILD_PIPELINE_NAME")},
			resource.NameValue{Name:"BuildTeamName", Value:os.Getenv("BUILD_TEAM_NAME")},
		},
	}

	err := json.NewEncoder(os.Stdout).Encode(response); if err != nil {
		log.Fatal("Error writing response to stdout")
	}
}
