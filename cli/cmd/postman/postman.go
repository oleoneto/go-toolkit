package postman

import (
	"time"

	"github.com/google/uuid"
)

type (
	PostmanRequest struct {
		Name     string  `json:"name"`
		Request  Request `json:"request"`
		Response []any   `json:"response"`
	}

	Request struct {
		Method string          `json:"method"`
		Header []PostmanHeader `json:"header"`
		Body   PostmanBody     `json:"body"`
		URL    URL             `json:"url"`
	}

	PostmanHeader struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Type  string `json:"type"`
	}

	Raw struct {
		Language string `json:"language"`
	}

	Options struct {
		Raw Raw `json:"raw"`
	}

	PostmanBody struct {
		Mode    string  `json:"mode"`
		Raw     string  `json:"raw"`
		Options Options `json:"options"`
	}

	Variable struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	URL struct {
		Raw      string     `json:"raw"`
		Host     []string   `json:"host"`
		Path     []string   `json:"path"`
		Variable []Variable `json:"variable"`
	}

	PostmanCollection struct {
		Name  string          `json:"name"`
		Items []PostmanFolder `json:"item"`
	}

	PostmanFolder struct {
		Name  string           `json:"name"`
		Items []PostmanRequest `json:"item"`
	}
)

func (pc *PostmanCollection) ToThunderClientCollection() (collection ThunderCollection, data ThunderRequests) {
	collectionId := uuid.NewString()
	collection.ID = collectionId
	collection.Name = pc.Name
	collection.Created = time.Now()

	for i, folder := range pc.Items {
		containerId := uuid.NewString()

		collection.Folders = append(collection.Folders, ThunderFolder{
			SortNum: (i + 1) * 10101,
			ID:      containerId,
			Name:    folder.Name,
			Created: time.Now(),
		})

		data = append(data, convertPostmanFolder(folder, collectionId, containerId)...)
	}

	return
}

func (pr *PostmanRequest) ToThunderClientRequest(collectionId, containerId string, position int) ThunderRequest {
	headers := []ThunderHeader{}
	for _, h := range pr.Request.Header {
		headers = append(headers, ThunderHeader{Name: h.Key, Value: h.Value})

	}

	r := ThunderRequest{
		ID:          uuid.NewString(),
		ColID:       collectionId,
		ContainerID: containerId,
		Name:        pr.Name,
		URL:         pr.Request.URL.Raw,
		Method:      pr.Request.Method,
		SortNum:     position,
		Created:     time.Now(),
		Modified:    time.Now(),
		Headers:     headers,
	}

	if pr.Request.Body.Raw != "" && pr.Request.Body.Options.Raw.Language != "" {
		r.Body = &ThunderBody{
			Raw:  pr.Request.Body.Raw,
			Type: pr.Request.Body.Options.Raw.Language,
		}
	}

	return r
}

func convertPostmanFolder(folder PostmanFolder, collectionId, containerId string) (data []ThunderRequest) {
	for i, item := range folder.Items {
		data = append(data, item.ToThunderClientRequest(collectionId, containerId, (i+1)*1000))
	}

	return
}
