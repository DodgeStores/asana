package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Task struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

type Tasks []Task

func (c *Client) GetTasksByTag(tag Tag) (Tasks, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/tag/%s/tasks", tag.GID), nil)
	if err != nil {
		return Tasks{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Tasks{}, err
	}

	var tasks struct {
		Data Tasks `json:"data"`
	}
	err = json.Unmarshal(bodyBytes, &tasks)
	if err != nil {
		return Tasks{}, err
	}

	return tasks.Data, nil
}
