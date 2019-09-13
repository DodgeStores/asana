package asana

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testTagsData = `
{"data":[{"id":90477931759759,"gid":"90477931759759","name":"Urgent & Important","resource_type":"tag"},{"id":742219771185599,"gid":"742219771185599","name":"Parkalot","resource_type":"tag"},{"id":742386996173529,"gid":"742386996173529","name":"Merch Task","resource_type":"tag"},{"id":742386996173536,"gid":"742386996173536","name":"Jeff Carter","resource_type":"tag"},{"id":742386996173537,"gid":"742386996173537","name":"James B.","resource_type":"tag"},{"id":742386996173538,"gid":"742386996173538","name":"Rick Roll","resource_type":"tag"},{"id":745800987241646,"gid":"745800987241646","name":"payroll","resource_type":"tag"},{"id":806725867143506,"gid":"806725867143506","name":"Bridgett","resource_type":"tag"},{"id":799161677367250,"gid":"799161677367250","name":"Link","resource_type":"tag"},{"id":1118672014021965,"gid":"1118672014021965","name":"Amy","resource_type":"tag"},{"id":1121785153966261,"gid":"1121785153966261","name":"@Veronda","resource_type":"tag"},{"id":1121785153966262,"gid":"1121785153966262","name":"#HR","resource_type":"tag"},{"id":1121785153966263,"gid":"1121785153966263","name":"#Roles","resource_type":"tag"},{"id":1121785153966264,"gid":"1121785153966264","name":"@RickTudor","resource_type":"tag"},{"id":1121785153966265,"gid":"1121785153966265","name":"#OPs","resource_type":"tag"},{"id":1121785153966266,"gid":"1121785153966266","name":"#PayrollVariations","resource_type":"tag"}]}
`

var (
	client *Client
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() {
	client = NewClient(nil)
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestGetTags(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testTagsData)
	})

	tags, err := client.GetTags()
	if err != nil {
		t.Fatalf("Error fetching tags: %v", err)
	}

	if len(tags) != 16 {
		t.Fatalf("Error expected tags length 15 got: %v", len(tags))
	}

	if tags[0].GID != "90477931759759" {
		t.Fatalf("Error expected gid '90477931759759' got %v", tags[0].GID)
	}

	if tags[0].Name != "Urgent & Important" {
		t.Fatalf("Error expected tag name 'Urgent & Important' got %v", tags[0].Name)
	}

	t.Logf("Tags:\n%v", tags)

}

func TestGetTag(t *testing.T) {

	tags := Tags{
		Tag{Name: "Hello", GID: "1"},
		Tag{Name: "Hi", GID: "2"},
		Tag{Name: "Howdy", GID: "3"},
	}

	//Find each test tag by name
	for _, tag := range tags {

		//Take all test tags and find the individual
		// one by name
		foundTag := tags.GetTagByName(tag.Name)

		if foundTag == nil {
			t.Errorf("tag %+v not found by name '%v' in array %+v", tag, tag.Name, tags)
		}

		if *foundTag != tag {
			t.Errorf("Expected tag %+v Got %+v", tag, *foundTag)
		}
	}

	//Make sure it doesn't find a tag
	// that's not there
	foundTag := tags.GetTagByName("Dumbo")
	if foundTag != nil {
		t.Errorf("Found tag name '%v' in %+v", "Dumbo", tags)
	}
}
