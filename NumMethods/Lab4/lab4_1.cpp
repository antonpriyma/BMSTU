#include <iostream>
#include <cmath>
#include <vector>

using namespace std;
int n = 10;

float b, a;

typedef vector<float> digits;

using namespace std;

struct Res{
    float a;
    float b;
};


struct Res linear(float sum_a, float sum_b, float sum_c, float sum_d) {
    float a;
    float b;
    a = (sum_c - sum_b * sum_d) / (sum_a*n - (sum_b * sum_b));
    b = (sum_d - a * sum_b) / (n);

    struct Res res;
    res.a =a;
    res.b =b;
    return  res;
};

struct Res logar(float sum_a, float sum_b, float sum_c, float sum_d) {
    float a;
    float b;

    a = (sum_c - sum_b * sum_d/n)/(sum_a - sum_b * sum_b/n);
    b = (sum_d - a * sum_b)/(n);

    struct Res res;
    res.a =a;
    res.b =b;
    return  res;
}

void CalculateLinear() {
    digits x = {3, 4, 5, 6, 7, 8, 9, 10, 11, 12};

    digits y(n);
    for (int i = 0; i < n; i++) {
        y[i] = x[i] + 0.1;
    }

    vector<digits> data(4);
    for (auto &d: data) {
        d.resize(n);
    }


    for (int i = 0; i < n; i++) {
        data[0][i] = x[i] * x[i];
        data[1][i] = x[i];
        data[2][i] = x[i] * y[i];
        data[3][i] = y[i];
    }

    float sum_a = 0, sum_b = 0, sum_c = 0, sum_d = 0;

    for (int i = 0; i < n; i++) {
        sum_a += data[0][i];
        sum_b += data[1][i];
        sum_c += data[2][i];
        sum_d += data[3][i];
    }

   struct Res res = linear(sum_a,sum_b,sum_c, sum_d);

    cout << "Linear: " << endl;
    cout << "a: " << res.a << endl;
    cout << "b: " << res.b << endl;


}

void CalculateLog(){
    digits x = {1, 1.5, 1.7, 1.75, 1.8, 1.85, 1.86, 1.88, 1.89, 1.9};
    digits y(n);
    digits y0(n);
    for (int i = 0; i < n; i++){
        y0[i] = log(x[i]);
    }

    float x0[n];
    for (int i = 0; i < n; i++){
        x0[i] = log(x[i]);
    }

    vector<digits> data(4);
    for (auto &d: data) {
        d.resize(n);
    }

    for (int i = 0; i < n; i++){
        data[0][i] = x0[i] * x0[i];
        data[1][i] = x0[i];
        data[2][i] = x0[i] * y0[i];
        data[3][i] = y0[i];
    }

    float sum_a = 0, sum_b = 0, sum_c = 0, sum_d = 0;
    for (int i = 0; i < n; i++){
        sum_a += data[0][i];
        sum_b += data[1][i];
        sum_c += data[2][i];
        sum_d += data[3][i];
    }

    struct Res res;

    res = logar(sum_a, sum_b, sum_c, sum_d);
    cout << "Logar: " << endl;
    cout << "a: " << exp(res.b) << endl;
    cout << "b: " << res.a << endl;

    float x1 = 0, y1 = 0, z = 0, diff3 = 0, diff1 = 0, z1 = 0;


    x1 = (x[0]+x[n-1])/2;
    y1 = sqrt(y0[0]*y0[n-1]);
    z1 = res.a * x1 + b;
    z = exp(res.b) * log(x1)+res.a;

    diff1 = abs(z1 - y1);
    diff3 = abs(z - y1);
    cout << endl;
    cout << fixed << "Linear diff: " << diff1 << endl << "Log diff: " << diff3 <<" ";
    cout << endl;
}



int main(int argc, const char *argv[]) {
    cout.precision(20);
    CalculateLinear();
    CalculateLog();
    return 0;
}
