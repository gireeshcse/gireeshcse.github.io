#include<stdio.h>
#include "usedbyC.h"

int main(int argc, char **argv){
    GoInt x = 12;
    GoInt y = 14;

    printf("About to call go function !");
    HelloGo();

    GoInt result = Multiply(x,y);

    printf("result : %d\n",(int)result);
    printf("It worked\n");
    return 0;
}