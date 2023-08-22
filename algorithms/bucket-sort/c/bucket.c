#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

// #define NBUCKET 5   // Number of buckets
#define NARRAY 1000

long int getMax(long int a[], int n)
{
  long int max = a[0];
  for (int i = 1; i < n; i++)
    if (a[i] > max)
      max = a[i];
  return max;
}

void bucket(long int a[], int n)
{
  long int max = getMax(a, n);
  long int *bucket = (long int *)malloc((max + 1) * sizeof(long int)); // Allocate bucket array
  if (bucket == NULL)
  {
    printf("Unable to allocate memory\n");
    return;
  }

  for (int i = 0; i <= max; i++)
  {
    bucket[i] = 0;
  }
  for (int i = 0; i < n; i++)
  {
    bucket[a[i]]++;
  }
  for (int i = 0, j = 0; i <= max; i++)
  {
    while (bucket[i] > 0)
    {
      a[j++] = i;
      bucket[i]--;
    }
  }

  free(bucket); // Free allocated memory
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

  printf("\nIniciando Execucao do Bucket Sort em C: \n");
  bucket(numbers, size);

  // printArr(numbers, size);

  free(numbers);

  return 0;
}