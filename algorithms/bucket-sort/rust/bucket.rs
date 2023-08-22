use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    // Abrir o arquivo de texto
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

    print!("\nIniciando Execucao do Bucket Sort em Rust: \n\n");

    // Chamar o Bucket Sort para ordenar os números
    bucket_sort(&mut numbers);

    // Imprimir os números ordenados lado a lado
    /*     for (index, num) in numbers.iter().enumerate() {
            if index != 0 {
                print!(", ");
            }
            print!("{}", num);
        }
    } */
}

fn bucket_sort(numbers: &mut Vec<i32>) {
    // Definir o número máximo de buckets e criar um vetor bidimensional de buckets
    let num_buckets = 10;
    let mut buckets: Vec<Vec<i32>> = vec![Vec::new(); num_buckets];

    // Distribuir os números nos buckets
    for &num in numbers.iter() {
        let index = ((num as f64) * (num_buckets as f64) / (i32::MAX as f64)) as usize;
        buckets[index].push(num);
    }

    // Ordenar os buckets e colocar os números de volta no vetor
    let mut index = 0;
    for bucket in buckets.iter_mut() {
        bucket.sort();
        for &num in bucket.iter() {
            numbers[index] = num;
            index += 1;
        }
    }
}
