#include "mandelbrot.h"

rgba_t mandelbrot_calculate_pixel_color(
		const int32_t x, 
		const int32_t y,
		const double shift_x, 
		const double shift_y,
		const double zoom,
		const int32_t max_iterations,
		const int32_t width,
		const int32_t height,
		const rgba_t(*color)(int32_t, int32_t)
) {
    const double scale = mandelbrot_calculate_scale(width, height) * zoom;
    int32_t iterations = mandelbrot_calculate_number_of_iterations(x, y, shift_x, shift_y, scale, max_iterations, width, height);
    return color(iterations, max_iterations);
}

rgba_t mandelbrot_color_orange(const int32_t iterations, const int32_t max_iterations) {
    if (iterations == max_iterations)
		return rgba_black;

    const int32_t v = 765 * iterations / max_iterations;
	if (v > 510) {
	    const rgba_t color = {.r = 255, .g = 255, .b = (uint8_t)(v % 255), .a = 255,};
	    return color;
	}
	else if (v > 255) {
		const rgba_t color = {.r = 255, .g = (uint8_t)(v % 255), .b = 0, .a = 255,};
		return color;
	}

	const rgba_t color = {.r = (uint8_t)(v % 255), .g = 0, .b = 0, .a = 255,};
	return color;
}

static int32_t mandelbrot_calculate_number_of_iterations(
	const int32_t x, 
	const int32_t y, 
	const double shift_x, 
	const double shift_y, 
	const double scale, 
	const int32_t iterations, 
	const int32_t width, 
    const int32_t height
) {
    const double normalized_x = mandelbrot_normalize_x(x, width, scale) + shift_x / scale;
    const double normalized_y = mandelbrot_normalize_y(y, height, scale) + shift_y / scale;

    int32_t iteration = 0;
	for (double r = 0.0, i = 0.0; r * r + i * i <= 2 * 2 && iteration < iterations; ++iteration) {
		const double r_temp = r * r - i * i + normalized_x;
		i = 2 * r * i + normalized_y;
	    r = r_temp;
    }

    return iteration;
}

static double mandelbrot_calculate_scale(const int32_t width, const int32_t height) {
    return (width + height) * 0.125;
}

static double mandelbrot_normalize_x(const int32_t x, const int32_t width, const double scale) {
    return (x - width / 2.0) / scale;
}

static double mandelbrot_normalize_y(const int32_t y, const int32_t height, const double scale) {
    return (y - height / 2.0) / scale;
}