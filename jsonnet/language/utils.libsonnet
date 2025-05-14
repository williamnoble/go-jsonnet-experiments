local private_add(a, b) = a + b;

{
  // fields with :: are hidden from resulting JSON output but can be exported
  // append_foo is syntatic sugar for: append_foo :: function(a) body
  // we don't really need ::, if our lib is a file we can use : (and allow public visibility)
  append_foo(a)::
    a + 'foo',

  public_add(a, b)::
    a + b,

  private_add: private_add,

  // Visible function using shorthand
  short_add(a, b): a + b,

  // Visible function using explicit function keyword
  literal_greet: function(name) 'Hello, ' + name + '!',

  formatGreeting(name, formal):
      if formal then
        'Greetings, ' + name + '.'
      else
        'Hey, ' + name + '!',
}
