package asana

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//Tag is the structure for asana
// tags assigned to tasks
type Tag struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

//Tags contains an array of pointers to
// Tag
type Tags []*Tag

//GetTags returns a list of tags based on the client's
// accessToken
func (c *Client) GetTags() (Tags, error) {

	resp, err := c.Request("GET", "/tags", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(bytesBody))

	tagData := struct {
		Tags Tags `json:"data"`
	}{}

	err = json.Unmarshal(bytesBody, &tagData)
	if err != nil {
		return nil, err
	}

	return tagData.Tags, nil
}

//GetTagByName gets a tag out of a list of tags by name
func (ts Tags) GetTagByName(name string) (*Tag, error) {
	for _, tag := range ts {
		if name == tag.Name {
			return tag, nil
		}
	}

	return nil, errors.New("tag not found")
}
