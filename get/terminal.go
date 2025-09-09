package get

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"

	"github.com/crazy3lf/colorconv"
	"github.com/fatih/color"
	"github.com/google/go-github/v74/github"
)

type Terminal struct {
	Issue *github.Issue
}

func (t Terminal) Title() string {
	if t.Issue == nil {
		return color.Set(color.Italic).Sprint("None")
	}
	return color.Set(color.Italic).Sprint(*t.Issue.Title)
}

func (t Terminal) Labels() []string {
	if t.Issue == nil || t.Issue.Labels == nil {
		return []string{}
	}
	labels := []string{}
	for _, label := range t.Issue.Labels {
		r, g, b, err := colorconv.HexToRGB(*label.Color)
		if err != nil {
			return labels
		}
		labels = append(labels, color.RGB(int(r), int(g), int(b)).Sprint(*label.Name))
	}
	return labels
}

func (t Terminal) Milestone() string {
	if t.Issue == nil || t.Issue.Milestone == nil {
		return color.Set(color.Italic).Sprint("None")
	}
	return color.Set(color.Italic).Sprint(*t.Issue.Milestone.Title)
}

func (t Terminal) Date() string {
	if t.Issue == nil {
		return "[]"
	}
	created, updated, closed := dateBase(t.Issue)
	return fmt.Sprintf("[Created: %s, Updated: %s, Closed: %s]", created, updated, closed)
}

func (t Terminal) Body() string {
	if t.Issue == nil || t.Issue.Body == nil {
		return color.Set(color.Italic).Sprint("Empty")
	}
	return string(markdown.Render(*t.Issue.Body, 80, 6))
}

func (t Terminal) State() string {
	if t.Issue == nil {
		return ""
	}
	state := color.Set(color.Italic).Sprint("None")
	switch *t.Issue.State {
	case "open":
		state = color.Set(color.Italic).Sprint("Open ❇️")
	case "closed":
		state = color.Set(color.Italic).Sprint("Closed ✔️")
	}
	return state
}

func (t Terminal) Assignees() []string {
	if t.Issue == nil || t.Issue.Assignees == nil {
		return []string{}
	}
	assignees := []string{}
	for _, assignee := range t.Issue.Assignees {
		if assignee.Name != nil {
			assignees = append(assignees, *assignee.Name)
		}
	}
	return assignees
}
