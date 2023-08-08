// Counting sort in C programming

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void countingSort(int arr[], int n)
{
  int max = arr[0];
  for (int i = 1; i < n; i++)
  {
    if (arr[i] > max)
    {
      max = arr[i];
    }
  }

  int *count = (int *)calloc(max + 1, sizeof(int));
  int *output = (int *)malloc(n * sizeof(int));

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

void printArr(int a[], int n)
{
  for (int i = 0; i < n; ++i)
    printf("%d ", a[i]);

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

  printf("\nIniciando Execucao do Counting Sort em C: \n\n");
  countingSort(numbers, size);
  printArr(numbers, size);

  free(numbers);
}