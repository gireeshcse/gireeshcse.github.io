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

    // To launch the add() kernel, which invokes it on the GPU
    add<<<1, 1>>>(N, x, y);

### Memory Allocation in C

#### malloc()

* **malloc()** allocates memory block of given size (in bytes) and returns a pointer to the beginning of the block. 
* **malloc()** doesn’t initialize the allocated memory.

        void* malloc(size_t size);
        arr = (int*)malloc(5 * sizeof(int)); // 5*4bytes = 20 bytes  

#### calloc()

* **calloc()** allocates the memory and also initializes the allocated memory block to zero.

        void* calloc(size_t num, size_t size);
        arr = (int*)calloc(5, sizeof(int));

**Return Value:** 

After successful allocation in malloc() and calloc(), a **pointer** to the block of memory is returned otherwise **NULL** value is returned which indicates the failure of allocation.

        // Deallocates memory previously allocated 
        free(arr); 

### Memory Allocation in C++

C++ supports **calloc()** and **malloc()** functions and also has two operators **new** and **delete** that perform the task of allocating and freeing the memory in a better and easier way.

        int N = 1<<20; // 1M elements
        float *x = new float[N];
        int *p = new int(25); // Intialize with 25

        int *q = NULL; 
        q = new int;   

        // Free memory
        delete [] x;
        delete p;
        delete q;

### Memory Allocation in CUDA

**Unified Memory**

* Single memory space accessible by all GPU's and CPU's in our system.	

        // Allocate Unified Memory -- accessible from CPU or GPU
        float *x, *y;
        cudaMallocManaged(&x, N*sizeof(float));
        cudaMallocManaged(&y, N*sizeof(float));

        // Free memory
        cudaFree(x);
        cudaFree(y);

* **race condition** can occur since multiple parallel threads would both read and write the same locations.So care has to be taken otherwise the output may be inapproriate.

### cudaDeviceSynchronize()

* Blocks until the device has completed all preceding requested tasks.Otherwise CPU continues to execute remaining while GPU is computing.
* returns an error if one of the preceding tasks has failed
* Wait for compute device to finish.
* Called in CPU code.

Although CUDA kernel launches are asynchronous, all GPU-related tasks placed in one stream (which is the default behavior) are executed sequentially.

        kernel1<<<X,Y>>>(...); // kernel start execution, CPU continues to next statement
        kernel2<<<X,Y>>>(...); // kernel is placed in queue and will start after kernel1 finishes, CPU continues to next statement
        cudaMemcpy(...); // CPU blocks until memory is copied, memory copy starts only after kernel2 finishes

using **cudaDeviceSynchronize()** is appropriate would be when you have several cudaStreams running, and you would like to have them exchange some information.


    // Wait for GPU to finish before accessing on host
    cudaDeviceSynchronize();

### Important Info

* CUDA files have the file extension .cu
* **nvcc**, the CUDA C++ compiler.

#### Compilation

    nvcc add.cu -o add_cuda

#### Profiling

    nvprof ./add_cuda

### Threads

**Execution configuration**: It tells how many parallel threads to use.

`<<<no_of_thread_blocks,threads_in_thread_block>>>`

#### Important 

* **gridDim.x**     : No of blocks in the grid.
* **blockIdx.x**    : index of current thread block.
* **blockDim.x**    : No of threads in a block(current exection thread block)
* **threadIdx.x**   : index of current thread in current block

#### Examples


        // For 256 parallel threads.

        add<<<1, 256>>>(N, x, y);
        __global__
        void add(int n, float *x, float *y)
        {
        int index = threadIdx.x;
        int stride = blockDim.x;// 256
        for (int i = index; i < n; i += stride)
            y[i] = x[i] + y[i];
        }

CUDA GPUs have many parallel processors grouped into Streaming Multiprocessors, or SMs. Each SM can run multiple concurrent thread blocks.

        int blockSize = 256;
        int numBlocks = (N + blockSize - 1) / blockSize;
        add<<<numBlocks, blockSize>>>(N, x, y);

index = blockIdx.x * blockDim.x + threadIdx.x;

index = 2 * 256 + 3 = 515

        __global__
        void add(int n, float *x, float *y)
        {
        int index = blockIdx.x * blockDim.x + threadIdx.x;
        int stride = blockDim.x * gridDim.x; // total number of threads in the grid
        for (int i = index; i < n; i += stride)
            y[i] = x[i] + y[i];
        }
### Scalability and Thread Use

No.of blocks - Multiple of number of Streaming Multiprocessors (SMs) on the device to balance the utilization.

    int nDevices;
    cudaGetDeviceCount(&nDevices);

    int numSMs;
    cudaDeviceGetAttribute(&numSMs,cudaDevAttrMultiProcessorCount,deviceId);
    add<<<32*numSMs,256>>>(1<<20,x,y);

Serial Host Implementation (Generally for Debugging)

    add<<<1,1>>>(1<<20,x,y);

### Querying device properties

    int nDevices,i,numSMs;
	cudaGetDeviceCount(&nDevices);
	for(i=0; i<nDevices; i++){
	
		cudaDeviceProp prop;
		cudaGetDeviceProperties(&prop,i);
		cudaDeviceGetAttribute(&numSMs,cudaDevAttrMultiProcessorCount,i);
		cout << "\n" << prop.name << "\n Clock Rate: " << prop.memoryClockRate
			 << "\n MemoryBusWidth: " << prop.memoryBusWidth << "\n Num of SMs: " << numSMs << "\n\n";

	}

    
### Caliculating total SMs in a system.

    int numSMs,nDevices,totalSMs = 0;
    cudaGetDeviceCount(&nDevices);
    for(i = 0; i < nDevices;i++)
    {
        cudaDeviceGetAttribute(&numSMs,cudaDevAttrMultiProcessorCount,i);
        totalSMs += numSMs;	
    }
    std::cout << "\n Total SMS " << totalSMs

### Hardware Implementation

The NVIDIA GPU architecture is built around a scalable array of multithreaded Streaming Multiprocessors (SMs). When a CUDA program on the host CPU invokes a kernel grid, the blocks of the grid are enumerated and distributed to multiprocessors with available execution capacity. The threads of a thread block execute concurrently on one multiprocessor, and multiple thread blocks can execute concurrently on one multiprocessor. As thread blocks terminate, new blocks are launched on the vacated multiprocessors.


### Wrap Size

Warp size is the number of threads in a warp.
The block of threads is actually divided into sub-blocks called "warps".
### Device Info

|   Device	|   Compute Capability	| Micro-architecture  	| Wrap Size  	| SMs|   Maximum number of threads per block	|
|--:	    |--:	                |:-:                   	|:-:	        |:-: |:-:	                                    |
|Tesla K40  |   3.5	                |   Kepler              |   32          | 15 |	       1024                     |
|GeForce940M|   5.0	                |   Maxwell             |   32	        |    |	       1024                     |
|   	    |   	                |   	                |   	        |    |	                                |





