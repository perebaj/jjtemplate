.PHONY: test
test:
	go run template.go --name test --output output
	cd ./output && make dev/test \
	&& make service && make run && make help && make dev/lint && make image

.PHONY: test-compose 
test-compose:
	go run template.go --name test-compose --output output --compose
	cd ./output && make dev/test \
	&& make service && make run && make help && make dev/lint && make image
