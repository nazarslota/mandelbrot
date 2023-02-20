#ifndef MANDELBROT_H
#define MANDELBROT_H

#include "rgba.h"

int mandelbrot_iterations(int w, int h, double x, double y, double x_shift, double y_shift, double zoom, int max_iterations) {
    double real = (x / ((double) h / (double) w)) / zoom + x_shift;
    double imag = y / zoom + y_shift;

    double z_real = 0.0;
    double z_imag = 0.0;

    int iterations = 0;

    while (z_real * z_real + z_imag * z_imag <= 4.0 && iterations < max_iterations) {
        double new_real = z_real * z_real - z_imag * z_imag + real;
        double new_imag = 2.0 * z_real * z_imag + imag;

        z_real = new_real;
        z_imag = new_imag;

        iterations++;
    }

    return iterations;
}

rgba_t mandelbrot_color(int iterations, int max_iterations) {
    if (iterations == max_iterations) {
        return rgba_new(0, 0, 0, 255);
    }

    const int v = 765 * iterations / max_iterations;
    if (v > 510) {
        return rgba_new(255, 255, (uint8_t) (v % 255), 255);
    } else if (v > 255) {
        return rgba_new(255, (uint8_t) (v % 255), 0, 255);
    }
    return rgba_new((uint8_t) (v % 255), 0, 0, 255);
}

#endif // !MANDELBROT_H
