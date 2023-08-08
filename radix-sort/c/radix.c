#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int findMax(int arr[], int n)
{
  int max = arr[0];
  for (int i = 1; i < n; i++)
  {
    if (arr[i] > max)
    {
      max = arr[i];
    }
  }
  return max;
}

void countingSortByDigit(int arr[], int n, int exp)
{
  int output[n];
  int count[10] = {0};

  for (int i = 0; i < n; i++)
  {
    count[(arr[i] / exp) % 10]++;
  }

  for (int i = 1; i < 10; i++)
  {
    count[i] += count[i - 1];
  }

  for (int i = n - 1; i >= 0; i--)
  {
    output[count[(arr[i] / exp) % 10] - 1] = arr[i];
    count[(arr[i] / exp) % 10]--;
  }

  for (int i = 0; i < n; i++)
  {
    arr[i] = output[i];
  }
}

void radixSort(int arr[], int n)
{
  int max = findMax(arr, n);

  for (int exp = 1; max / exp > 0; exp *= 10)
  {
    countingSortByDigit(arr, n, exp);
  }
}

void printArray(int array[], int size)
{
  for (int i = 0; i < size; ++i)
  {
    printf("%d  ", array[i]);
  }
  printf("\n");
}

int main()
{
  FILE *file;
  int number;
  int *numbers = NULL;
  int size = 0;
  char line[1000];
  int i;

  file = fopen("./number-generator/numbers.txt", "rb");
  if (NULL == file)
  {
    printf("Unable to open file\n");
    exit(-1);
  }

  while (fgets(line, sizeof(line), file) != NULL)
  {
    char *token = strtok(line, ",");
    while (token != NULL)
    {
      size++;
      numbers = (int *)realloc(numbers, size * sizeof(int));
      if (numbers == NULL)
      {
        printf("Unable to aloc memory\n");
        return 1;
      }
      numbers[size - 1] = atoi(token);
      token = strtok(NULL, ",");
    }
  }

  fclose(file);

  printf("\nIniciando Execucao do Radix Sort em C: \n\n");
  radixSort(number, size);
  printArray(number, size);

  free(numbers);
}