#include <windows.h>
#include "fontsize.h"

void get_console_fontsize(struct fontsize *fs)
{
    CONSOLE_FONT_INFO info;

    GetCurrentConsoleFont(GetStdHandle(STD_OUTPUT_HANDLE), FALSE, &info);
    fs->width = info.dwFontSize.X;
    fs->height = info.dwFontSize.Y;
}
