local U2 = import 'u2.libsonnet';
local utils = import 'utils.libsonnet';

{
  foo: utils.append_foo('F is for '),
  bar: utils.append_bar('B is for '),
  x: U2,
}
