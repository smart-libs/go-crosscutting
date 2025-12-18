module githup.com/smart-libs/go-crosscutting/assertions/impl/validator/playground

go 1.24.0

require (
	github.com/go-playground/validator/v10 v10.29.0
	github.com/smart-libs/go-crosscutting/assertions/lib v0.0.0
	github.com/smart-libs/go-crosscutting/serror/lib v0.0.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
)

replace github.com/smart-libs/go-crosscutting/serror/lib => ./../../../../serror/lib

replace github.com/smart-libs/go-crosscutting/assertions/lib => ../../../../assertions/lib
