#include <iostream>
#include <cmath>

double alpha1(double x) {
    return 0;
}

double alpha2(double x) {
    return 1 + x;
}

double f(double x, double y) {
    return -x + y * (alpha2(x));
}

double q(int i, int j, int n, int m) {
    if ((i == 0 || i == n) && (j == 0 || j == m)) {
        return 0.25;
    } else if ((i == 0 || i == n) && (j != 0 && j != m)) {
        return 0.5;
    } else if ((j == 0 || j == m) && (i != 0 && i != n)) {
        return 0.5;
    }
    return 1;
}

double methodCell(double xa, double xb, double ya, double yb, int n, int m) {
    double xh((xb - xa) / n);
    double yh((ya - yb) / m);
    double s(0);
    for (int i = 0; i <= n; i++) {
        for (int j = 0; j <= m; j++) {
            s += f(xa + i * xh + xh / 2.0, ya + j * yh + yh / 2.0);
        }
    }
    return xh * yh * s;
}

double doubleIntegrate(double a, double b, double c, double d, int n, int m) {
    double hx = abs(a - b) / (double) n;
    double hy = abs(c - d) / (double) m;

    double res = 0;
    for (int i = 0; i <= n; i++) {
        for (int j = 0; j <= m; j++) {
            res += q(i, j, n, m) * f(a + i * hx, c + j * hy);
        }
    }
    return hx * hy * res;
}

int main() {
    double n = 2;
    double m = 1;
    double res = 10;
    double res1 = 10;
    double res2 = 20;
    double h = 10;
    int i = 0;
    while (abs(res - res2) / 3 >= 0.01) {
        i++;
        n *= 2;
        res = doubleIntegrate(0, 1, 0, 1, n, n);
        res2 = doubleIntegrate(0, 1, 0, 1, n * 2, n * 2);
    }
    std::cout<<i<< std::endl;
    std::cout << res2 << std::endl;
    n = 2;
    m = 1;
    res = 10;
    res1 = 10;
    res2 = 20;
    h = 10;
    i = 0;
    while (abs(res - res2) / 3 >= 0.01) {
        i++;
        n *= 2;
        res = methodCell(0, 1, 0, 1, n, n);
        res2 = methodCell(0, 1, 0, 1, n * 2, n * 2);
    }
    std::cout<<i<< std::endl;
    std::cout << res2 << std::endl;
    return 0;
}
