void baz() { }

void foo() { }

void bizz() { }

void buzz() {
  bizz();
}

int main() {
  foo();
  baz();
  buzz();
}

