
all:
	clear && cd cmd-bfe && go build -o ../bfe && cd .. && go build -o bfc cmd-bfc/main.go;

clean:
	clear; rm ./bfe ./bfc;

