.PHONY: test
test:
	go run template.go --name test --output output
	cd ./output && make test \
	&& make service && make run && make help && make lint
