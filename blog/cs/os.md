# Operating System.

* Controls the execution of applications.
* Acts as interface b/w the Hardware and Applications.
* Provides services to users and a platform for programs to run on.

## Types of OS

* Batched OS (Ex: Payroll System, Transactions Process, etc.)

    The batch processing is used where the more set up time was reduced by processing jobs in batches where similar jobs were run together.

* Multi-Programmed OS (Ex: Windows, UNIX)

* Timesharing OS (Ex: Multics)

* Distributed OS (Ex: LOCUS, etc.)

    System software over a collection of independent, networked, communication, and physically separate computational nodes. They handle jobs which are serviced by multiple CPUs.

* Real-Time OS (Ex: PSOS, VRTX etc.)

    These are event-driven and preemptive, meaning the OS is capable of monitoring the relevant priority of competing tasks and make changes to the task priority.

OS is visualised as Resource manager

* Resources

    - Hardware

        * CPU
        * Memory
        * I/O
        * Registers

    - Software

        * Device Drivers
        * Semaphores
        * Msg
        * Monitors


### Objectives/Goals 

* Convenience (User friendly)
* Efficiency (Resources Utilization) 
* Ability to evolve.
* Portability
* Reliability
* Scalability

### Functions

* Resource Management
* Memory Management
* Process Management
* File Management
* Error Detection
* Security


User interacts with os through command interpreter. It translates the user inputs/request into System Call Interface(SCI)/Application Program Interface(API)


### Command Interpreter 

* A that interprets the commands of the user.
* Acts as interface between the program and Kernel

### MultiProgram

In Main memory, multiple programs will be stored at any given point of time, while one program is being processed, other waits in queue , once a process is either completed or goes for I/O or preempted, the process picks up some other program to process.

#### Advantages

* It enables multi-tasking
* CPU utilization is high 

#### Types of Multiprogramming

* Non-Preemptive

    Ensures that a process relinquishes control of the CPU only when it finishes with its current CPU burst.

    Example: DOS, Windows 3.0

    - Voluntary Release
    - Completion
    - I/O, System Call
    - Event Services

    Disadv

    - Starvation
    - Priority not possible.

* Preemptive

    Running process can be forced to release .

    - Forcefully
    - Complete
    - I/O
    - Timer, Priority

#### Architecture support for multiprogramming

* DMA (Direct Memory Access)

    Process can directly access the memory for I/O operation without processor intervention with the help of DMA controller.

* Address Translation

    CPU issues a virtual/logical memory address. MMU (Memory Management Unit) translates it into actual physical address of the MM.

* 2-mode operation

    CPU has 2 different modes to run the processes.

    - User Mode

        - Non-Privileged / Pre-emptive
        - All user programs, all application programs.
    
    - Kernel Mode

        - System/supervisory/Priviliged/Protected/Non-Premptive
        - Process/Service routine, OS services (Scheduler, dispatcher etc.)


    Program Status Word (PSW) or Process Status Word:

    - A bit in PSW commonly referred as the command bit/mode bit indicates whether it is running on UM(0) or KM(1).
    
    Note: System calls are usually invoked by using s/w interrupts.

    fork() -> system call -> s/w interrupts -> Interrupt Service Routine (ISR)

    Interrupt handler controls a table called "dispatch table" which has all the entries for every system call.

    * **fork()**
        - System call
        - Creates a separate & duplicate process
        - contains copy of address space of original process (communication b/w parent and child)
        - Process identifier new process is returned to the parent & zero to the child process.
        - **exec()** system call is used after the fork() by one of the processes to replace process memory space with a new program (including threads). program id as a parameter

        - n forks - creates total of 2^n processes and 2^n -1 child processes
            ```
            #include<stdio.h>
            #include<unistd.h>
                int main(){
                int id = -1;
                fork();
                fork();
                id = fork();
                printf("\nHello - %d\n",id); // printed 8 times
                return 0;
            }
            ```
            Output
            ```

            Hello - 310635

            Hello - 0


            Hello - 310637
            Hello - 310638

            Hello - 0

            Hello - 0

            Hello - 310639

            Hello - 0
            ```

## Multiprocessor system

* It involves the processing of different computer programs at the same time mostly by a computer system with two or more CPUs that are sharing single memory.

* Increases throughput and is cost-effective also as all processors share the same resources.

## RAID (Redundant Arrays of Independent Disks)

* Data storage virtualization technology that combines multiple hard disks.
* Used to improve overall performance and reliability of data storage.
* Main purpose is to achieve data redundancy to reduce data loss.

### Levels

#### Level 6 (P+Q Redundancy)

* 

#### Level 5 (Block interleaved distributed Parity or Block-Level stripping with Distributed Parity)

* Parity rotates among the drives.
* RAID 5 requires at least three disks.

#### Level 4 (Block interleaved parity or Block-Level stripping with Dedicated Parity)

* Instead of duplicating data, this adopts a parity-based approach.
* One disk is dedicated to parity.
* Parity is caliculated using XOR function. XOR(0,0,0,1) = 1 XOR(0,1,1,0) = 0. Even number 1's ones results in parity 0 and an odd number of 1's results in parity 1.
* Reliability: 1 (1 Disk failure)
* Capacity: (N-1) * B

#### Level 3 (Bit interleaved parity)

#### Level 2 (Memory style error correcting code)

#### Level 1 (Mirroring and Duplexing)

* More than one copy of each block is stored in a separate disk. Thus, every block has two (or more ) copies, lying on different disks.

* For Mirroring Level - 2 
    
    - Reliability: 1 to N/2
    - Capacity: N * B / 2 (Only half the space is being used to store data.) 


#### Level 0 (Non-Redundant Striping)

* Blocks are stripped(spread) across disks.
* The entire space is being used to store data. Since there is no duplication, N disks each having B blocks are fully utilized.



### Credits 

[geekforgeeks](https://www.geeksforgeeks.org/raid-redundant-arrays-of-independent-disks/)
[interviewbit](https://www.interviewbit.com/operating-system-interview-questions/)
[wikipedia](https://en.wikipedia.org/wiki/Standard_RAID_levels)