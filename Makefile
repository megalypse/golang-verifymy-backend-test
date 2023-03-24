IMAGE_ID := ${shell docker images 'verifymy_backend_test_golang_server' -a -q}
clean:
	docker compose stop server
	docker rmi ${IMAGE_ID} -f

DB_IMAGE_ID := ${shell docker images 'verifymy_mysql_db' -a -q}
clean-db:
	docker compose down
	docker rmi ${DB_IMAGE_ID} -f

run-compose:
	docker compose up -d

run-compose-clean: clean
	docker compose up -d

run-compose-clean-all: clean clean-db
	docker compose down --remove-orphans
	docker compose up -d


