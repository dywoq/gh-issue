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
	"strconv"
	"strings"

	"github.com/dywoq/dywoqlib/err"
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
