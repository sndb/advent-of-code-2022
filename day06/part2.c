#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>

enum {
	N = 14,
	MAPSIZE = 1 << 8,
};

int
main(void)
{
	int map[MAPSIZE] = {0};
	int c;
	int i = 0;
	int j = 0;
	while (i != N && (c = getchar()) != EOF && c != '\n') {
		int d = j - map[c];
		if (d < N && d <= i) {
			i = j - map[c];
		} else {
			i++;
		}
		map[c] = j;
		j++;
	}
	printf("%d\n", j);
	return 0;
}
