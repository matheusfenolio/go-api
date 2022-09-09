run:
	go run src/cmd/main.go

build:
	go build ./src/cmd

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

