package asana

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testTagsData = `{"data":[{"id":90477931759759,"gid":"90477931759759","name":"Urgent & Important","resource_type":"tag"}]}`

func TestGetTags(t *testing.T) {

	var (
		client *Client
		mux    *http.ServeMux
		server *httptest.Server
	)

	client = NewClient(nil)
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client.BaseURL = server.URL

	defer server.Close()

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testTagsData)
	})

	tags, err := client.GetTags()
	if err != nil {
		t.Fatalf("Error fetching tags: %v", err)
	}

	if len(tags) != 1 {
		t.Fatalf("Error expected tags length 1 got: %v", len(tags))
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
		foundTag, err := tags.GetTagByName(tag.Name)
		if err != nil {
			t.Errorf("Error finding tag by name '%v': %+v", tag.Name, err)
		}

		if foundTag != tag {
			t.Errorf("Expected tag %+v Got %+v", tag, foundTag)
		}
	}

	//Make sure it doesn't find a tag
	// that's not there
	_, err := tags.GetTagByName("Dumbo")
	if err == nil {
		t.Errorf("Found tag name '%v' in %+v. Expected error", "Dumbo", tags)
	}
}
