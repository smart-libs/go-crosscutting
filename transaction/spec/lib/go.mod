module github.com/smart-libs/go-crosscutting/transaction/spec/lib

go 1.23.0

require (
	github.com/smart-libs/go-crosscutting/event/spec/lib v0.0.1
	github.com/stretchr/testify v1.11.1
	go.uber.org/mock v0.6.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/smart-libs/go-crosscutting/event/spec/lib => ../../../event/spec/lib
