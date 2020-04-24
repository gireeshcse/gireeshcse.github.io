### Nodejs

* 1 Process
* 1 thread
* 1 event loop
* 1 JS Engine(V8/libuv)
* 1 Nodejs engine

* CPU intensive code

    Blocks other process from being executed

In nodejs application try to avoid operations like synchronous network calls or infinite calls.
Only I/O operations run in parallel i.e asynchronously

Solution to the above is multiple process(Like Cluster API).But this results in shared memory and communication of data must be via JSON.

**Worker threads improve performance on CPU intensive operations not on I/O operations**

### Worker Threads

* One Process
* Multiple threads
* one event loop per thread 
* one JS engine instance per thread.
* one NodeJS instance per thread

Worker Threads have the following 

* ArrayBuffers

    To transfer memory from one thread to another.

* SharedArrayBuffer

    It lets us share memory between threads (limited to binary data).

* Atomics

    lets us do some processes execution concurrently, more efficiently and allows us to implement conditions variables in JavaScript

* MessagePort

    used for communicating between different threads. It can be used to transfer structured data, memory regions and other MessagePorts between different Workers.

* MessageChannel

    Represents an asynchronous, two-way communications channel used for communicating between different threads.

* WorkerData

     Is used to pass startup data. An arbitrary JavaScript value that contains a clone of the data passed to this thread’s Worker constructor. The data is cloned as if using postMessage()

### Example

        const { Worker } = require('worker_threads');

        const worker = new Worker(`
        const { parentPort } = require('worker_threads');
        parentPort.once('message',
            message => parentPort.postMessage({ pong: message }));  
        `, { eval: true });
        worker.on('message', message => console.log(message));      
        worker.postMessage('ping')

### Credits

[https://nodesource.com/blog/worker-threads-nodejs](https://nodesource.com/blog/worker-threads-nodejs)