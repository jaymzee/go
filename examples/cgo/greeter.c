#include <stdio.h>
#include "greeter.h"

int OtherNumber = 88;

int greet(const char *name, int year, char *out) {
    int n = sprintf(out, "Greetings, %s from %d! We come in peace :)",
                    name, year);

    return n;
}
