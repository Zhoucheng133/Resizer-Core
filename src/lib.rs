use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use image::imageops::FilterType;

fn c_str_to_string(ptr: *const c_char) -> Option<String> {
    if ptr.is_null() { return None; }
    unsafe {
        Some(CStr::from_ptr(ptr).to_string_lossy().into_owned())
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn GetSize(
    path: *const c_char,
) -> *mut c_char {
    if path.is_null() {
        return std::ptr::null_mut();
    }
    let c_str = unsafe { CStr::from_ptr(path) };
    let r_str = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return std::ptr::null_mut(),
    };

    match image::image_dimensions(r_str) {
        Ok((width, height)) => {
            let size_str = format!("{}x{}", width, height);
            match CString::new(size_str) {
                Ok(c_string) => c_string.into_raw(),
                Err(_) => std::ptr::null_mut(),
            }
        }
        Err(_) => std::ptr::null_mut(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn Resize(
    path: *const c_char,
    width: i32,
    height: i32,
    output: *const c_char,
) -> *mut c_char {
    let input_path = match c_str_to_string(path) {
        Some(s) => s,
        None => return CString::new("Error: Null input path").unwrap().into_raw(),
    };
    let out_path = match c_str_to_string(output) {
        Some(s) => s,
        None => return CString::new("Error: Null output path").unwrap().into_raw(),
    };

    if width <= 0 || height <= 0 {
        return CString::new("Error: Invalid dimensions").unwrap().into_raw();
    }

    let result = (|| -> Result<(), Box<dyn std::error::Error>> {
        let img = image::open(input_path)?;
        let scaled = img.resize(width as u32, height as u32, FilterType::Lanczos3);
        scaled.save(out_path)?;
        Ok(())
    })();

    match result {
        Ok(_) => CString::new("OK").unwrap().into_raw(),
        Err(e) => CString::new(format!("Error: {}", e)).unwrap().into_raw(),
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn free_string(s: *mut c_char) {
    unsafe {
        if s.is_null() { return; }
        let _ = CString::from_raw(s);
    }
}