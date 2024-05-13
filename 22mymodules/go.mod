module github.com/AryanGuptaJi/mymodules

go 1.22.2

require github.com/gorilla/mux v1.8.1

// go build .
// go mod tidy
// go mod verify

// go list

// go list all

//  go list -m all
// go list -m -versions github.com/gorilla/mux

// go mod why github.com/gorilla/mux

// go mod graph

// Edit go mod file through cmd
// go mod edit -go 1.17
// go mod edit -module module_name

// go run main.go                                If we'll run this, will fetch from web itself
// go mod vendor
// go run -mod=vendor main.go                    If we'll run this, will fetch from vendor(as node modules)