

GOPATH=$(CURDIR)/server
GOPACKAGE=restjester
SERVERBIN=${GOPATH}/bin/restjester



all:
	GOPATH=${GOPATH} go install ${GOPACKAGE}

native:
	GOPATH=${GOPATH} go install ${GOPACKAGE}

test:
	GOPATH=${GOPATH} go test ${GOPACKAGE}

run: stop all
	${SERVERBIN}

stop:
	- killall ${GOPACKAGE}

clean:
	rm -f ${SERVERBIN}


