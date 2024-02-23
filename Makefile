dev:
	fresh
dev-db:
	docker compose up db -d
prod:
	docker compose up --build
tidy:
	go mod tidy