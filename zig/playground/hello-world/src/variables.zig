// const :声明常量
// @import: 导入包或者zig文件
// @ 内置函数以@开头
const std = @import("std");

pub fn main() !void {
    // undefined 表示未初始化
    var variable: u16 = undefined;
    variable = 6;
    std.debug.print("变量是 {}\n", .{variable});
}

// 解构 表达式只能出现在块内（不在容器范围内）
fn destructing() !void {
    const print = @import("std").debug.print;
    // {} 块 用于限制变量作用域
    {
        var x: u32 = undefined;
        var y: u32 = undefined;
        var z: u32 = undefined;
        // 元组
        const tuple = .{ 1, 2, 3 };
        // 解构元组
        x, y, z = tuple;

        print("tuple: x = {}, y = {}, z = {}\n", .{ x, y, z });
        // 数组
        const array = [_]u32{ 4, 5, 6 };
        // 解构数组
        x, y, z = array;

        print("array: x = {}, y = {}, z = {}\n", .{ x, y, z });
        // 向量定义
        const vector: @Vector(3, u32) = .{ 7, 8, 9 };
        // 解构向量
        x, y, z = vector;

        print("vector: x = {}, y = {}, z = {}\n", .{ x, y, z });
    }
    {
        var x: u32 = undefined;

        const tuple = .{ 1, 2, 3 };

        x, var y: u32, const z = tuple;

        print("x = {}, y = {}, z = {}\n", .{ x, y, z });

        // y 可变
        y = 100;

        // 可以用 _ 丢弃不想要的值
        _, x, _ = tuple;

        print("x = {}", .{x});
    }
}

// 块
fn block() !void {
    // 限制变量作用域
    // 块也可以是一个表达式。当块带有标签时，break 语句可以从块中返回一个值。
    var y: i32 = 123;

    const x = blk: {
        y += 1;
        //使用 break 标签 返回一个值
        break :blk y;
    };

    std.debug.print("{}", .{x});
}

// 容器
//在 Zig 中，容器 是充当命名空间的任何语法结构，用于保存变量和函数声明。容器也可以是可实例化的类型定义。结构体、枚举、联合、不透明类型，甚至 Zig 源文件本身都是容器。然而，容器不能包含语句（语句是描述程序运行操作的一个单位）。

// 注释
// 普通注释 /// 文档注释 //! 顶层文档注释

/// usingnamespace 直接讲一个容器中的pub声明混入到当前容器中
fn using_namespace() !void {
    const T = struct {
        // 将std混入到当前容器T中
        usingnamespace @import("std");
    };
    // T也有了debug
    T.debug.print("helloworld", .{});
}

/// threadlocal
threadlocal var xx: i32 = 1234;
fn threadlocal_1() !void {
    const thread1 = try std.Thread.spawn(.{}, testTls, .{});
    const thread2 = try std.Thread.spawn(.{}, testTls, .{});
    testTls();
    thread1.join();
    thread2.join();
}
fn testTls() void {
    // 1234
    std.debug.print("x is {}\n", .{xx});
    xx += 1;
    // 1235
    std.debug.print("x is {}\n", .{xx});
}
