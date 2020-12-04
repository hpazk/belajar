void main() {
  List<int> numberList = [1, 2, 3, 4, 5];
  print(numberList);

  // dynamic list
  List profile = ['Hpazk', 26, true];
  print(profile[0]);

  // add to list
  numberList.add(6);
  print(numberList);

  // insert to list
  numberList.insert(1, 8);

  // update value
  numberList[0] = 0;

  // iteration with for
  for (int i = 0; i < numberList.length; i++) {
    print(i);
  }

  // iteration with foreach
  numberList.forEach((num) => print(num));

  /*
  .remove(<contain>)
  .removeAt(<index>)
  .removeLast()
  .removeRange(<index min>, <index.max>)
  */
}
