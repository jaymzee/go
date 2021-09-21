#include <windows.h>
#include <stdio.h>
#include "winsz.h"

int stty_size(struct WindowSize *ws)
{
    FILE *fp;
    char buf[1024];

    /* Open the command for reading. */
    fp = popen("/usr/bin/stty size", "r");
    if (fp == NULL) {
        return 0;
    }
    /* Read the output. */
    fgets(buf, sizeof(buf), fp);
    pclose(fp);

    // check the cmd actually executed
    int len = strlen(buf);
    if (len < 4 || len > 10) {
        return 0;
    }
    sscanf(buf, "%hi %hi", &ws->rows, &ws->cols);

    // indicate success
    return 1;
}

void GetConsoleWindowSize(struct WindowSize *ws)
{
    short fontW, fontH;

    // get font size
    CONSOLE_FONT_INFO font;
    HANDLE h = GetStdHandle(STD_OUTPUT_HANDLE);
    GetCurrentConsoleFont(h, FALSE, &font);
    fontW = font.dwFontSize.X;
    fontH = font.dwFontSize.Y;

    // get rows and columns using windows API
    CONSOLE_SCREEN_BUFFER_INFO csbi;
    GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &csbi);
    ws->cols = csbi.srWindow.Right - csbi.srWindow.Left + 1;
    ws->rows = csbi.srWindow.Bottom - csbi.srWindow.Top + 1;

    // get rows and columns using stty
    // stty_size(ws);

    // if font size is sensible use that to calculate the window size
    if (fontW > 1 && fontW < 256 && fontH > 1 && fontH < 256) {
        ws->xres = ws->cols * fontW;
        ws->yres = ws->rows * fontH;
    }

    // query terminal directly
    int x=0, y=0;
    printf("\x1b[14t");
    scanf("\x1b[4;%d;%dt", &y, &x);
    scanf("\x1b[4;%d;%dt", &y, &x);
    if (x > ws->xres) {
        ws->xres = x;
    }
    if (y > ws->yres) {
        ws->yres = y;
    }
    printf("(%d, %d)\n", x, y);
}
