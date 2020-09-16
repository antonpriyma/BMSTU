#include <iostream>
#include <cmath>

float x = 0.6;
float y = 0.1;

double f1(double x, double y) { return sin(x + 1) - y - 1; }

double f2(double x, double y) { return 2 * x + cos(y) - 2; }

double df1dx() { return -cos(x + 1); }

double df1dy() { return -1; }

double df2dx() { return 2; }

double df2dy() { return sin(y); }

int main() {
    double x_sled, y_sled;
    double res[3];
    double J;
    int i = 0;
    float eps = 0.001;
    float x = 0.6;
    float y = 0.1;
    do {
        i++;
        //Якобиан
        J = df1dx() * df2dy() - df1dy() * df2dx();
        //вычисление очередного значения х
        x_sled = x + (((-f1(x, y)) * df2dy() - (-f2(x, y)) * df1dy()) / J);
        y_sled = y + (((-f2(x, y)) * df1dx() - (-f1(x, y)) * df2dx()) / J);
        if (abs(x_sled - x) < eps && abs(y_sled - y) < eps)
            break;
        x = x_sled;
        y = y_sled;
    } while (true);
    //double[] mas = { x_sled,y_sled };
    res[0] = x_sled;
    res[1] = y_sled;
    res[2] = i;
    std::cout << "x = " << res[0] << std::endl;
    std::cout << "y = " << res[1] << std::endl;
    std::cout << "iters = " << res[2] << std::endl;
}