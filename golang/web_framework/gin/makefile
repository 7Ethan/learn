build:
	GOOS=linux GOARCH=amd64 go build 
	docker build -t gin_demo .
run:
	docker run -t -p 8080:8080 gin_demo