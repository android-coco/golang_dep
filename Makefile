#/bin/bash
# This is how we want to name the binary output
OUTPUT=bin/dep
SRC=dep/dep.go

# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`



# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME}"

local:
	rm -f ./bin/dep
	go build ${LDFLAGS} -o ${OUTPUT} ${SRC}

debug:
	rm -f ./bin/dep
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT} ${SRC}

release:
	rm -f ./bin/dep
	CGO_ENABLED=0 GOOS=linux go build ${LDFLAGS} -o ${OUTPUT} ${SRC}
	cd .. && tar -zcvf dep_release.tar.gz --exclude build --exclude .git --exclude .gitmodules --exclude pkg --exclude src --exclude Makefile dep && mv -f dep_release.tar.gz dep/build/dep_release_${BUILD_TIME}.tar.gz && rm -f 2/bin/dep

clean:
	rm -f ./bin/dep
	rm -f ./build/dep*
