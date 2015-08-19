fn main(){
    let mut x = 1;
    let y = x;
    x += 1;
    // Both values can be independently used
    println!("x + 1 = {} ,y : {}",x + 1,y); //因为实在栈上分配内存，而下面的例子是在堆上分配内存。

    // `a` is a pointer to a heap allocated integer
    let a = box 5;

    println!("a contains: {}", a);

    // "Move" `a` into `b`
    // Here's what happens under the hood: the pointer `a` gets copied (*not*
    // the data on the heap, just its address) into `b`. Now both are pointers
    // to the *same* heap allocated data. But now, `b` *owns* the heap
    // allocated data; `b` is now in charge of freeing the memory in the heap.
    let b = a;

    // After the previous move, `a` can no longer be used
    // Error! `a` can no longer access the data, because it no longer owns the
    // heap memory
    //println!("a contains: {}", a);
}