test:
	go test
upload:
	docker build --tag=dekokun/parrot .
	docker push dekokun/parrot
