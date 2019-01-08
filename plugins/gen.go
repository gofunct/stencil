//go:generate stencil assets -p plugins -s="/init" -o init.go -v Init init
//go:generate stencil assets -p plugins -s="/grpc" -o grpc.go -v Grpc grpc
//go:generate stencil assets -p plugins -s="/certs" -o certs.go -v Certs certs

package plugins
