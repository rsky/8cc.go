CFLAGS=-Wall -std=gnu99 -g
OBJS=main.o lex.o string.o

8cc: $(OBJS) 8gg
	$(CC) $(CFLAGS) -o $@ $(OBJS)
	./build-hook.sh

clean:
	rm -f 8cc *.o tmp.*
	rm -f 8gg

8gg: 8cc.go adapter.go
	GOOS=linux  GOARCH=amd64 go build -o 8gg.linux 8cc.go adapter.go
	GOOS=darwin GOARCH=amd64 go build -o 8gg.mac   8cc.go adapter.go

