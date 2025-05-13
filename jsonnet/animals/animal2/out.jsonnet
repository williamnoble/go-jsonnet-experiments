// create a local to act as a Class inside a file
// we call the class in our object
local newAnimal(name, age) = {
  name: name,
  age: age,
};

{
  x: newAnimal('foo', 21),
}
