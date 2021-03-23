pre:
	@echo "=== TSAK.clips === [Preinstallation of some stuff]"
	go get github.com/keysight/clipsgo
	go get github.com/client9/misspell/cmd/misspell
	go get golang.org/x/tools/cmd/godoc
	go get github.com/llorllale/go-gitlint/cmd/go-gitlint
	go get github.com/psampaz/go-mod-outdated
	go get golang.org/x/tools/cmd/goimports
	go mod download gotest.tools
	go mod download github.com/stretchr/testify
	go get gotest.tools/gotestsum
	go get github.com/stretchr/testify/assert@v1.6.1
