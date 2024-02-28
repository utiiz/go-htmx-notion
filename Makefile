build:
	go build -o ./bin/app ./cmd

run: tailwindcss templ build
	air

test:
	go test -v ./... -count=1 

tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i internal/static/css/input.css -o internal/static/css/styles.css

tailwindcss_watch:
	bun run tailwindcss --config tailwind.config.js -i internal/static/css/input.css -o internal/static/css/styles.css --watch

templ:
	templ generate -watch -proxy=http://localhost:3000

up:
	migrate -source "file://internal/database/migrations" -database "postgresql://tuteaz:$(password)@localhost:5432/notion?sslmode=disable" up

down:
	migrate -source "file://internal/database/migrations" -database "postgresql://tuteaz:$(password)@localhost:5432/notion?sslmode=disable" down

create_migration:
	migrate create -ext sql -dir internal/database/migrations -format unix $(name)

postgres_run:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=tuteaz -e POSTGRES_PASSWORD=$(password) -e POSTGRES_DB="notion" -d postgres:latest

postgres_start:
	docker container start postgres
