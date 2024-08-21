run:
	go run cmd/main.go
table:
	migrate create -dir migrations -ext sql db
up:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/test?sslmode=disable" up
down:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/test?sslmode=disable" down
force:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/test?sslmode=disable" force