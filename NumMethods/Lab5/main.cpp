#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>
#include "matrix.h"

using namespace std;

double n = 10, A = 1, B = 58.414103393034324;

double p = -6.0;
double q = 8.0;
double f = 10.0;


double func(double x) {
    return (5.0 * exp(4 * x) / 4) - (3.0 * exp(2 * x) / 2) + (5.0) / 4;
}


int main() {

    double  h = 1 / n;

    vector<double> y1, y, di, ai, bi, ci;


    bi.push_back(q * h * h - 2);
    ci.push_back(1 + p * h / 2);
    di.push_back(f * h * h - A * (1 - p * h / 2));

    for (int i = 2; i < n - 1; i++) {
        ai.push_back(1 - (p * h) / 2);
        bi.push_back(q * h * h - 2);
        ci.push_back(1 + p * h / 2);
        di.push_back(f * h * h);
    }

    bi.push_back(q * h * h - 2);
    ai.push_back(1 - p * h / 2);
    ci.push_back(0);
    di.push_back(f * h * h - B * (1 + p * h / 2));

    matrix(y1, ai, bi, ci, di);

    y.push_back(A);
    y.resize(n);
    copy(y1.begin(), y1.end(), y.begin() + 1);
    y.push_back(B);


    for (int i = 0; i <= n; i++) {
        cout << y[i] << " " << func(i * h) << endl;
    }
    cout << endl;

    double diff = 0;
    for (int i = 0; i <= n; i++) {
        diff += abs(y[i] - func(i * h)) * abs(y[i] - func(i * h));
    }

    cout << "diff: " << diff / (n + 1) << endl;
    return 0;
}