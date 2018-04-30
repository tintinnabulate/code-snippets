// Include standard IO library
#include <stdio.h>

// This could be a foreign function exposed by another programming language, for example.
// Pass in two pointers to ints, `a` and `b`,
void get_state(int *a, int *b) {
    // dereference the pointers
    *a = 42; // the address `a` now points to 42,
    *b = 1; // the address `b` now points to 1.
}

int main(void) {
    // Declare two allocations in memory for ints, `x` and `y`,
    int x;
    int y;

    // pass to `get_state` the addresses of the memory allocations of `x` and `y`
    get_state(&x, &y);

    // print the following:
    //   x: 42, y: 1
    printf("x: %d, y: %d\n", x, y);

    // exit.
    return 0;
}
