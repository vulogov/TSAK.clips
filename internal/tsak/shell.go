package tsak

import (
	"strings"

	"github.com/peterh/liner"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK.clips/internal/clips"
)

func Shell() {
	Init()
	log.Debug("[ tsak.clips ] tsak.Shell() is reached")
	line := liner.NewLiner()
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range clips.Completer {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})
	if value, err := line.Prompt("> "); err == nil {
		log.Debug("Got: ", value)
	} else if err == liner.ErrPromptAborted {
		log.Error("Aborted")
	} else {
		log.Error("Error reading line: ", err)
	}
}
