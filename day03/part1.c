#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

enum {
	BUFSIZE = 100,
	N = 53,
};

int
priority(char c)
{
	if (c >= 'a' && c <= 'z')
		return c - 'a' + 1;
	if (c >= 'A' && c <= 'Z')
		return c - 'A' + 27;
	assert(0);
}

int
main(void)
{
	char buf[BUFSIZE];
	int sum = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		int c[N] = {0};
		int len = strlen(buf) - 1;
		for (int i = 0; i < len; i++) {
			int p = priority(buf[i]);
			if (i < len / 2) { /* first compartment */
				c[p] = 1;
			} else if (c[p]) { /* second compartment */
				sum += p;
				break;
			}
		}
	}
	printf("%d\n", sum);
}
