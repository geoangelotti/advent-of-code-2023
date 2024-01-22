CC = gcc

CFLAGS = -Wall -Werror
LIBS = 

all:

lib.o: lib.c
	$(CC) $(CFLAGS) -c -o lib.o lib.c

test: lib.o test.c
	$(CC) $(CFLAGS) -o test.out test.c lib.o

clean:
	rm -rf *.o *.out