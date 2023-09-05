const std = @import("std");
const io = @import("io");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("calories.txt", .{});
    defer file.close();

    var buffer: [1024]u8 = undefined;
    var stream = std.io.fixedBufferStream(&buffer);

    var i: usize = 0;
    while (i < 5) : (i += 1) {
        file.reader().streamUntilDelimiter(stream.writer(), '\n', null) catch |err| switch (err) {
            error.NoSpaceLeft => {
                break;
            },
            else => {
                return err;
            },
        };

        std.debug.print("{s}", .{buffer});
        std.debug.print("{d}", .{buffer[0..]});
        std.fmt.format(writer: anytype, comptime fmt: []const u8, args: anytype)
        std.fmt.formatBuf(buf: []const u8, options: FormatOptions, writer: anytype)
        // std.debug.print("{d}\n", .{try std.fmt.parseInt(i32, , 10)});
        buffer = undefined;
    }
    // try stream.reader().streamUntilDelimiter(stream.writer(), '\n', null);
}
