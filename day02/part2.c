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

enum {
	LOSE,
	DRAW,
	WIN,
};

int
shape(char s)
{
	switch (s) {
	case 'A':
		return ROCK;
	case 'B':
		return PAPER;
	case 'C':
		return SCISSORS;
	}
	abort();
}

int
outcome(char o)
{
	switch (o) {
	case 'X':
		return LOSE;
	case 'Y':
		return DRAW;
	case 'Z':
		return WIN;
	}
	abort();
}

int
lose(char s, char o)
{
	return (s + (o + 2)) % 3;
}

int
main(void)
{
	char buf[BUFSIZE];
	int total = 0;
	while (fgets(buf, BUFSIZE, stdin)) {
		char s = shape(buf[0]);
		char o = outcome(buf[2]);
		int score = o * 3 + lose(s, o) + 1;
		total += score;
	}
	printf("%d\n", total);
	return 0;
}
