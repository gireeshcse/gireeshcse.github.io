#include<stdio.h>

#include "callC.h"

void cHello(){
    printf("Hello from C!");
}

void printMessage(char* message){
    printf("Message from Go : %s\n",message);
}