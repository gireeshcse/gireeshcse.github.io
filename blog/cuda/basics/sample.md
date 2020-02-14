    __global__

It tells CUDA c++ compiler that this function runs on the GPU and can be called from CPU Code.

    /**
    * CUDA kernel function to add the elements of 2 arrays on the GPU
    */
    __global__
    void add(int n, float *x, float *y)
    {
        for(int i=0;i<n;i++)
        {
            y[i] = x[i] + y[i];
        }
    }

### Memory Allocation in C

#### malloc()

* **malloc()** allocates memory block of given size (in bytes) and returns a pointer to the beginning of the block. 
* **malloc()** doesn’t initialize the allocated memory.

    void* malloc(size_t size); 

#### calloc()

**calloc()** allocates the memory and also initializes the allocated memory block to zero.

    void* calloc(size_t num, size_t size);