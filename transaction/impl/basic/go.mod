module github.com/smart-libs/go-crosscutting/transaction/impl/basic

go 1.23.0

require (
	github.com/joomcode/errorx v1.2.0
	github.com/smart-libs/go-crosscutting/event/spec/lib v0.0.1
	github.com/smart-libs/go-crosscutting/transaction/spec/lib v0.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	go.uber.org/mock v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/smart-libs/go-crosscutting/event/spec/lib => ../../../event/spec/lib
	github.com/smart-libs/go-crosscutting/transaction/spec/lib => ../../spec/lib
)
