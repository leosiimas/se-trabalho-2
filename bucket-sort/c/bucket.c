#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// #define NBUCKET 5   // Number of buckets
#define NARRAY 1000

int getMax(int a[], int n) // function to get maximum element from the given array
{
  int max = a[0];
  for (int i = 1; i < n; i++)
    if (a[i] > max)
      max = a[i];
  return max;
}

void bucket(int a[], int n) // function to implement bucket sort
{
  int max = getMax(a, n); // max is the maximum element of array
  int bucket[max], i;
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

  printf("\nIniciando Execucao do Bucket Sort em C: \n\n");
  bucket(numbers, size);
  printArr(numbers, size);

  free(numbers);
}