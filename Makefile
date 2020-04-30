BINARY="ingress-admin"
VERSION=1.0.2
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	#@go build -o ${BINARY} -tags=jsoniter
	echo "default"

build-frontend:
	@cd frontend/vue-admin-template/ ; npm run build:prod

build-backend:
	@\cp -rf frontend/vue-admin-template/dist/* backend/dist/
	@cd backend/ ; GOOS=linux CGO_ENABLED=0  go build -tags netgo -a -v -ldflags="-s -w" -o ${BINARY}

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

docker:
	@\cp -rf backend/dist  build/
	@\cp -rf backend/config  build/
	@\cp -rf backend/${BINARY}  build/
	@cd build ; docker build --no-cache -t ingressadmin/ingress-admin:${VERSION} . 

clean:
	@cd backend/ ; if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@cd build/ ; if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default fmt fmt-check install test vet clean
