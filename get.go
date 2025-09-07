package main

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/google/go-github/v74/github"
)

func getLabels(i *github.Issue) []string {
	if i == nil || i.Labels == nil {
		return []string{}
	}
	labels := []string{}
	for _, label := range i.Labels {
		if label.Name != nil {
			labels = append(labels, *label.Name)
		}
	}
	return labels
}

func getAssignees(i *github.Issue) []string {
	if i == nil || i.Assignees == nil {
		return []string{}
	}
	assignees := []string{}
	for _, assignee := range i.Assignees {
		if assignee.Name != nil {
			assignees = append(assignees, *assignee.Name)
		}
	}
	return assignees
}

func getFormattedMarkdownBody(i *github.Issue) string {
	if i == nil {
		return ""
	}
	if i.Body == nil {
		return "<empty>"
	}
	body := *i.Body
	return string(markdown.Render(body, 80, 6))
}

func getTitle(i *github.Issue) string {
	if i == nil {
		return "<empty>"
	}
	return *i.Title
}

func getMilestone(i *github.Issue) string {
	if i == nil || i.Milestone == nil {
		return "<empty>"
	}
	return *i.Milestone.Title
}

func getDate(i *github.Issue) string {
	created := "<empty>"
	if i.CreatedAt != nil {
		created = i.CreatedAt.String()
	}
	updated := "<empty>"
	if i.UpdatedAt != nil {
		updated = i.UpdatedAt.String()
	}
	closed := "<empty>"
	if i.ClosedAt != nil {
		closed = i.ClosedAt.String()
	}
	return fmt.Sprintf("[Created: %v, Updated: %v, Closed: %v]", created, updated, closed)
}
