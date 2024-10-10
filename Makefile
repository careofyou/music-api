include .env

stop_containers:
	@echo "Остановить другие контейнеры"
	if [ $$(docker ps -q) ]; then \
		echo "контейнеры остановлены"; \
		docker stop $$(docker ps -q); \
	else \
		echo "не найдено запущеных контейнеров"; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container:
	docker start $(DB_DOCKER_CONTAINER)

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build: 
	if [ -f "${BINARY}"]; then \
		rm ${BINARY}; \
		echo "${BINARY} deleted"; \
	fi
	@echo "building binary"
	go build -o ${BINARY} cmd/server/*.go

run: build
	./${BINARY}

stop:
	@-pkil -SIGTERM 0f "./{BINARY}"
	echo "API stopped"
