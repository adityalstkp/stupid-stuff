run:
	go build -o bin/stupid cmd/stupid-cli/main.go && ./bin/stupid --scroll-dur $(scroll-dur)

run-cacher:
	go build -o bin/cacher cmd/cacher/main.go && ./bin/cacher
