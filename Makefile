main:
	go build -o main cmd/main.go

run: main
	./main

.PHNOY : clean
clean: 
	-rm -f main
