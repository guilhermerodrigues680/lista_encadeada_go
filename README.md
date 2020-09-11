
.replit

```txt
# run = "rm -f go.mod && go env -w GOPATH=$(pwd) && go env -w GOBIN=$(pwd)/bin && go env | grep 'GOPATH\\|GOBIN' && go run src/github.com/guilhermerodrigues680/lista_encadeada_go/cmd/main.go"

# run = "rm -f go.mod && go env -w GOPATH=$(pwd) && go env -w GOBIN=$(pwd)/bin && go env | grep 'GOPATH\\|GOBIN' && cd src/github.com/guilhermerodrigues680/lista_encadeada_go/cmd && go build -o bin_go_listas && ./bin_go_listas"

run = "./src/github.com/guilhermerodrigues680/lista_encadeada_go/cmd/bin_go_listas"
```