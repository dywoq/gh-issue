package get

import (
	"fmt"

	"github.com/google/go-github/v74/github"
)

type Markdown struct {
	Issue *github.Issue
}

func (m Markdown) Title() string {
	if m.Issue == nil {
		return "_None_"
	}
	return *m.Issue.Title
}

func (m Markdown) State() string {
	if m.Issue == nil {
		return ""
	}
	color := "#57b76aff"
	name := "Open ❇️"
	switch *m.Issue.State {
	case "closed":
		color = "#a77aff"
		name = "Closed ✔️"
	}
	return fmt.Sprintf("<span style=\"color:%s;\">%s</span>", color, name)
}

func (m Markdown) Labels() []string {
	if m.Issue == nil || m.Issue.Labels == nil {
		return []string{}
	}
	labels := []string{}
	for _, label := range m.Issue.Labels {
		if label != nil {
			labels = append(labels, fmt.Sprintf("<span style=\"color:#%s;\">%s</span>", *label.Color, *label.Name))
		}
	}
	return labels
}

func (m Markdown) Milestone() string {
	if m.Issue == nil || m.Issue.Milestone == nil {
		return "_None_"
	}
	return fmt.Sprintf("_%s_", *m.Issue.Milestone.Title)
}

func (m Markdown) Assignees() []string {
	if m.Issue == nil || m.Issue.Assignees == nil {
		return []string{}
	}
	assignees := []string{}
	for _, assignee := range m.Issue.Assignees {
		if assignee.Login != nil {
			assignees = append(assignees, *assignee.Login)
		}
	}
	return assignees
}

func (m Markdown) Date() string {
	if m.Issue == nil {
		return ""
	}
	created, updated, closed := dateBase(m.Issue)
	return fmt.Sprintf("[ Created: %s, Updated: %s, Closed: %s ]", created, updated, closed)
}

func (m Markdown) Link() string {
	if m.Issue == nil {
		return "_None_"
	}
	return fmt.Sprintf("https://github.com/%s/%s/issues/%d", *m.Issue.Repository.Owner, *m.Issue.Repository.Name, *m.Issue.Number)
}

func (m Markdown) Body() string {
	if m.Issue == nil || m.Issue.Body == nil {
		return "_Empty_"
	}
	return *m.Issue.Body
}
