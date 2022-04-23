INSTALLDIR = ${HOME}/.local/bin
CONFIGDIR = ${HOME}/.config/go-emoji

.PHONY: test config install

install: lsemoji toemoji
	mv ./lsemoji ${INSTALLDIR}/
	mv ./toemoji ${INSTALLDIR}/

lsemoji:
	go build -o ./lsemoji cmd/lsemoji/lsemoji.go

toemoji:
	go build -o ./toemoji cmd/toemoji/toemoji.go

test:
	go test ./...

config: ${CONFIGDIR}
	cp ./opt/config.yaml ${CONFIGDIR}

${CONFIGDIR}:
	mkdir -p ${CONFIGDIR}
