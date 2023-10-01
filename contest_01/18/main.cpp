
#include <iostream>

using namespace std;

int main(){
    int n;
    std::cin >> n;
    string temp = "";
    while(n>0) {
        char s=((n-1)%26+'A');
		temp=s+temp;
		n=(n-1)/26;
    }
    std::cout << temp;
}
