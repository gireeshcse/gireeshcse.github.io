#include <iostream>
#include <cmath>
#include <cstring>

void pattern1(int n){
    for(int i = 0 ;i < n; i++){
        for(int j = 0 ; j < n; j++){
            std::cout << "* ";
        }
        std::cout << "\n";
    }
}

void pattern2(int n){
    for(int i = 0 ;i < n; i++){
        for(int j = 0 ; j <= i; j++){
            std::cout << "* ";
        }
        std::cout << "\n";
    }
}

void pattern3(int n){
    for(int i = 1 ;i <= n; i++){
        for(int j = 1 ; j <= i; j++){
            std::cout << j <<" ";
        }
        std::cout << "\n";
    }
}

void pattern4(int n){
    for(int i = 1 ;i <= n; i++){
        for(int j = 1 ; j <= i; j++){
            std::cout << i <<" ";
        }
        std::cout << "\n";
    }
}

void pattern5(int n){
    for(int i = 1 ;i <= n; i++){
        for(int j = i ; j <= n; j++){
            std::cout <<"* ";
        }
        std::cout << "\n";
    }
}

void pattern5_1(int n){
    for(int i = 0 ;i < n; i++){
        for(int j = 0 ; j < n - i  ; j++){
            std::cout <<"* ";
        }
        std::cout << "\n";
    }
}

void pattern6(int n){
    for(int i = 1 ;i <= n; i++){
        for(int j = 1 ; j <= (n - i + 1); j++){
            std::cout << j << " ";
        }
        std::cout << "\n";
    }
}

int main(){

pattern6(10);

}