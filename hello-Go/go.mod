module github.com/Muturi-002/Go-Learn

go 1.20
//Only suitable for local development. Should be used together with require()
replace github.com/Muturi-002/Go-Learn/mystrings v0.0.0 => ../mystrings
//Should be used for both local development and external development.
require (
github.com/Muturi-002/Go-Learn/mystrings v0.0.0
)