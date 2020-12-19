APP=app
BACKGROUND=background
SERVING=serving

MODULE=github.com/nafiar/tkpd-recom-engine

build_app:
	@go build -o ./cmd/${APP}/${APP} ${MODULE}/cmd/${APP}

run_app:
	./cmd/${APP}/${APP}

build_background:
	@go build -o ./cmd/${BACKGROUND}/${BACKGROUND} ${MODULE}/cmd/${BACKGROUND}

run_background:
	./cmd/${BACKGROUND}/${BACKGROUND}

build_serving:
	@go build -o ./cmd/${SERVING}/${SERVING} ${MODULE}/cmd/${SERVING}

run_serving:
	./cmd/${SERVING}/${SERVING}