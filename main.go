package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"trab2/metrics"
)

func main() {
	paths := [9]metrics.PathValue{
		{Text: "bucket-sort", OutputPath: "output/c/", Addrs: "./algorithms/bucket-sort/c/bucket.o", Language: "C"},
		{Text: "bucket-sort", OutputPath: "output/go/", Addrs: "./algorithms/bucket-sort/go/bucket.o", Language: "GO"},
		{Text: "bucket-sort", OutputPath: "output/rust/", Addrs: "./algorithms/bucket-sort/rust/bucket.o", Language: "Rust"},
		{Text: "counting-sort", OutputPath: "output/c/", Addrs: "./algorithms/counting-sort/c/counting.o", Language: "C"},
		{Text: "counting-sort", OutputPath: "output/go/", Addrs: "./algorithms/counting-sort/go/counting.o", Language: "GO"},
		{Text: "counting-sort", OutputPath: "output/rust/", Addrs: "./algorithms/counting-sort/rust/counting.o", Language: "Rust"},
		{Text: "radix-sort", OutputPath: "output/c/", Addrs: "./algorithms/radix-sort/c/radix.o", Language: "C"},
		{Text: "radix-sort", OutputPath: "output/go/", Addrs: "./algorithms/radix-sort/go/radix.o", Language: "GO"},
		{Text: "radix-sort", OutputPath: "output/rust/", Addrs: "./algorithms/radix-sort/rust/radix.o", Language: "Rust"},
	}

	for _, path := range paths {
		done := make(chan bool)

		// Inicia a medição em uma go routine em background
		go metrics.HardwareMetrics(done, path)

		// Executa o Código
		execute(path)

		// Ao final da execução avisa a goroutine para parar a execução
		done <- true

		// Espera 30 segundos para a próxima execução, motivo: deixar a temperatura do computador Baixar
		time.Sleep(30 * time.Second)
	}
}

func execute(path metrics.PathValue) {
	cmd := exec.Command(path.Addrs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o código binário: ", err)
		panic("Erro ao executar o código do endereço: " + path.Addrs)
	}
}
