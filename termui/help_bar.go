package termui

import (
	"fmt"
	"strings"

	"github.com/MichaelMure/git-bug/util/colors"
)

type helpBar []struct {
	keys string
	text string
}

func (hb helpBar) Render() string {
	var builder strings.Builder
	for i, entry := range hb {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(fmt.Sprintf("[%s] ", entry.keys))
		builder.WriteString(colors.Bold(fmt.Sprintf("%s", entry.text)))
	}
	return builder.String()
}
