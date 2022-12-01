#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

enum { BUFSIZE = 32 };

int
icmp(const void *p, const void *q)
{
	int i = *(int *)p;
	int j = *(int *)q;
	if (i < j)
		return +1;
	if (i > j)
		return -1;
	return 0;
}

int
main(void)
{
	char buf[BUFSIZE];
	int calories[1000];
	int run = 0;
	int i = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		if (!isdigit(buf[0])) {
			calories[i++] = run;
			run = 0;
		} else {
			int n = atoi(buf);
			run += n;
		}
	}
	qsort(calories, i, sizeof(int), icmp);
	int top3 = calories[0] + calories[1] + calories[2];
	printf("%d\n", top3);
	return 0;
}
