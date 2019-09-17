package asana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//Task holds the GID and name of a given
// task
type Task struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

//CreateTaskComment takes a string and comments on a given task
func (c *Client) CreateTaskComment(task *Task, msg string) error {
	if task.GID == "" {
		return NoGID
	}

	jsonMsg := struct {
		Data map[string]interface{} `json:"data"`
	}{
		map[string]interface{}{
			"text": msg,
		},
	}

	jsonBytes, err := json.Marshal(jsonMsg)
	if err != nil {
		return err
	}

	resp, err := c.Request("POST", fmt.Sprintf("/tasks/%s/stories", task.GID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(bodyBytes))
	}

	return nil

}

//GetTasksByTag fetches a list of tasks using a
// given Tag's GID if available
func (c *Client) GetTasksByTag(tag Tag) ([]Task, error) {
	if tag.GID == "" {
		return nil, NoGID
	}

	resp, err := c.Request("GET", fmt.Sprintf("/tags/%s/tasks", tag.GID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tasks struct {
		Data []Task `json:"data"`
	}
	err = json.Unmarshal(bodyBytes, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks.Data, nil
}
