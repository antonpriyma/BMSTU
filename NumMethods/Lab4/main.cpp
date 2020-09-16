//#include <iostream>
//#include <vector>
//#include <cmath>
//
//using namespace std;
//
//typedef vector<float> digits;
//
//void matrix(digits &x, digits &a, digits &b, digits &c,
//            digits &d) {
//
//    int m = d.size();
//    digits alpha, beta;
//
//    for (int i = 0; i < m; i++) {
//
//        if (i == 0) {
//            alpha.push_back((float) -c[i] / b[i]);
//            beta.push_back((float) d[i] / b[i]);
//        } else {
//            alpha.push_back((float) -c[i] / (b[i] + alpha[i - 1] * a[i - 1]));
//            beta.push_back((float) (d[i] - a[i - 1] * beta[i - 1]) / (b[i] + alpha[i - 1] * a[i - 1]));
//        }
//    }
//
//    x.push_back(beta[m - 1]);
//
//    for (int i = m - 2; i >= 0; i--) {
//
//        double r = alpha[i] * beta[m - 1] + beta[i];
//
//        x.push_back(r);
//
//    }
//    reverse(x.begin(), x.end());
//
//}
//
//float calculate(vector<digits> coef, float x, float a, float H) {
//    int i = 0;
//    while (!(a + i * H <= x && x < a + ((i + 1) * H))) {
//        i++;
//    }
//
//    float Xi = a + i * H;
//    return coef[0][i] + coef[2][i] * (x - Xi) + coef[1][i] * pow((x - Xi), 2) + coef[3][i] * pow((x - Xi), 3);
//}
//
//vector<digits> count_coefs(float H, digits y) {
//    int N = y.size() - 1;
//
//    // a
//    vector<digits> res(4);
//    digits d(N + 1);
//    for (int i = 0; i < N + 1; ++i) {
//        d[i] = y[i];
//    }
//    res[0] = d;
//    d.clear();
//
//
//    // c
//    digits a(N - 1, 4.0);
//    digits b(N - 2, 1.0);
//    digits c(N - 2, 1.0);
//    digits d0(N - 1);
//
//    for (int i = 0; i < N - 1; ++i) {
//        d0[i] = 3 *
//                (y[i + 2] - 2 * y[i + 1] + y[i]) / (H * H);
//    }
//
//    digits solved;
//    matrix(solved, a, b, c, d0);
//
//    d.resize(N + 1);
//    for (int i = 1; i < N; ++i) {
//        d[i] = solved[i - 1];
//    }
//    res[1] = d;
//    d.clear();
//
//    // b
//    d.resize(N);
//    for (int i = 0; i < N; ++i) {
//        d[i] = (y[i + 1] - y[i]) / H -
//               H * (res[1][i + 1] + 2 * res[1][i]) / 3;
//    }
//    res[2] = d;
//    d.clear();
//
//
//    // d
//    d.resize(N);
//    for (int i = 0; i < N; ++i) {
//        d[i] = (res[1][i + 1] - res[1][i]) / (3 * H);
//    }
//    res[3] = d;
//
//    return res;
//}
//
//int main() {
//    int N = 5;
//    digits y(N + 1);
//    float A = 0, B = 1;
//    float H = (B - A) / N, V = 0;
//
//    for (auto &y0 : y) {
//        y0 = sin(V);
//        V += H;
//    }
//
//    auto coefs = count_coefs(H, y);
//
//
//    for (float i = 0; i < N + 0.1; i += 0.5) {
//        float x = A + H * i;
//        float y0 = sin(x);
//        float yy = calculate(coefs, x, A, H);
//        float delta = abs(y0 - yy);
//        cout << "x = " << x << "   ";
//        cout << "y0 = " << y0 << "   ";
//        cout << "y* = " << yy << "   ";
//        cout << "diff = " << delta << endl;
//    }
//}
