
.PHONY: lambda
lambda:
	docker build -f Dockerfile -t fastfib:latest .

.PHONY: testlambda
testlambda:
	docker build -f Dockerfile.test -t testfastfib:latest .
