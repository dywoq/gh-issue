// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issue

import (
	"errors"

	"github.com/dywoq/dywoqlib/err"
	"github.com/google/go-github/v74/github"
)

// Close closes the Github issue with given id, repo, token, its owner and token.
// If owner, repo or token is empty, it returns error.
func Close(owner, repo, token string, id int) err.Context {
	if owner == "" || repo == "" || token == "" {
		return err.NewContext(errors.New("github.com/dywoq/gh-issue/pkg/issue: owner, repo or token is empty"), "source is issue.Close(string, string, string, int) err.Context")
	}
	c, ctx := client(token)
	req := &github.IssueRequest{State: github.Ptr("closed")}
	_, _, err2 := c.Issues.Edit(ctx, owner, repo, id, req)
	if err2 != nil {
		return err.NewContext(err2, "source is issue.Close(string, string, string, int) err.Context")
	}
	return err.NoneContext()
}
