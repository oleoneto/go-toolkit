package postman

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type (
	ThunderRequest struct {
		ID          string          `json:"_id"`         // 15d101c0-7067-4d03-9517-bf9ca044b59c
		ColID       string          `json:"colId"`       // 15d101c0-7067-4d03-9517-bf9ca044b59c
		ContainerID string          `json:"containerId"` // 15d101c0-7067-4d03-9517-bf9ca044b59c
		Name        string          `json:"name"`        // Create Resource
		URL         string          `json:"url"`         // {{scheme}}://{{baseURL}}/{{namespace}}/resource
		Method      string          `json:"method"`      // POST
		SortNum     int             `json:"sortNum"`     // 30000
		Created     time.Time       `json:"created"`     // 2023-02-24T19:27:49.444Z
		Modified    time.Time       `json:"modified"`    // 2023-02-24T19:27:49.444Z
		Headers     []ThunderHeader `json:"headers"`
		Params      []any           `json:"params,omitempty"`
		Body        *ThunderBody    `json:"body,omitempty"`
		Tests       []any           `json:"tests,omitempty"`
	}

	ThunderHeader struct {
		Name  string `json:"name"`  // client-id
		Value string `json:"value"` // {{client-id}}
	}

	ThunderBody struct {
		Type string `json:"type,omitempty"` // json
		Raw  string `json:"raw,omitempty"`  // "{\"name\": \"Resource Name\"}"
		Form []any  `json:"form,omitempty"` // []
	}

	ThunderFolder struct {
		ID          string    `json:"_id"`         // 15d101c0-7067-4d03-9517-bf9ca044b59c
		Name        string    `json:"name"`        // Create Resource
		ContainerID string    `json:"containerId"` // 15d101c0-7067-4d03-9517-bf9ca044b59c
		SortNum     int       `json:"sortNum"`     // 30000
		Created     time.Time `json:"created"`     // 2023-02-24T19:27:49.444Z
	}

	ThunderCollection struct {
		ID      string          `json:"_id"`     // 15d101c0-7067-4d03-9517-bf9ca044b59c
		Name    string          `json:"colName"` // Example
		SortNum int             `json:"sortNum"` // 30000
		Created time.Time       `json:"created"` // 2023-02-24T19:27:49.444Z
		Folders []ThunderFolder `json:"folders"`
	}

	ThunderEnvironment struct {
		ID       string    `json:"_id"`      // 15d101c0-7067-4d03-9517-bf9ca044b59c
		Name     string    `json:"colName"`  // Example
		SortNum  int       `json:"sortNum"`  // 30000
		Created  time.Time `json:"created"`  // 2023-02-24T19:27:49.444Z
		Modified time.Time `json:"modified"` // 2023-02-24T19:27:49.444Z
		Default  bool      `json:"default"`
		Data     []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		}
	}

	ThunderRequests []ThunderRequest
)

func (tc *ThunderCollection) WriteFile(dir string) error {
	fdata, err := json.Marshal(tc)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dir+"/thunderCollection.json", fdata, 0744)
}

func (tr *ThunderRequests) WriteFile(dir string) error {
	cdata, err := json.Marshal(tr)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dir+"/thunderclient.json", cdata, 0744)
}

func (te *ThunderEnvironment) WriteFile(dir string) error {
	cdata, err := json.Marshal(te)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dir+"/thunderEnvironment.json", cdata, 0744)
}
