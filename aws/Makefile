.PHONY: deps clean build deploy

deps:
	go get -u ./...

clean: 
	rm -rf ./somefortune/somefortune
	rm fortune-stack.yaml
	
build:
	GOOS=linux GOARCH=amd64 go build -o somefortune/somefortune ./somefortune

deploy:
	sam package --template-file template.yaml --output-template-file fortune-stack.yaml --s3-bucket fortune-app
	sam deploy --template-file fortune-stack.yaml --stack-name fortunestack --capabilities CAPABILITY_IAM