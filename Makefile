APP=app

MODULE=github.com/nafiar/tkpd-recom-engine

build_app:
	@go build -o ./cmd/${APP}/${APP} ${MODULE}/cmd/${APP}

run_app:
	./cmd/${APP}/${APP}