#ifndef MANDELBROT_H
#define MANDELBROT_H

#include <stdint.h>

struct rgba {
	uint8_t r, g, b, a;
} typedef rgba_t;

static const rgba_t rgba_black = { .r = 0, .g = 0, .b = 0, .a = 255 };

rgba_t mandelbrot_calculate_pixel_color(
	int32_t x,
	int32_t y,
	double shift_x,
	double shift_y,
	double zoom,
	int32_t iterations,
	int32_t width,
	int32_t height,
	const rgba_t(*color)(int32_t, int32_t)
);

rgba_t mandelbrot_color_orange(int32_t iterations, int32_t max_iterations);

static int32_t mandelbrot_calculate_number_of_iterations(
	int32_t x,
	int32_t y,
	double shift_x,
	double shift_y,
	double scale,
	int32_t iterations,
	int32_t width,
	int32_t height
);

static double mandelbrot_calculate_scale(int32_t width, int32_t height);

static double mandelbrot_normalize_x(int32_t x, int32_t width, double scale);

static double mandelbrot_normalize_y(int32_t y, int32_t height, double scale);

#endif // !MANDELBROT_H