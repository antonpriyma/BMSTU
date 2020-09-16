//#include <iostream>
//#include <vector>
//#include <cmath>
//
//using namespace std;
//
//float a = 0, b = 10;
//int p = -6;
//int q = 8;
//int n = 10;
//
//typedef vector<float> digits;
//digits c;
//
//float func(){
//    return 10;
//}
//
//void solveMatrix(int n, float *a, float *b, float *cc, float *d, float *x) {
//    float y[n], alfa[n], beta[n];
//    y[0] = b[0];
//    alfa[0] = -cc[0] / y[0];
//    beta[0] = d[0] / y[0];
//
//    for (int i = 1; i < n - 1; i++) {
//        y[i] = b[i] + a[i - 1] * alfa[i - 1];
//        alfa[i] = -cc[i] / y[i];
//        beta[i] = (d[i] - a[i - 1] * beta[i - 1]) / y[i];
//    }
//
//    y[n - 1] = b[n - 1] + a[n - 2] * alfa[n - 2];
//    beta[n - 1] = (d[n - 1] - a[n - 2] * beta[n - 2]) / y[n - 1];
//
//    x[n - 1] = beta[n - 1];
//    for (int i = n - 2; i >= 0; i--) {
//        x[i] = alfa[i] * x[i + 1] + beta[i];
//    }
//
//    for (int i = 1; i < n - 1; i++) {
//        c.push_back(x[i]);
//        cout << " y: " << c[i] << endl;
//    }
//    cout << endl;
//    c.push_back(52.25);
//}
//
//void method() {
//    double h = 1.0 / n;
//    c.clear();
//
//    double x[n];
//    float f[n];
//    c.push_back(a);
//
//
//    for (int i = 0; i < n; i++) {
//        x[i] = i * h;
//        cout << "x: " << x[i] << endl;
//    }
//    cout << endl;
//    for (int i = 0; i < n; i++) {
//        f[i] = func() + 1;
//        cout << "f: " << f[i] << endl;
//    }
//    cout << endl;
//
//    digits cc(n - 1);
//
//
//    for (int i = 0; i < n - 1; i++) {
//        cc[i] = (1 + p * h / 2);
//    }
//
//    digits bb(n);
//    for (i = 0; i < n; i++) {
//        bb[i] = (h * h * q - 2);
//    }
//
//    digits aa(n - 1);
//    for (i = 0; i < n - 1; i++) {
//        aa[i] = (1 - p * h / 2);
//    }
//
//    digits xx(n);
//
//    for (i = 0; i < n; i++) {
//        xx[i] = 0;
//    }
//
//    digits dd(n);
//
//    for (i = 0; i < n; i++) dd[i] = f[i] * h * h;
//
//    dd[0] -= a * (1 - (h / 2) * p);
//    dd[n - 1] -= b * (1 + (h / 2) * p);
//
//    solveMatrix(n, aa, bb, cc, dd, xx);
//
//    digits result(n);
//    for (int i = 0; i < n; i++) {
//        result[i] = abs(c[i] - (-x[i] * x[i] * x[i] / 2 - 3 * x[i] * x[i] / 8 - 7 * x[i] / 16 - 1599 / 1600 +
//                                1599 * pow(2.7, 4 * x[i]) / 1600));
//    }
//
//    for (int i = 1; i < n - 1; i++) {
//        cout << "diff: " << result[i] << endl;
//    }
//}
//
//int main() {
//    method();
//}
