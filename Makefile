clean:
	rm -rf youtube-discord-bot youtube-discord-bot.zip

build:
	go build ./...

package: build
	zip youtube-discord-bot.zip youtube-discord-bot

deploy: package
	terraform init
	terraform apply
