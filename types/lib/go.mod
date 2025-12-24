module github.com/smart-libs/go-crosscutting/types/lib

go 1.22

require (
	github.com/smart-libs/go-crosscutting/types/impl/decimal/shopspring v0.0.1
	github.com/stretchr/testify v1.8.1
	golang.org/x/text v0.16.0
)

require (
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/leekchan/accounting v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/smart-libs/go-crosscutting/types/impl/decimal/shopspring => ../impl/decimal/shopspring
