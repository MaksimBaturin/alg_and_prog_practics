#include <iostream>
#include <string>

using namespace std;

int main(){
    string text, answer, temp;
    int counter;
    
    answer="";
    std::cin >> text;
    
    for (int i {0}; i <= text.length(); i++)
        {
        if (text[i] == text[i+1]){
            counter++;
        }
        else{
            counter++;
            if (counter == 1){
                answer+=text[i];
            }
            else{
                answer+=text[i];
                answer+=std::to_string(counter);
            }
            counter = 0;
        }
    }
    std::cout <<answer<< endl;
}
