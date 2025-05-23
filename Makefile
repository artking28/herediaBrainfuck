
all:
	clear; go build -o bfe cmd-bfe/main.go; \
	go build -o bfc cmd-bfc/main.go 

clean:
	clear; rm ./bfe;