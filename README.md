# Linearly Homomorphic Time-Lock Puzzles

Implementation of the Linearly-Homomorphic Time-Lock Puzzle algorithm as described in section 4.1 "Linearly Homomorphic" of the paper [Homomorphic Time-Lock Puzzles and Applications](https://eprint.iacr.org/2019/635.pdf) by Malavolta et al.

This implementation also features the extension mentioned in section 5.1 "Semi-Compact Scheme for Branching Programs" which allows for larger message spaces.

## Setup

1. `git clone <url>`
2. `asdf install` (optional)
3. `go test -count 1 -race ./...`

## Useful Commands

```sh
go run <package-path>
go build [<package-path>]

go test [<package-path>][/...] [-v] [-cover] [-race] [-short] [-parallel <number>]
go test -bench=. [<package-path>] [-count <number>] [-benchmem] [-benchtime 2s] [-memprofile <name>]

go test -coverprofile <name> [<package-path>]
go tool cover -html <name>
go tool cover -func <name>

go fmt [<package-path>]

go mod init [<module-path>]
go mod tidy
```

## Useful Resources

- [Go - Learn](https://go.dev/learn)
- [Go - Documentation](https://go.dev/doc)
- [Go - A Tour of Go](https://go.dev/tour)
- [Go - Effective Go](https://go.dev/doc/effective_go)
- [Go - Playground](https://go.dev/play)
- [Go by Example](https://gobyexample.com)
- [100 Go Mistakes and How to Avoid Them](https://100go.co)
