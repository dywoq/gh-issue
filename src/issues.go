package main

import (
	"fmt"
	"net/http"
)

// issueDelete deletes a GitHub issue using the REST API.
// It sends a DELETE request to the GitHub API endpoint with a personal access token.
func issuesDelete(owner, repo string, id int, token string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", owner, repo, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("github.com/dywoq/gh-issue-deleter: unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
