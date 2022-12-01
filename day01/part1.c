#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

enum { BUFSIZE = 32 };

int
main(void)
{
	char buf[BUFSIZE];
	int max = 0;
	int run = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		if (!isdigit(buf[0])) {
			if (run > max)
				max = run;
			run = 0;
		} else {
			int n = atoi(buf);
			run += n;
		}
	}
	printf("%d\n", max);
	return 0;
}
