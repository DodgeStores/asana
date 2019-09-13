package asana

import (
	"encoding/json"
	"io/ioutil"
)

//Tag is the structure for asana
// tags assigned to tasks
type Tag struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

//Tags is an array of type Tag
type Tags []Tag

//GetTags returns a list of tags based on the client's
// accessToken
func (c *Client) GetTags() ([]Tag, error) {

	resp, err := c.Request("GET", "/tasks", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tagData := struct {
		Tags []Tag `json:"data"`
	}{}

	err = json.Unmarshal(bytesBody, &tagData)
	if err != nil {
		return nil, err
	}

	return tagData.Tags, nil
}

//GetTagByName gets a tag out of a list of tags by name
func (ts *Tags) GetTagByName(name string) *Tag {
	for _, tag := range *ts {
		if name == tag.Name {
			return &tag
		}
	}

	return nil
}
