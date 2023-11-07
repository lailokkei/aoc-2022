use std::{
    cmp::Reverse,
    collections::BinaryHeap,
    fs::{self},
};

fn sum_top_k(lines: &String, k: i32) -> i32 {
    let mut top_k = BinaryHeap::new();
    let mut current = 0;

    for line in lines.lines() {
        if line.is_empty() {
            top_k.push(Reverse(current));

            if top_k.len() as i32 > k {
                top_k.pop();
            }
            current = 0;
        } else {
            let int_thing: i32 = line.parse().unwrap();
            current += int_thing;
        }
    }

    let mut sum = 0;

    for idk in top_k.iter() {
        sum += idk.0;
    }

    return sum;
}

fn main() {
    let mut contents = fs::read_to_string("data/1.txt").unwrap();
    contents += "\n";

    println!("{}", sum_top_k(&contents, 1));
    println!("{}", sum_top_k(&contents, 3));
}
