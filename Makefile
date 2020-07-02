run: build
	./setenv.sh && ./localbot

clean:
	rm -rf localbot lambdabot youtube-discord-bot.zip

build:
	go build ./cmd/lambdabot/...
	go build ./cmd/localbot/...

package: build
	zip youtube-discord-bot.zip lambdabot

deploy: package
	terraform init
	terraform apply
