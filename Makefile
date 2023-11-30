postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_PASSWORD=donghao -e POSTGRES_USER=root -d postgres:16-alpine

createDB:
	docker exec -it postgres16 createdb --username=root bank

dropDB:
	docker exec -it postgres16 dropdb bank

up:
	 migrate -path db/migration -database "postgresql://root:donghao@localhost:5432/bank?sslmode=disable" -verbose up

down:
	migrate -path db/migration -database "postgresql://root:donghao@localhost:5432/bank?sslmode=disable" -verbose down