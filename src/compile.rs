use std::fs;
use std::path::Path;
use std::process::Command;
use std::io::{Error, ErrorKind};

pub fn write_to_rust_file(transpiled_code: &str, file_name: &str) -> Result<(), Error> {
    // Determine the file extension based on the provided file_name
    let mut output_file = String::from(file_name);
    if let Some(extension) = Path::new(file_name).extension() {
        if let Some(ext_str) = extension.to_str() {
            // Replace the existing extension with ".rs" if it's not already ".rs"
            if ext_str != "rs" {
                output_file.push_str(".rs");
            }
        }
    } else {
        // If no extension is provided, append ".rs" to the file_name
        output_file.push_str(".rs");
    }

    // Write the transpiled_code to the determined output file
    fs::write(&output_file, transpiled_code)?;

    Ok(())
}

pub fn compile_to_executable(rust_file: &str, output_exe: &str) -> Result<(), Error> {
    let output = Command::new("rustc")
        .arg(rust_file)
        .arg("-o")
        .arg(output_exe) 
        .output()?;

    // Check if compilation was successful
    if output.status.success() {
        println!("Compilation successful!");
    } else {
        if let Some(code) = output.status.code() {
            eprintln!("Compilation failed with error code: {}", code);
        }
        if !output.stdout.is_empty() {
            eprintln!("Compiler output:\n{}", String::from_utf8_lossy(&output.stdout));
        }
        if !output.stderr.is_empty() {
            eprintln!("Compiler errors:\n{}", String::from_utf8_lossy(&output.stderr));
        }
        // Compilation failed, delete the temporary Rust file
        if let Err(err) = fs::remove_file(rust_file) {
            eprintln!("Failed to delete temporary Rust file: {}", err);
        }
        return Err(Error::new(ErrorKind::Other, "Compilation failed"));
    }

    Ok(())
}
