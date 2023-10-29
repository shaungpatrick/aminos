aminos:
	go build -o build/aminos app/main.go && \
	docker build -t amino -f build/Dockerfile . && \
	rm build/aminos