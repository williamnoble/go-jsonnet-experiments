local newAnimal(myCatParams={}) = {
  local defaults = {
    name: 'Frogspawn',
    age: 80,
    x: 99,
  },
  local _config = defaults + myCatParams,

  name: _config.name,
  age: _config.age,
  x: std.get(myCatParams, 'age'),
};

{
  newAnimal:: newAnimal,
}
