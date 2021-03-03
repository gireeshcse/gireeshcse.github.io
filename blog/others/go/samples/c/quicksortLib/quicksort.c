#include<stdio.h>
#include "quicksort.h"

// void main(){

// 	int arr[10] = {51,1,86,74,5,8,77,5,74,40};

// 	int low = 0;
//  	int high = 9;
// 	quicksort(arr , low  , high);
// 	for(low = 0; low <= high; low++){
// 		printf("%d\t",arr[low]);
// 	}
	
// }

void quicksort(int *arr, int low ,int high){
	int part_index;
	if(low < high){
		part_index = partition(arr, low, high);
		quicksort(arr,low,part_index-1);
		quicksort(arr,part_index+1,high);
		
	}

}

void swap(int *num1, int *num2)
{
	int temp = *num2;
	*num2 = *num1;
	*num1 = temp;
}
int partition(int *arr, int low, int high){

	int pivot = arr[high];
	int i = low;
	int j;
	for(j=low; j < high;j++){
		if(arr[j] < pivot)
		{
			//i++;			
			swap(&arr[i],&arr[j]);
			i++;
		}
	}
	swap(&arr[i],&arr[high]);
	return i;
	
}

int add(int a, int b){
	return a + b;
}
