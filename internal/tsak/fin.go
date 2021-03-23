package tsak

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK.clips/internal/banner"
)

func Fin() {
	log.Debug("[ tsak.clips ] tsak.Fin() is reached")
	banner.Banner("[ Zay Gezunt ]")
}
