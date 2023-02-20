#ifndef MANDELBROT_RGBA_H
#define MANDELBROT_RGBA_H

#include <stdint.h>

struct rgba {
    uint8_t r;
    uint8_t g;
    uint8_t b;
    uint8_t a;
} typedef rgba_t;

rgba_t rgba_new(uint8_t r, uint8_t g, uint8_t b, uint8_t a) {
    rgba_t c = {r, g, b, a};
    return c;
}

#endif // !MANDELBROT_RGBA_H
