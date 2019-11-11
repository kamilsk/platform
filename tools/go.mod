module tools

go 1.11

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190921135421-dca3d7403570

require (
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.21.0
	github.com/spf13/afero v1.2.2 // indirect
	golang.org/x/tools v0.0.0-20191010075000-0337d82405ff
)
