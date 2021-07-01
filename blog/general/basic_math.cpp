#include <iostream>
#include <cmath>
#include <array>

using namespace std;

bool isPrime(int n){

    if(n <= 1) return false;

    if(n == 2) return true;

    if(n % 2 == 0) return false;

    int m = sqrt(n);

    for(int i = 3; i <= m; i += 2){
        if (n % i == 0){
            return false;
        }
    } 

    return true;
}

bool* sieve(int n){
    std::array<bool,sizeof(bool) * (n+1)> prime;
    prime.fill(true)

    return prime;
}

int main(){

    int num;
    cout << "Please enter a number:";
    cin >> num;
    if (isPrime(num) == true){
        cout << "\n " << num << " is prime\n";
    }else{
        cout << "\n " << num << " is not prime\n";
    }
}

