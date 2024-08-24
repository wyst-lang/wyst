<p align="center">
<img src="img/wyst-bg.jpg" style="height: 128px">
</p>

<h1 align="center">Wyst
</h1>

<h3 align="center">Wyst is a programing language that is inspired by C and transpiles to Go</h3>

```c
include io;

void main() {
  printf("Hello, World!");
}
```

# For developers

If the grammar has been updated run: `antlr -Dlanguage=Go -o parser Wyst.g4`
