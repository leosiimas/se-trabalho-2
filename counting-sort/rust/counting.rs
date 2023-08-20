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

    // Chamar o Counting Sort para ordenar os números
    counting_sort(&mut numbers);

    // Imprimir os números ordenados lado a lado
    for (index, num) in numbers.iter().enumerate() {
        if index != 0 {
            print!(", ");
        }
        print!("{}", num);
    }
}

fn counting_sort(numbers: &mut Vec<i32>) {
    if numbers.is_empty() {
        return;
    }

    let max_num = *numbers.iter().max().unwrap();
    let min_num = *numbers.iter().min().unwrap();
    let range = (max_num - min_num + 1) as usize;

    let mut count_array = vec![0; range];
    let mut output_array = vec![0; numbers.len()];

    // Contar a ocorrência de cada número
    for &num in numbers.iter() {
        count_array[(num - min_num) as usize] += 1;
    }

    // Calcular as posições finais no array de saída
    for i in 1..range {
        count_array[i] += count_array[i - 1];
    }

    // Construir o array de saída ordenado
    for &num in numbers.iter().rev() {
        let index = (num - min_num) as usize;
        output_array[count_array[index] - 1] = num;
        count_array[index] -= 1;
    }

    // Copiar o array ordenado de volta para o vetor original
    numbers.copy_from_slice(&output_array);
}
