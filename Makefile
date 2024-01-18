.PHONY: test
test:
	mkdir -p output 
	go run template.go --name test --output output
	cd ./output && make dev/test \
	&& make service && make run && make help && make dev/lint && make image

.PHONY: test-compose 
test-compose:
	mkdir -p output 
	go run template.go --name test-compose --output output --compose
	cd ./output && make dev/test \
	&& make service && make run && make help && make dev/lint && make image
