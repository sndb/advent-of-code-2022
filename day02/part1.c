#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

enum { BUFSIZE = 8 };

enum {
	ROCK,
	PAPER,
	SCISSORS,
};

int
shape(char s)
{
	switch (s) {
	case 'X':
	case 'A':
		return ROCK;
	case 'Y':
	case 'B':
		return PAPER;
	case 'Z':
	case 'C':
		return SCISSORS;
	}
	abort();
}

int
gamescore(int a, int b)
{
	if (a == b)
		return 3;
	if ((b + 1) % 3 == a)
		return 6;
	return 0;
}

int
main(void)
{
	char buf[BUFSIZE];
	int total = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		char a = shape(buf[2]);
		char b = shape(buf[0]);
		int score = a + 1 + gamescore(a, b);
		total += score;
	}
	printf("%d\n", total);
	return 0;
}
