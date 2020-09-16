//
// Created by Anton Priyma on 10.04.2020.
//
#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>

using namespace std;

void matrix(vector<double> &x, vector<double> &a, vector<double> &b, vector<double> &c,
            vector<double> &d) {

    int m = d.size();
    vector<double> alpha, beta;

    for (int i = 0; i < m; i++) {

        if (i == 0) {
            alpha.push_back((float) -c[i] / b[i]);
            beta.push_back((float) d[i] / b[i]);
        } else {
            alpha.push_back((float) -c[i] / (b[i] + alpha[i - 1] * a[i - 1]));
            beta.push_back((float) (d[i] - a[i - 1] * beta[i - 1]) / (b[i] + alpha[i - 1] * a[i - 1]));
        }
    }

    x.push_back(beta[m - 1]);

    for (int i = m - 2; i >= 0; i--) {

        double r = alpha[i] * beta[m - 1] + beta[i];

        x.push_back(r);

    }
    reverse(x.begin(), x.end());

}
