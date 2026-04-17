

// It finds objects that are no longer used and frees their memory.
Without GC:

Memory leaks : heaf memory stays alive , not like stack one, so if no GC, it increse the memory usage , and kill the application
Manual free() errors

Go uses:
✔ Concurrent
Runs alongside your program (not fully stop-the-world)

✔ Tri-color mark-and-sweep
Core algorithm

1. Mark (find live objects)
2. Sweep (remove dead objects)
3. Allocate new memory

Objects are divided into 3 sets:
Color	Meaning
White	: Not visited (candidate for deletion)
Grey	: Discovered but not fully scanned
Black	: Fully scanned (safe )

✔ Stack:
Local variables
Managed automatically
No GC needed
✔ Heap:
Dynamically allocated data (new, make)
Managed by GC


How to reduce GC pressure?
Reduce allocations
Reuse memory
Use sync.Pool
var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}
Reuses objects → reduces GC pressure
