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
	int r[3][N] = {0};
	int sum = 0;
	int i = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		int len = strlen(buf) - 1;
		for (int j = 0; j < len; j++) {
			int p = priority(buf[j]);
			r[i % 3][p] = 1;
		}
		if (i % 3 == 2) { /* group */
			for (int j = 0; j < N; j++) {
				if (r[0][j] && r[1][j] && r[2][j]) {
					sum += j;
					break;
				}
			}
			memset(r, 0, sizeof(r));
		}
		i++;
	}
	printf("%d\n", sum);
}
