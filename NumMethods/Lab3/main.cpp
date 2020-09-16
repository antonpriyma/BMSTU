#include <iostream>
#include <math.h>
#include <vector>


int main() {
//    std::vector<float> x = {0.1, 0.198, 0.297, 0.401, 0.497, 0.599, 0.701, 0.799, 0.900, 1.089};
//    std::vector<float> y = {0.1, 0.199, 0.303, 0.400, 0.499, 0.600, 0.701, 0.803, 0.903, 1.089};

    const int aCoef = 2;
    const int bCoef = 1;
    std::vector<float > x = {2, 7, 20, 54, 148};
    std::vector<float > y ;

    std::vector<float> xLog;

    for (float xi: x){
        xLog.push_back(log(xi));
        y.push_back(aCoef*log(xi) + bCoef);
    }

    float sumA = 0, sumB = 0, C = 0, D = 0, N = x.size();

    for (int i = 0; i < N; i++) {
        sumA += xLog[i];
        C += xLog[i] * y[i];
        sumB += y[i];
        D += xLog[i] * xLog[i];
    }

    float a_temp = (N * C - (sumA * sumB)) / ((N * D) - (sumA * sumA));
    float b_temp = (sumB - a_temp * sumA) / N;

    float a = exp(b_temp);
    float b = a_temp;

    std::cout << "y = x*" << a << " + " << b << std::endl;
    //  std::cout << a << "\n" << b << std::endl;
    for (int i = 0; i < N; i++) {
        std::cout << "(" << xLog[i] << "," << y[i] << ")" << std::endl;
    }



    return 0;
}

