## Competitive Programming

#### Primes

```
// #include <cmath>
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
```

##### Sieve of Eratosthenes

To generate all prime from 2 to a given number n.

It begins by assumming that all numbers are prime.It takes the first prime number and removes all of its multiples. It then applies the same method to next prime number.

```
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
```

#### GCD(Greatest Common Divisor)

 GCD of two numbers a and b is the greatest number that divides evenly into both a and b.

 ```
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
```

 ##### Euclid’s algorithm

* We begin by expressing the larger number in terms of the smaller number plus a remainder: 
* wE iterate over the two numbers until a remainder of 0 is found.

Example: GCD of 2336 and 1314
2336 = 1314 x 1 + 1022
We now do the same with 1314 and 1022:
1314 = 1022 x 1 + 292
We continue this process until we reach a remainder of 0:
1022 = 292 x 3 + 146
292 = 146 x 2 + 0

```
int gcd_euclid(int a, int b){
    if(b == 0) return a;
    return gcd_euclid(b, a % b);
}
```

##### LCM(Least common Multiple) 

* LCM of 6 and 9 is 18

```
int lcm(int a, int b){
    return (a *b ) / gcd_euclid(a,b);
}
```



####  Credits

[Top Coder](https://www.topcoder.com/thrive/articles/Mathematics%20for%20Topcoders)