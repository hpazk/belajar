// basic function
void main() {
  greet();
  // invoking profile function
  profile('Hpazk', 26);

  // // invoking average function
  var avg = average(10, 5);
  print(avg);

  // invoking greetNewUser function
  greetNewUser('Zkh', 25, true);
  greetNewUser('Ltg', 25);

  greetNewUser2(name: 'Ltg', age: 25);
}

void greet() {
  print('Hello World');
}

// function parameters
void profile(String name, age) {
  print('your name is $name, age $age');
}

// function return type
double average(num val1, num val2) {
  return (val1 + val2) / 2;
}

// positional optional parameters
void greetNewUser([String name, int age, isVerified]) {
  print('Hello $name, age $age, verified: $isVerified');
}

// named optional parameters and default parameter
void greetNewUser2({String name, int age, bool isVerified = false}) {
  //isVerified is default vaule parameter
  print('Hello $name, age $age, verified: $isVerified');
}
