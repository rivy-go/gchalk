go generate ./...
cp ./internal/generator/gawkgen/gawkgen.go.txt ./internal/generator/gawkgen/gawkgen.go
go run ./internal/generator/gawkgen/gawkgen.go > generated.go