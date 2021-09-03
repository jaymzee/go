#include <stdint.h>

struct WindowSize {
    int16_t rows;
    int16_t cols;
    int16_t xres;
    int16_t yres;
};

void GetConsoleWindowSize(struct WindowSize *ws);
