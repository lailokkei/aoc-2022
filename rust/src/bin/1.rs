use std::{
    fs::File,
    io::{BufRead, BufReader},
};

struct Max {
    max: i32,
    current: i32,
}

impl Max {
    fn new() -> Max {
        return Max { max: 0, current: 0 };
    }

    fn add(&mut self, food: i32) {
        self.current += food;
    }

    fn end(&mut self) {
        if self.current > self.max {
            self.max = self.current
        }
        self.current = 0;
    }
}

fn main() {
    let file = File::open("data/1.txt").expect("file missing");
    let reader = BufReader::new(file);

    let mut max = Max::new();

    for line in reader.lines() {
        let thing = line.unwrap();
        if thing.is_empty() {
            max.end()
        } else {
            let int_thing = thing.parse().unwrap();
            max.add(int_thing);
        }
    }

    println!("{}", max.max);
}
