// a local isn't accessible outside the file
local newAnimalPrivate
(name, age, is_cat=true) = {
  name: name + 'private',
  age: age + 2,
  is_cat: is_cat,
};

// expose a method which calls the local
{
  callNewAnimal(name, age, is_cat)::
    newAnimalPrivate(name, age, is_cat),
}


// public, standard constructor
{
  newAnimal(name, age, is_cat=true):: {
    name: name,
    age: age,
    is_cat: is_cat,
  },
}

// public, standard constructor, concatenation/addition
{
  newAnimalTwo(name, age, is_cat=true):: {
    name: name + 'two',
    age: age + 2,
    is_cat: is_cat,
  },
}

