# important!


Value Types             |       Reference Types
================================================
int                     |           slices  
float                   |           maps
string                  |           channels
bool                    |           pointers
struct                  |           functions



for reference types you don't need to worry about updating the original address location,
when argument is passed in a function.


This is similar to Java.

Where Objects, HashMaps etc, are passing the copy of the value to the function.
Java and Go are pass by value.

Only thing to remember is, the object that is passed to the function, the value of the contents would be
Meta data.
-----------
addressLocationReferencePointer: &afad2342address
lenght: 123223
capacity: 232
...

The copy of the value itself is a pointer to a memory location in heap.
