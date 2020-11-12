void main() {
  var kucing = Animal('kitty', 2, 2.5);
  print(kucing.name);
  kucing.sleep();
  kucing.eat();
  print(kucing.weight);
}

class Animal {
  String name;
  int age;
  double weight;

  Animal(this.name, this.age, this.weight);

  void eat() {
    print('$name is eating.');
    weight = weight + 0.2;
  }

  void sleep() {
    print('$name is sleeping.');
  }

  void poop() {
    print('$name is pooping.');
    weight = weight - 0.1;
  }
}
