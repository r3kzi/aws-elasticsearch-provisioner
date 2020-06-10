image:
	docker build --no-cache -t elasticsearch-provisioner -t rekzi/elasticsearch-provisioner:latest .
test:
	go test -cover -mod=vendor ./...