BINARY_NAME = cli
DATABASE_NAME = cli_storage
SOURCES = cmd/main/main.go

all: check_go build

check_go:
	@command -v go > /dev/null 2>&1 || { echo "Go не установлен. Установите Go для продолжения."; exit 1; }

build_dir:
	mkdir -p build

build: build_dir 
	go build -o build/$(BINARY_NAME) $(SOURCES)
	sudo cp build/$(BINARY_NAME) /usr/bin
	mkdir -p ~/.config/cli
	touch ~/.config/cli/config.yml
	cp config.yml ~/.config/cli/config.yml

run_%: build
	CONFIG="config.yml" ./build/$(BINARY_NAME) cat $*  


COMPOSE_DIR = docker


docker-cat:
	sudo docker build -t cli -f $(COMPOSE_DIR)/DockerfileCat .
	sudo docker run -it --rm cli

docker-count:
	sudo docker build -t cli -f $(COMPOSE_DIR)/DockerfileCount .
	sudo docker run -it --rm cli

docker-ls:
	sudo docker build -t cli -f $(COMPOSE_DIR)/DockerfileLs .
	sudo docker run -it --rm cli


clear:
	docker-compose -f $(COMPOSE_DIR)/docker-compose.yaml down
	docker-compose -f $(COMPOSE_DIR)/docker-compose.yaml up --build --remove-orphans

sudo-clear:
	sudo docker system prune -a

sudo-compose: sudo-clear
	sudo docker-compose -f $(COMPOSE_DIR)/docker-compose.yaml down
	sudo docker-compose -f $(COMPOSE_DIR)/docker-compose.yaml up --build --remove-orphans
