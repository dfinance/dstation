package config

import (
	"text/template"
)

const defaultConfigTemplate = `# This is a TOML config file to configurate connection to VM.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# VM network address to connect.
vm_address = "{{ .Address }}"

# VM data server listen address.
vm_data_listen = "{{ .DataListen }}"

# VM retry settings.

## Retry max attempts.
## Default is 0 - infinity attempts.
vm_retry_max_attempts = {{ .MaxAttempts }}

## Request timeout per attempt in ms.
## Default is 0 - infinite (no timeout).
vm_retry_req_timeout_ms = {{ .ReqTimeoutInMs }}
`

var configTemplate *template.Template

func init() {
	var err error
	tmpl := template.New("vmConfigTemplate")

	if configTemplate, err = tmpl.Parse(defaultConfigTemplate); err != nil {
		panic(err)
	}
}
