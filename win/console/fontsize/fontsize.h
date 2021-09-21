#include <stdint.h>

struct fontsize {
    uint16_t width;
    uint16_t height;
};

void get_console_fontsize(struct fontsize *fs);
