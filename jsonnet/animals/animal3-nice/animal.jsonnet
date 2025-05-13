local Animal(name, age) = {
  // local variable
  local n_eyes = 6,

  // reference with self.
  foo:: 3,

  // function accessible by self.
  f:: function(foo) {
    f: foo,
  },

  name: name,
  age: age,
  eyes: n_eyes,
  num: self.foo,
  l: self.f('foo'),
};

{
  newAnimal:: Animal,
}
