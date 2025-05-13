local animal = import 'AnimalClass.libsonnet';

{
  one: animal.newAnimal('firstExample', 1),
  two: animal.newAnimalTwo('secondExample', 21),
  outside: animal.callNewAnimal('thirdExample', 21, true),
}

# one calls newAnimal directly, it's not a local so it's exported implictly
