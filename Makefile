APP=app
BACKGROUND=background

MODULE=github.com/nafiar/tkpd-recom-engine

build_app:
	@go build -o ./cmd/${APP}/${APP} ${MODULE}/cmd/${APP}

run_app:
	./cmd/${APP}/${APP}

build_background:
	@go build -o ./cmd/${BACKGROUND}/${BACKGROUND} ${MODULE}/cmd/${BACKGROUND}

build_background:
	./cmd/${BACKGROUND}/${BACKGROUND}