module github.com/musianisamuele/macpass

go 1.21.3

require (
	github.com/coreos/go-iptables v0.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/jcmturner/gokrb5/v8 v8.4.4
	github.com/spf13/viper v1.17.0
	golang.org/x/term v0.13.0
	gotest.tools/v3 v3.5.1
	internal/comunication v1.0.0
)

replace internal/comunication => ./internal/comunication

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.3.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
