local animalTemplate = import 'animal.jsonnet';

{
  local myCatParams = {
    name: 'Bruno',
    age: 3,
    redundant: 'lol',
  },

  local myCat = animalTemplate.newAnimal(myCatParams),
  local secondCat = animalTemplate.newAnimal(),
  cat: myCat,
  second: secondCat,
}
