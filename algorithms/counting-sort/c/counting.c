// Counting sort in C programming
#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void countingSort(long int arr[], int n)
{
  long int max = arr[0];
  for (int i = 1; i < n; i++)
  {
    if (arr[i] > max)
    {
      max = arr[i];
    }
  }

  long int *count = (long int *)calloc(max + 1, sizeof(long int));
  long int *output = (long int *)malloc(n * sizeof(long int));

  if (!count || !output)
  {
    printf("Erro de alocação de memória!\n");
    return;
  }

  for (int i = 0; i < n; i++)
  {
    count[arr[i]]++;
  }

  for (int i = 1; i <= max; i++)
  {
    count[i] += count[i - 1];
  }

  for (int i = n - 1; i >= 0; i--)
  {
    output[count[arr[i]] - 1] = arr[i];
    count[arr[i]]--;
  }

  for (int i = 0; i < n; i++)
  {
    arr[i] = output[i];
  }

  free(count);
  free(output);
}

void printArr(long int a[], int n)
{
  for (int i = 0; i < n; ++i)
    printf("%ld ", a[i]);

  printf("\n");
}
int main()
{
  FILE *file;
  long int number; // Use um tipo adequado para seus números
  long int *numbers = NULL;
  long int capacity = 1000; // Tamanho inicial do array
  long int size = 0;

  file = fopen("./number-generator/numbers.txt", "rb");
  if (NULL == file)
  {
    printf("Unable to open file\n");
    exit(-1);
  }

  numbers = (long int *)malloc(capacity * sizeof(long int)); // Alocar espaço inicial
  if (numbers == NULL)
  {
    printf("Unable to allocate memory\n");
    return 1;
  }

  while (fscanf(file, "%ld,", &number) != EOF)
  {
    if (size >= capacity)
    { // Realocar em incrementos maiores
      capacity *= 2;
      numbers = (long int *)realloc(numbers, capacity * sizeof(long int));
      if (numbers == NULL)
      {
        printf("Unable to reallocate memory\n");
        return 1;
      }
    }
    numbers[size] = number;
    size++;
  }

  fclose(file);

  printf("\nIniciando Execucao do Counting Sort em C: \n");
  countingSort(numbers, size);
  /*   printArr(numbers, size); */

  free(numbers);
}