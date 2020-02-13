    __global__

It tells CUDA c++ compiler that this function runs on the GPU and can be called from CPU Code.

    __global__
    void add(int n, float *x, float *y)
    {
        for(int i=0;i<n;i++)
        {
            y[i] = x[i] + y[i];
        }
    }