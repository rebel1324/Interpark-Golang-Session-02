build: format check build-local

build-local:
	go mod download
	go build -o "${output_path}" -v
format:
	go fmt
check:
	staticcheck .
