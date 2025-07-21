//! By convention, main.zig is where your main function lives in the case that
//! you are building an executable. If you are making a library, the convention
//! is to delete this file and start with root.zig instead.

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("Hello world! \n", .{});

    const out = std.io.getStdOut().writer();
    const err = std.io.getStdErr().writer();

    // 获取buffer//
    var out_buffer = std.io.bufferedWriter(out);
    var err_buffer = std.io.bufferedWriter(err);

    // 获取writer句柄//
    var out_writer = out_buffer.writer();
    var err_writer = err_buffer.writer();

    // 通过句柄写入buffer//
    try out_writer.print("Hello {s}!\n", .{"out"});
    try err_writer.print("Hello {s}!\n", .{"err"});

    // 尝试刷新buffer//
    try out_buffer.flush();
    try err_buffer.flush();
}

const std = @import("std");

/// This imports the separate module containing `root.zig`. Take a look in `build.zig` for details.
const lib = @import("hello_world_lib");
