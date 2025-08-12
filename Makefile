REPO=go-cognito

build:
	docker build -t $(REPO) .

run:
	docker run \
		-e AWS_ACCESS_KEY_ID \
		-e AWS_SECRET_ACCESS_KEY \
		-e AWS_DEFAULT_REGION=us-east-2 \
		-p 8080:8080 \
		$(REPO)