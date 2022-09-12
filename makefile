run:
	go run src/main.go

build:
	go build ./src/main.go

test:
	go test ./src/internal/customer -coverprofile=coverage.out
	go tool cover -html=coverage.out

coverage:
	go test ./src/internal/customer -coverprofile=src/coverage.out

image:
	docker build -t api .

container:
	docker run -it --rm --name my-api -p 8080:8080 -e HOST=192.168.1.16 localhost/api

compose:
	docker-compose up

apply:
	terraform apply

destroy:
	terraform destroy

