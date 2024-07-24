use std::alloc::{alloc, Layout};
use std::ptr::null_mut;

// Malloc
pub fn _0x6d616c6c6f63(size: usize) -> *mut u8 {
    if size == 0 {
        return null_mut();
    }

    // Create a memory layout for the requested size
    let layout = Layout::from_size_align(size, std::mem::align_of::<u8>()).unwrap();

    // Allocate the memory
    unsafe {
        let ptr = alloc(layout);
        if ptr.is_null() {
            // Allocation failed
            null_mut()
        } else {
            ptr
        }
    }
}
