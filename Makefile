IMAGE_ID := ${shell docker images 'verifymy_backend_test_golang_server' -a -q}
DB_IMAGE_ID := ${shell docker images 'verifymy_mysql_db' -a -q}
TESTS_IMAGE_ID := ${shell docker images 'verifymy_backend_test_golang_server_t' -a -q}
TESTS_DB_IMAGE_ID := ${shell docker images 'verifymy_mysql_t_db' -a -q}

down:
	docker compose down --remove-orphans

clean:
	docker compose stop server
	docker rmi ${IMAGE_ID} -f

clean-db:
	docker compose stop mysql_db
	docker rmi ${DB_IMAGE_ID} -f

clean-test-db:
	docker compose stop mysql_t_db
	docker rmi ${TESTS_DB_IMAGE_ID} -f

clean-tests:
	docker compose stop t_server
	docker rmi ${TESTS_IMAGE_ID} -f

run:
	docker compose up server -d

run-tests:
	docker compose up t_server -d

run-clean: clean-db clean down run

run-tests-clean: clean-tests clean-test-db down run-tests

