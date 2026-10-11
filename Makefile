# Go
format:
	go fmt .

build: format
	go build .

run:
	go run.

test:
	go test .

# DB
# TODO: env file, obviously.
start_db_docker_local:
	docker run \
		-p 5432 \
		-e POSTGRES_USER=amin \
		-e POSTGRES_PASSWORD=swordfish \
		-e POSTGRES_DB=goblet \
		--name=goblet2-db \
		-it \
		postgres:18.6
