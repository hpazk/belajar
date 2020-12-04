import 'dart:io';

void main() {
  stdout.write('Enter your name: ');
  var name = stdin.readLineSync();

  stdout.write('Enter your Age: ');
  var age = int.parse(stdin.readLineSync());

  print('your name is $name, age $age');
}
