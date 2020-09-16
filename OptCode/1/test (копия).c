#include <stdio.h>
#include <stdlib.h>

void phi() {
    int w, x, y, z;
    scanf("%d", &x);
    x = x - 3;
    if (x < 3) {
        y = x * 2;
        w = y;
    } else {
        y = x - 3;
    }
    w = x - y;
    z = x + y;
    printf("w: %d, x: %d, y: %d, z: %d",
           w, x, y, z);
}

int main() {
	
	// Arithmetic expressions
	int a, b, c, d;
    scanf("%d%d%d%d", &a, &b, &c, &d);
    int e = a * b + c / d;
    int f = e ^ (a << 2);
    float qq = 3.14;
    float qqq = 2.00600e+003;
    printf("%d %d\n", e, f);
	
	// Phi functions
	phi();
	
	// Memory
	int n;
    scanf("%d", &n);
    int *arr = (int*) malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        arr[i] = i * i;
	}
	for (int i = 0; i < n; i++) {
		printf("%d", arr[i]);
	}
	printf("\n");
	
	int *ref = &n;
	int n2 = *ref;

    return 0;
    
}

