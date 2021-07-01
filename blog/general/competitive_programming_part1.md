## Competitive Programming

#### Primes

```
#include <iostream>
#include <cmath>

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
```

##### Sieve of Eratosthenes

To generate all prime from 2 to a given number n.

It begins by assumming that all numbers are prime.It takes the first prime number and removes all of its multiples. It then applies the same method to next prime number.
