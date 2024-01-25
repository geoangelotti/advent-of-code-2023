#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>

#define BUFFER_COUNT 1024

#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define MIN(a, b) ((a) < (b) ? (a) : (b))

long processpart1(char *filepath)
{
	long product = 1;
	long *values;
	values = (long *)calloc(0, sizeof(long));
	size_t size = 0;
	long acc = 0;
	int in = 0;
	char buffer[1];
	int f = open(filepath, O_RDONLY);
	while (read(f, &buffer, 1) > 0)
	{
		if (buffer[0] >= '0' && buffer[0] <= '9')
		{
			if (in == 0)
			{
				in = 1;
			}
			acc = acc * 10 + buffer[0] - '0';
		}
		else
		{
			if (in)
			{
				size += 1;
				values = realloc(values, size * sizeof(long));
				values[size - 1] = acc;
				acc = 0;
				in = 0;
			}
		}
	}
	close(f);
	for (int i = 0; i < size / 2; i++)
	{
		long time = values[i];
		long distance = values[size / 2 + i];
		long timeheld = 0;
		for (long j = 0; j <= time; j++)
		{
			timeheld = j;
			if (timeheld * (time - timeheld) > distance)
			{
				break;
			}
		}
		product *= time - 2 * timeheld + 1;
	}
	return product;
}

int main()
{
	printf("%ld\n", processpart1("input.txt"));
	return 0;
}