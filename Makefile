.PHONY: containers
containers:
	docker-compose up -d

.PHONY: run
run: containers
	goreman start