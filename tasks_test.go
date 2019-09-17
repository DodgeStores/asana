package asana

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tasksFromTagData = `
	{"data":[{"id":505120560636486,"gid":"505120560636486","name":"Take Office Photos to Walgreens","resource_type":"task"}]}
`
var invalidJSON = `
	{"data":[{"id":505120560636486,"gid":"505120560636486","name":"""Take Office Photos to Walgreens"resource_type":"task"}]}
`

func TestComment(t *testing.T) {
	var (
		client *Client
		mux    *http.ServeMux
		server *httptest.Server
	)

	client = NewClient(nil)
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL

	//Successful Case
	mux.HandleFunc("/tasks/1/stories", func(w http.ResponseWriter, r *http.Request) {})

	err := client.CreateTaskComment(&Task{GID: "1"}, "Hi")

	server.Close()

	if err != nil {
		t.Error(err)
	}

	//No GID
	err = client.CreateTaskComment(&Task{}, "Hi")

	if err != NoGID {
		t.Errorf("Expected %v Got %v", NoGID, err)
	}

	//Send http error code
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL

	mux.HandleFunc("/tasks/1/stories", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	server.Close()

	if err == nil {
		t.Errorf("Expected error when returned non 2xx status code. Got nil")
	}

}

func TestGetTasksByTag(t *testing.T) {
	var (
		client *Client
		mux    *http.ServeMux
		server *httptest.Server
	)

	client = NewClient(nil)
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL

	//Successful Case
	mux.HandleFunc("/tags/90477931759759/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, tasksFromTagData)
	})

	tags, err := client.GetTasksByTag(Tag{
		GID:  "90477931759759",
		Name: "Urgent & Important",
	})
	if err != nil {
		t.Errorf("Expected to fetch task. Got: %+v", err)
	}

	if len(tags) != 1 {
		t.Errorf("Expected tasks length 1 Got: %d", len(tags))
	}

	testName := "Take Office Photos to Walgreens"
	if tags[0].Name != testName {
		t.Errorf("Expected %s. Got: %s", testName, tags[0].Name)
	}

	testGID := "505120560636486"
	if tags[0].GID != testGID {
		t.Errorf("Expected %s. Got: %s", testGID, tags[0].GID)
	}

	server.Close()

	//Invalid Json

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL

	mux.HandleFunc("/tags/90477931759759/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, invalidJSON)
	})

	tags, err = client.GetTasksByTag(Tag{
		GID:  "90477931759759",
		Name: "Urgent & Important",
	})
	if err == nil {
		t.Errorf("Expected invalid json, got %v", err)
	}
	t.Logf("Correct Invalid JSON Error: %+v", err)

	server.Close()

}
