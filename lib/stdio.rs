#[macro_export]
macro_rules! printf {
    ($($arg:tt)*) => {
        println!($($arg)*);
    };
}