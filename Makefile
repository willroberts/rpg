test:
	# Pre-install dependencies to avoid future recompilation
	go install -v ./...
	go test -v ./... | grep -v vendor
