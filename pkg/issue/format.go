package issue

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dywoq/dywoqlib/lib/err"
)

func FormatToIntSlice(arg string) ([]int, err.Context) {
	if arg == "" {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/issue: arg is empty"))
		err2.SetMore("source is issue.FormatToIntSlice(string) ([]int, err.Context)")
		return []int{}, err2
	}
	ids := []int{}
	parts := strings.SplitSeq(arg, ",")
	for part := range parts {
		id, err2 := strconv.Atoi(part)
		if err2 != nil {
			return []int{}, err.NewContext(err2, "source is issue.FormatToIntSlice(string) ([]int, err.Context)")
		}
		ids = append(ids, id)
	}
	return ids, err.NoneContext()
}
