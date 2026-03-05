build:
	docker build -t local-present .

run:
	docker run -d -p 3999:3999 -v `pwd`:/app local-present
