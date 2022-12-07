#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
#include <stdint.h>

enum { BUFSIZE = 1 << 8 };

const size_t DISK = 70000000;
const size_t NEED = 30000000;

struct file {
	char *name;
	size_t size;
};

struct dir {
	char *name;
	struct dir *parent;
	struct node *content;
	size_t size; // cache
};

enum node_type {
	NT_DIR,
	NT_FILE,
};

struct node {
	union {
		struct file *file;
		struct dir *dir;
	};
	enum node_type type;
	struct node *next;
};

struct dir *root = NULL;
struct dir *cwd = NULL;

struct dir *
dir_create(const char *name, struct dir *parent)
{
	struct dir *d = malloc(sizeof(*d));
	d->name = strdup(name);
	d->parent = parent;
	d->content = NULL;
	d->size = 0;
	return d;
}

struct node *
node_create(void *obj, enum node_type type)
{
	struct node *n = malloc(sizeof(*n));
	switch (type) {
	case NT_DIR:
		n->file = obj;
		break;
	case NT_FILE:
		n->dir = obj;
		break;
	default:
		assert(0);
	}
	n->type = type;
	n->next = NULL;
	return n;
}

struct file *
file_create(const char *name, size_t size)
{
	struct file *f = malloc(sizeof(*f));
	f->name = strdup(name);
	f->size = size;
	return f;
}

void
cd(const char *name)
{
	if (strcmp(name, "/") == 0) {
		if (!root) {
			root = dir_create(name, NULL);
		}
		cwd = root;
	} else if (strcmp(name, "..") == 0) {
		cwd = cwd->parent;
	} else {
		struct dir *d = dir_create(name, cwd);
		struct node *n = node_create(d, NT_DIR);
		if (!cwd->content) {
			cwd->content = n;
		} else {
			struct node *p;
			for (p = cwd->content; p->next; p = p->next)
				;
			p->next = n;
		}
		cwd = d;
	}
}

void
add_file(const char *name, size_t size)
{
	struct file *f = file_create(name, size);
	struct node *n = node_create(f, NT_FILE);
	if (!cwd->content) {
		cwd->content = n;
	} else {
		struct node *p;
		for (p = cwd->content; p->next; p = p->next)
			;
		p->next = n;
	}
}

size_t
cache_size(struct dir *d)
{
	size_t size = 0;
	struct node *p;
	for (p = d->content; p; p = p->next) {
		switch (p->type) {
		case NT_DIR:
			size += cache_size(p->dir);
			break;
		case NT_FILE:
			size += p->file->size;
			break;
		default:
			assert(0);
		}
	}
	d->size = size;
	return size;
}

int
starts_with(const char *prefix, const char *s)
{
	for (; *prefix; prefix++, s++) {
		if (*prefix != *s)
			return 0;
	}
	return 1;
}

void
part1(struct dir *d, size_t *s)
{
	if (d->size <= 100000) {
		*s += d->size;
	}
	struct node *p;
	for (p = d->content; p; p = p->next) {
		if (p->type == NT_DIR) {
			part1(p->dir, s);
		}
	}
}

void
part2(struct dir *d, size_t *s)
{
	if (DISK - root->size + d->size >= NEED && d->size < *s)
		*s = d->size;
	struct node *p;
	for (p = d->content; p; p = p->next) {
		if (p->type == NT_DIR) {
			part2(p->dir, s);
		}
	}
}

int
main(void)
{
	char buf[BUFSIZE];
	while (fgets(buf, BUFSIZE, stdin)) {
		buf[strlen(buf) - 1] = '\0';
		if (starts_with("$ cd ", buf)) {
			cd(buf + 5);
		} else if (isdigit(buf[0])) {
			size_t size;
			char name[BUFSIZE];
			sscanf(buf, "%zu %s", &size, name);
			add_file(name, size);
		}
	}
	cache_size(root);
	size_t sum1 = 0;
	size_t sum2 = SIZE_MAX;
	part1(root, &sum1);
	part2(root, &sum2);
	printf("%zu %zu\n", sum1, sum2);
	return 0;
}
