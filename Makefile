start-ota-manage:
	cd ota-manage
	docker compose up --build -d

build-ota-updater:
	cd client-ui
	yarn install
	yarn run generate
	cp dist/_nuxt/* ./../ota-updater/client-ui/*
	cd ../ota-updater
	go build cmd/cli/sys-upgrade.go
	go build cmd/daemon/daemon.go
	tar -zcvf ota-updater-dist.tar.gz ./client-ui ./sys-upgrade ./daemon
