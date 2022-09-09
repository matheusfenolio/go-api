run:
	go run src/cmd/main.go

build:
	go build ./src/cmd

image:
	podman build -t api .

container:
	podman run -it --rm --name my-api -p 8080:8080 -e HOST=192.168.1.16 localhost/api

compose:
	podman-compose up

apply:
	terraform apply

destroy:
	terraform destroy

