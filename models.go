package resource

import (
	_ "fmt"
	_ "strconv"
	"time"
)

type Source struct {
	Test string `json:"test,omitempty"`
}

type Request struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type Response struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}

type Version struct {
	Ref       string    `json:"ref"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type GetConfig struct {
	Test string `json:"test,omitempty"`
}

type PutConfig struct {
	Test string `json:"test,omitempty"`
}

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Metadata []NameValue

//func (m Metadata) Merge(input map[string]interface{}) Metadata {
//
//	var pairs []NameValue
//
//	for _, item := range m {
//		append(pairs, item)
//	}
//
//	for k, v := range input {
//		pair := NameValue{Name: k, Value: fmt.Sprintf("%v", v)}
//		append(m, pair)
//	}
//
//	return m
//}
