#include <iostream>
#include <cmath>
#include <cstring>

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
    // bool prime[n+1] is wrong since it returns local variable 
    bool *prime = new bool[n+1];
    // void* memset( void* str, int ch, size_t n);
    // fill a byte string with a byte value
    // memset works byte by byte
    // for int it can set 0 or -1
    memset(prime, true, sizeof(prime)*(n+1));
    
    prime[0] = false;
    prime[1] = false;

    int m = sqrt(n);
    for(int p = 2; p <= m; p++){
        if(prime[p]){
            // if it is prime then eliminate the multiples of that number
            // k start square of prime because less than it are already marked 
            for(int k = p*p;k<= n;k+=p){
                prime[k] = false;
            }
        }
    }

    return prime;
}

int gcd(int a, int b){
    int min = std::min(a,b);
    int result = 1;
    for(int i = min;i >= 2;i--){
        if((a % i == 0) && (b % i == 0)){
            return i;
        }
    }
    return 1;
}

int gcd_euclid(int a, int b){
    if(b == 0) return a;
    return gcd_euclid(b, a % b);
}

int lcm(int a, int b){
    return (a *b ) / gcd_euclid(a,b);
}

int main(){

    int num,option,num1,num2;
    cout << "\n1. Check a number is prime or not \n2. Prime numbers upto N(Sieve of Eratosthenes)" 
        << "\n3.Normal GCD\n4.Euclid's GCD\n5. LCM\nPlease select the option:";
    cin >> option;
    switch (option)
    {
        case 1:{
            cout << "\nPlease enter a number:";
            cin >> num;
            if(isPrime(num)){
                cout << "\n" << num << " is a prime";
            }else{
                cout << "\n" << num << " is not a prime";
            }
            break;
        }
        case 2:{
            cout << "\nPlease enter a number(N):";
            cin >> num;
            bool *primes = sieve(num);
            for(int i=0;i<=num;i++){
                if(primes[i]){
                    cout << i << " ";
                }
            }
            cout << "\n";
            delete[] primes; //allocated memory must be deleted
            break;
        }
        case 3:{
            cout << "\nPlease enter a & b  numbers:";
            cin >> num1 >> num2;
            cout << gcd(num1,num2) << " is the GCD of " << num1 << " and " << num2 <<  "\n";
            break;
        }
        case 4:{
            cout << "\nPlease enter a & b  numbers:";
            cin >> num1 >> num2;
            cout << gcd_euclid(num1,num2) << " is the GCD of " << num1 << " and " << num2 <<  "\n";
            break;
        }
        case 5:{
            cout << "\nPlease enter a & b  numbers:";
            cin >> num1 >> num2;
            cout << lcm(num1,num2) << " is the LCM of " << num1 << " and " << num2 <<  "\n";
            break;
        }
        default:
            cout << "\n Invalid option";
            break;
    }
    return 0;
}

