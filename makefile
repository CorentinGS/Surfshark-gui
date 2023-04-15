NAME = "surfshark-gui"

ci:
	./scripts/ci.sh

build:
	go build -ldflags="-s -w" -o bin/$(NAME)

compress:
	upx -9 bin/$(NAME)

compress-debug:
	upx -9 -d bin/$(NAME)

compress-best:
	upx -9 --best bin/$(NAME)

install:
	cp bin/$(NAME) /usr/local/bin/$(NAME)

uninstall:
	rm /usr/local/bin/$(NAME)

clean:
	rm -rf bin


all: build compress-best install

