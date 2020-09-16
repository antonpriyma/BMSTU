//#include <iostream>
//
//using namespace std;
//
//int main() {
//    int N;
//    cin >> N;
//    int a, b, c;
//    int i;
//    int min1 = 0, min2 = 0, summ = 0, ming = 10000;
//    for (i = 0; i < N; i++) {
//        cin >> a >> b >> c;
//
//        if ((a < b) && (a < c)) {
//            min1 = a;
//
//            if ((b < c) && (b - min1) % 5 != 0) {
//                min2 = b;
//            }
//
//            if ((c < b) && (b - min1) % 5 != 0) {
//                min2 = c;
//            }
//        }
//
//        if ((b < a) && (b < c)) {
//            min1 = b;
//
//            if ((a < c) && (a - min1) % 5 != 0) {
//                min2 = a;
//            }
//
//            if ((c < a) && (c - min1) % 5 != 0) {
//                min2 = c;
//            }
//
//        }
//
//        if ((c < a) && (c < b)) {
//            min1 = c;
//
//            if ((a < b) && (a - min1) % 5 != 0) {
//                min2 = a;
//            }
//
//            if ((b < a) && (b - min1) % 5 != 0) {
//                min2 = b;
//            }
//        }
//
//        summ += min1;
//        if ((min2 != 0) && (min2 - min1 < ming)) {
//            ming = min2 - min1;
//        }
//    }
//
//    if (summ % 5 != 0) {
//        cout << summ;
//    } else if ((summ % 5 == 0) && (ming != 10000)) {
//        cout << summ + ming;
//    } else {
//        cout << 0;
//    }
//
//    return 0;
//}
//
//
