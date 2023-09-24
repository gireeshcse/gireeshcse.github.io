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

void pattern7(int n){
    int m = 1;
    for(int i = 0 ;i < n; i++){

        // Printing spaces
        for(int k=i;k<n-1;k++){
            std::cout << " ";
        }

        // printing *
        for(int j = 0 ; j < m; j++){
            std::cout << "*";
        }
        std::cout << "\n";
        m += 2;
    }
}

void pattern7_1(int n){
    for(int i = 1 ;i <= n; i++){

        // Printing spaces
        for(int k=i;k<n;k++){
            std::cout << " ";
        }

        // printing *
        for(int j = 0 ; j < (i*2)-1; j++){
            std::cout << "*";
        }
        std::cout << "\n";
    }
}

void pattern8(int n){
    for(int i = n ;i >= 1; i--){

        // Printing spaces
        for(int k=i;k<n;k++){
            std::cout << " ";
        }

        // printing *
        for(int j = 0 ; j < (i*2)-1; j++){
            std::cout << "*";
        }
        std::cout << "\n";
    }
}

void pattern9(int n){
    pattern7_1(10);
    pattern8(10);
}

void pattern10(int n){

    for(int i = 0;i<n;i++){
        for(int j = 0;j<=i;j++){
            std::cout << "*";
        }
        std::cout << "\n";
    }

    for(int i = n-1;i>0;i--){
        for(int j = 0;j<i;j++){
            std::cout << "*";
        }
        std::cout << "\n";
    }

}

int toggleV(int v){
    if(v == 0)
        return 1;

    return 0;
}

void pattern11(int n){

    int v = 1;
    for(int i = 0;i<n;i++){
        int k = v;
        for(int j = 0;j<=i;j++){
            std::cout << k;
            k = toggleV(k);
        }
        std::cout << "\n";
        v = toggleV(v);

    }
}

void pattern12(int n){

    for(int i = 1;i<=n;i++){
        for(int j = 1;j<=i;j++){
            std::cout << j;
        }

        for(int j=(n*2)-(i*2);j>0;j--){
            std::cout << " ";
        }

        for(int j = i;j>=1;j--){
            std::cout << j;
        }

        std::cout << "\n";
    }
}

void pattern13(int n){
    int k = 1;
    for(int i=1;i<=n;i++){
        for(int j=1;j<=i;j++){
            std::cout << k << " ";
            k++;
        }
        std::cout << "\n";
    }
}

void pattern14(int n){
    for(int i=0;i<n;i++){
        for(char j='A';j<=('A'+i);j++){
            std::cout << j;
        }
        std::cout << "\n";
    }
}

void pattern15(int n){
    for(int i=n;i>0;i--){
        for(char j='A';j<('A'+i);j++){
            std::cout << j;
        }
        std::cout << "\n";
    }
}

void pattern16(int n){
    char c = 'A';
    for(int i=0;i<n;i++){
        c = 'A' + i;
        for(char j=0;j<=i;j++){
            std::cout << c;
        }
        std::cout << "\n";
    }
}

void pattern17(int n){
    char c = 'A';
    for(int i=1;i<=n;i++){
        // print spaces
        for(int k=i;k<n;k++){
            std::cout << " ";
        }
        for(char j=1;j<=(2*i)-1;j++){
            if(j<=i){
                c = 'A'+j-1;
            }else{
                c = c-1;
            }
                
            std::cout << c;
        }
        std::cout << "\n";
    }
}

void pattern18(int n){
    char c = 'A'+(n-1);
    for(int i=1;i<=n;i++){
        char k = c;
        for(char j=1;j<=i;j++){
                
            std::cout << k;
            k += 1;
        }
        c = c -1;
        std::cout << "\n";
    }
}


void pattern19(int n){
    for(int i = 0 ;i < n; i++){
        for(int j = n ; j > i; j--){
            std::cout <<"*";
        }
        //print spaces
        for(int j=0;j<i*2;j++){
            std::cout <<" ";
        }
        for(int j = n ; j >i; j--){
            std::cout <<"*";
        }
        std::cout << "\n";
    }

    for(int i = n ;i > 0; i--){
        for(int j = n ; j >= i; j--){
            std::cout <<"*";
        }
        //print spaces
        for(int j=(i*2 - 2);j>0;j--){
            std::cout <<" ";
        }
        for(int j = n ; j >= i; j--){
            std::cout <<"*";
        }
        std::cout << "\n";
    }
}

void pattern20(int n){
    for(int i =1;i<=n;i++){
        for(int j=1;j<=i;j++){
            std::cout << "*";
        }
        // spaces
        for(int k=0;k<(n-i)*2;k++)
        {
            std::cout << " ";
        }

        for(int j=1;j<=i;j++){
            std::cout << "*";
        }
        
        std::cout << "\n";
    }
    for(int i =n-1;i>0;i--){
        for(int j=i;j>=1;j--){
            std::cout << "*";
        }

        // spaces
        for(int k=0;k<(n-i)*2;k++)
        {
            std::cout << " ";
        }

        for(int j=i;j>=1;j--){
            std::cout << "*";
        }
        std::cout << "\n";
    }
}


void pattern21(int n){
    for(int i=0;i<n;i++){
        if(i%n == 0 || i % n == n-1){
            for(int j=0;j<n;j++){
                std::cout << "*";
            }
        }else{
            std::cout << "*";
            //print spaces
            for(int j = 0;j<n-2;j++){
                std::cout << " ";
            }
            std::cout << "*";
        }
        std::cout << "\n";
    }
}

int minVal(int i,int j,int n){
    int top = i;
    int left = j;
    int right = (2*n) - 2 - j;
    int bottom = (2*n) - 2 - i;

    int min = top;
    if(left < min){
        min = left;
    }
    if(right < min){
        min = right;
    }
    if(bottom < min){
        min = bottom;
    }

    return min;
}

void pattern22(int n){
    for(int i = 0; i < (n*2)-1;i++){
        for(int j = 0; j < (n*2)-1;j++){
            
            std::cout << (n - minVal(i,j,n));
        }
        std::cout << "\n";
    }
}

int main(){

pattern22(4);

}