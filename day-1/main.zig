const std = @import("std");
const io = @import("io");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("calories.txt", .{});
    defer file.close();
    var buf_reader = 
}
