#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

char *GetSize(const char *path);

char *Resize(const char *path, int32_t width, int32_t height, const char *output);

void free_string(char *s);
