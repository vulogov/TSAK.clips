package tsak

import (
	"fmt"

	"github.com/apex/log"

	"github.com/vulogov/TSAK.clips/internal/banner"
	"github.com/vulogov/TSAK.clips/internal/conf"
)

func Version() {
	Init()
	log.Debug("[ tsak.clips ] tsak.Version() is reached")
	banner.Banner(fmt.Sprintf("[ tsak.clips %v ]", conf.BVersion))
	banner.Table()
}
