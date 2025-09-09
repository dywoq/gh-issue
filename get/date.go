package get

import (
	"fmt"

	"github.com/google/go-github/v74/github"
)

// dateBase returns the issue date of creating, updating, and closing.
func dateBase(i *github.Issue) (string, string, string) {
	created := "<empty>"
	if i.CreatedAt != nil {
		created = fmt.Sprintf("%d:%d:%d, %d %s", i.CreatedAt.Hour(), i.CreatedAt.Minute(), i.CreatedAt.Second(), i.CreatedAt.Day(), i.CreatedAt.Month())
	}
	updated := "<empty>"
	if i.UpdatedAt != nil {
		updated = fmt.Sprintf("%d:%d:%d, %d %s", i.UpdatedAt.Hour(), i.UpdatedAt.Minute(), i.UpdatedAt.Second(), i.UpdatedAt.Day(), i.UpdatedAt.Month())
	}
	closed := "<empty>"
	if i.ClosedAt != nil {
		closed = fmt.Sprintf("%d:%d:%d, %d %s", i.ClosedAt.Hour(), i.ClosedAt.Minute(), i.ClosedAt.Second(), i.ClosedAt.Day(), i.ClosedAt.Month())
	}
	return created, updated, closed
}
