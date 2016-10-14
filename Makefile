

GOPATH=$(CURDIR)/server
GOPACKAGE=restjester
SERVERBIN=${GOPATH}/bin/restjester



all:
	GOPATH=${GOPATH} go install ${GOPACKAGE}


test:
	GOPATH=${GOPATH} go test ${GOPACKAGE}


run: stop all
	${SERVERBIN}


stop:
	- killall ${GOPACKAGE}


clean:
	rm -f ${SERVERBIN}


xcompile:

	# windows 64
	mkdir -p releases/windows/amd64
	GOPATH=${GOPATH} GOOS=windows GOARCH=amd64 go build -o releases/windows/amd64/restjester ${GOPACKAGE}  

	# OSX 64
	mkdir -p releases/OSX/amd64
	GOPATH=${GOPATH} GOOS=darwin GOARCH=amd64 go build -o releases/darwin/amd64/restjester ${GOPACKAGE}  

	# linux 64
	mkdir -p releases/linux/amd64
	GOPATH=${GOPATH} GOOS=linux GOARCH=amd64 go build -o releases/linux/amd64/restjester ${GOPACKAGE}  

	# linux arm ( chromeboook )
	mkdir -p releases/linux/arm
	GOPATH=${GOPATH} GOOS=linux GOARCH=arm go build -o releases/linux/arm/restjester ${GOPACKAGE}  
	
	# linux arm64 ( better chromeboook )
	mkdir -p releases/linux/arm64
	GOPATH=${GOPATH} GOOS=linux GOARCH=arm64 go build -o releases/linux/arm64/restjester ${GOPACKAGE}  
