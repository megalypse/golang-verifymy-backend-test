# GoLang Clean Architecture

# HOW TO RUN THE APPLICATION

1. Run `make run-compose`

Check `http://localhost:3001/swagger/index.html` for Swagger Documentation.

---
## Walkthrough
Use swagger to be able easily call the endpoints and check their signatures.

# Makefile commands
* `down`: The same as `docker compose down --remove-orphans`
* `clean`: Cleans the server image. Will fail if there's no image built
* `clean-db`: Cleans the database image. Will fail if there's no image built
* `clean-test-db`: Cleans the test database image. Will fail if there's no image built
* `clean-test`: Cleans the test container image. Will fail if there's no image built
* `run`: Run the server using compose
* `run-tests`: Run the application tests
* `run-clean`: Runs the server after cleaning up its image. Will fail if there's no image built
* `run-clean-all`: Runs the server after cleaning up both its image and the database's. Will fail if there's no image built
* `run-tests-clean`: Runs the tests after cleaning up both its image and the test database's. Will fail if there's no image built
