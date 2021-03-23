package tsak

import (
	"github.com/apex/log"

	tlog "github.com/vulogov/TSAK.clips/internal/log"
)

func Init() {
	tlog.Init()
	log.Debug("[ tsak.clips ] tsak.Init() is reached")
}
