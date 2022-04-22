INSTALL_DIR = ${HOME}/.local/bin/

install: lsemoji toemoji
	mv ./lsemoji ${INSTALL_DIR}
	mv ./toemoji ${INSTALL_DIR}

lsemoji:
	go build -o ./lsemoji cmd/lsemoji/lsemoji.go

toemoji:
	go build -o ./toemoji cmd/toemoji/toemoji.go

test:
	go test ./...

