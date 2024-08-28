package main

/*
#include <stdio.h>
#include <stdlib.h>

void printHello(const char *name) {
	printf("Hello, %s!\n", name);
	fflush(stdout);
}

int add(int a, int b) {
	return a + b;
}

char *reserveCharMemory(int size, char *ptr) {
	ptr = (char *)malloc(size);
	if (ptr == NULL) {
		printf("Memory allocation failed\n");
		fflush(stdout);
	}
	// Print the memory address
	printf("Memory address [C]: %p\n", ptr);
	fflush(stdout);
	// Some garbage value in the memory
	printf("Value: %c\n", *ptr);
	fflush(stdout);
	return ptr;
}

void releaseCharMemory(char *ptr) {
	free(ptr);
}

*/
import "C"
import (
	"fmt"
)

func main() {
	C.printHello(C.CString("Carlos"))

	a := C.add(1, 2)
	fmt.Printf("add from c:%d\n", a)

	// Memory allocation in C
	var ptr *C.char
	ptr = C.reserveCharMemory(100, ptr)
	if ptr == nil {
		fmt.Println("Memory allocation failed")
	}
	defer C.releaseCharMemory(ptr)
	*ptr = 'C'
	fmt.Printf("Memory address [Go]: %p\n", ptr)
	fmt.Printf("Value: %c\n", *ptr)
}
