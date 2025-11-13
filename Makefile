createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres simplebank
dropdb:
	docker exec -it postgres-container dropdb --username=postgres simplebank
migrateup:
	migrate -path db/migrations -database "postgresql://postgres:passwrod123@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://postgres:passwrod123@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc: 
	sqlc generate
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
