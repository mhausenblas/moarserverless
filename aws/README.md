# fortune Lambda


In order to build the app, clone this repo and work in the `aws/` subdirectory. Also, make sure you've got the `aws` CLI and the [SAM CLI](https://github.com/awslabs/aws-sam-cli) installed.

## Preparation

Set up the S3 bucket holding the Lambda function:

```bash
$ aws s3api create-bucket \
      --bucket mh9-fortune-app \
      --create-bucket-configuration LocationConstraint=eu-west-1 \
      --region eu-west-1
```

## Lambda functions and HTTP API

How to build and deploy the Lambda functions and a HTTP API with [SAM](https://github.com/awslabs/serverless-application-model). 

### Local development

To get started I did `sam init --name app --runtime go1.x` initially and developed the function. In order for the local simulation to work, you need to have Docker running. Note: Local testing the API is at time of writing limited, since [CORS is locally not supported](https://github.com/awslabs/aws-sam-cli/issues/323), yet.

For each code iteration, in `app/` do:

```bash
# 1. run emulation of Lambda and API Gateway locally (via Docker):
$ sam local start-api

# 2. update Go source code (add functionality, fix bugs)

# 3. create new binaries which are also automagically synced into the local SAM runtime:
$ make build
```

If you change anything in the SAM/CF [template file](template.yaml) then you need to re-start the local API emulation.

### Development and deployment in live environment

For each code iteration, in `app/` do a `make deploy`.

```bash
# get the HTTP API endpoint:
$ aws cloudformation describe-stacks \
      --stack-name fortunestack | \
      jq '.Stacks[].Outputs[] | select(.OutputKey=="FortuneAPIEndpoint").OutputValue' -r

# invoke function via HTTP API endpoint from previous step:
curl  https://XXXXXXXXXX.execute-api.eu-west-1.amazonaws.com/Prod/some/
```

### Clean up

```bash
$ aws cloudformation delete-stack \
      --stack-name fortunestack
```
