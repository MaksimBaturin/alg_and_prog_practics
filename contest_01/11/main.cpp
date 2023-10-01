#include <iostream>

int main(){
    int a, b, c;
    std::cin >> a >> b >> c;
    if (abs(a-b)<abs(a-c))
    {
        std::cout << "B " << abs(a-b);
    }
    else
    {
        std::cout << "C " << abs(a-c);
    }
}
