#include <iostream>

int main() {
    int n, left, right, leftmax, rightmax;
    std::cin >> n;

	long water;

    long heights[n];

    for (int i = 0; i < n; i++) {
        std::cin >> heights[i];
    }

    left = 0;
    right = n - 1;
    leftmax = 0;
    rightmax = 0;
    water = 0;

    for (; left < right;) {
        if (heights[left] <= heights[right]) {
            if (heights[left] >= leftmax) {
                leftmax = heights[left];
            } else {
                water += leftmax - heights[left];
            }
            left++;
        } else {
            if (heights[right] >= rightmax) {
                rightmax = heights[right];
            } else {
                water += rightmax - heights[right];
            }
            right--;
        }
    }

    std::cout << water << std::endl;
}
