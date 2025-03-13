prepare:
	pip install -r server/requirements.txt
	cd client && go mod tidy
	python -m grpc_tools.protoc -I. --python_out=server --grpc_python_out=server scrape.proto
	protoc --go_out=client/proto --go-grpc_out=client/proto scrape.proto

build:
	docker-compose build

up:
	docker-compose up

dbg:
	docker-compose -f docker-compose.dbg.yml up
	