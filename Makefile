test:
	# Pre-install dependencies to avoid future recompilation
	go install -v ./...
	go test -v ./... | grep -v vendor
	golint
	golint rpg/...
	go vet github.com/willroberts/rpg
	go vet github.com/willroberts/rpg/rpg
