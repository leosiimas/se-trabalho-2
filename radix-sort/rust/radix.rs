use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file =
        File::open("./number-generator/numbers.txt").expect("Não foi possível abrir o arquivo");
    let reader = BufReader::new(file);

    // Ler os números do arquivo e armazená-los em um vetor
    let mut numbers: Vec<i32> = Vec::new();
    for line in reader.lines() {
        if let Ok(line) = line {
            let parts: Vec<&str> = line.split(',').collect();
            for part in parts {
                if let Ok(num) = part.trim().parse::<i32>() {
                    numbers.push(num);
                }
            }
        }
    }

    // Chamar o Radix Sort para ordenar os números
    radix_sort(&mut numbers);

    // Imprimir os números ordenados
    // Imprimir os números ordenados lado a lado
    for (index, num) in numbers.iter().enumerate() {
        if index != 0 {
            print!(", ");
        }
        print!("{}", num);
    }
}

fn radix_sort(numbers: &mut Vec<i32>) {
    if numbers.is_empty() {
        return;
    }

    let max_num = *numbers.iter().max().unwrap();
    let num_digits = num_digits(max_num);

    for digit in 0..num_digits {
        counting_sort_by_digit(numbers, digit);
    }
}

fn num_digits(num: i32) -> usize {
    if num == 0 {
        return 1;
    }

    let mut count = 0;
    let mut num = num;
    while num != 0 {
        num /= 10;
        count += 1;
    }
    count
}

fn counting_sort_by_digit(numbers: &mut Vec<i32>, digit: usize) {
    let mut count_array = vec![0; 10];
    let mut output_array = vec![0; numbers.len()];

    for &num in numbers.iter() {
        let digit_val = (num / 10_i32.pow(digit as u32)) % 10;
        count_array[digit_val as usize] += 1;
    }

    for i in 1..10 {
        count_array[i] += count_array[i - 1];
    }

    for &num in numbers.iter().rev() {
        let digit_val = (num / 10_i32.pow(digit as u32)) % 10;
        output_array[count_array[digit_val as usize] as usize - 1] = num;
        count_array[digit_val as usize] -= 1;
    }

    numbers.copy_from_slice(&output_array);
}
