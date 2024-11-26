# Go - Garbage Collector


# Runtime/MemStats

-   Alloc is bytes of allocated heap objects.
	Alloc uint64

-   TotalAlloc is cumulative bytes allocated for heap objects.
	TotalAlloc increases as heap objects are allocated, but unlike Alloc and HeapAlloc, it does not decrease when
	objects are freed.
    TotalAlloc uint64

-   Sys is the total bytes of memory obtained from the OS.
    Sys is the sum of the XSys fields below. Sys measures the virtual address space reserved by the Go runtime for the
    heap, stacks, and other internal data structures. It's likely that not all of the virtual address space is backed
    by physical memory at any given moment, though in general it all was at some point.
    Sys uint64

-   Lookups is the number of pointer lookups performed by the runtime.
    This is primarily useful for debugging runtime internals.
	Lookups uint64

-   Mallocs is the cumulative count of heap objects allocated. The number of live objects is Mallocs - Frees.
	Mallocs uint64

-   Frees is the cumulative count of heap objects freed.
	Frees uint64

-   HeapAlloc is bytes of allocated heap objects.
    "Allocated" heap objects include all reachable objects, as well as unreachable objects that the garbage collector has
    not yet freed. Specifically, HeapAlloc increases as heap objects are allocated and decreases as the heap is swept  and unreachable objects are freed. Sweeping occurs incrementally between GC cycles, so these two processes occur simultaneously, and as a result HeapAlloc tends to change smoothly (in contrast with the sawtooth that is typical of stop-the-world garbage collectors).
	HeapAlloc uint64

-   HeapSys is bytes of heap memory obtained from the OS.
	HeapSys measures the amount of virtual address space reserved for the heap. This includes virtual address space
	that has been reserved but not yet used, which consumes no physical memory, but tends to be small, as well as virtual
    address space for which the physical memory has been returned to the OS after it became unused (see HeapReleased
	for a measure of the latter).
	
	HeapSys estimates the largest size the heap has had.
	HeapSys uint64

-   HeapIdle is bytes in idle (unused) spans.
	
    Idle spans have no objects in them. These spans could be (and may already have been) returned to the OS, or they can be reused for heap allocations, or they can be reused as stack memory. HeapIdle minus HeapReleased estimates the amount of memory that could be returned to the OS, but is being retained by the runtime so it can grow the heap without requesting more memory from the OS. If this difference is significantly larger than the heap size, it indicates there was a recent transient spike in live heap size.
        HeapIdle uint64

-   HeapInuse is bytes in in-use spans.
	
	 In-use spans have at least one object in them. These spans
	 can only be used for other objects of roughly the same
	 size.
	
	 HeapInuse minus HeapAlloc estimates the amount of memory
	 that has been dedicated to particular size classes, but is
	 not currently being used. This is an upper bound on
	 fragmentation, but in general this memory can be reused
	 efficiently.
	HeapInuse uint64

-   HeapReleased is bytes of physical memory returned to the OS.
	
	 This counts heap memory from idle spans that was returned
	 to the OS and has not yet been reacquired for the heap.
	HeapReleased uint64

-   HeapObjects is the number of allocated heap objects.
	
	 Like HeapAlloc, this increases as objects are allocated and
	 decreases as the heap is swept and unreachable objects are
	 freed.
	HeapObjects uint64

	
-   GCSys is bytes of memory in garbage collection metadata.
	GCSys uint64

	 OtherSys is bytes of memory in miscellaneous off-heap
	 runtime allocations.
	OtherSys uint64

-   NextGC is the target heap size of the next GC cycle.
	
	 The garbage collector's goal is to keep HeapAlloc ≤ NextGC.
	 At the end of each GC cycle, the target for the next cycle
	 is computed based on the amount of reachable data and the
	 value of GOGC.
	NextGC uint64

- LastGC is the time the last garbage collection finished, as
	 nanoseconds since 1970 (the UNIX epoch).
	LastGC uint64

- PauseTotalNs is the cumulative nanoseconds in GC
	 stop-the-world pauses since the program started.
	
	 During a stop-the-world pause, all goroutines are paused
	 and only the garbage collector can run.
	PauseTotalNs uint64

- PauseNs is a circular buffer of recent GC stop-the-world
	 pause times in nanoseconds.
	
	 The most recent pause is at PauseNs[(NumGC+255)%256]. In
	 general, PauseNs[N%256] records the time paused in the most
	 recent N%256th GC cycle. There may be multiple pauses per
	 GC cycle; this is the sum of all pauses during a cycle.
	PauseNs [256]uint64

-   PauseEnd is a circular buffer of recent GC pause end times,
	 as nanoseconds since 1970 (the UNIX epoch).
	
	 This buffer is filled the same way as PauseNs. There may be
	 multiple pauses per GC cycle; this records the end of the
	 last pause in a cycle.
	PauseEnd [256]uint64

-   NumGC is the number of completed GC cycles.
	NumGC uint32

- NumForcedGC is the number of GC cycles that were forced by
	 the application calling the GC function.
	NumForcedGC uint32

-   GCCPUFraction is the fraction of this program's available
	 CPU time used by the GC since the program started.
	
	 GCCPUFraction is expressed as a number between 0 and 1,
	 where 0 means GC has consumed none of this program's CPU. A
	 program's available CPU time is defined as the integral of
	 GOMAXPROCS since the program started. That is, if
	 GOMAXPROCS is 2 and a program has been running for 10
	 seconds, its "available CPU" is 20 seconds. GCCPUFraction
	 does not include CPU time used for write barrier activity.
	
	 This is the same as the fraction of CPU reported by
	 GODEBUG=gctrace=1.
	GCCPUFraction float64