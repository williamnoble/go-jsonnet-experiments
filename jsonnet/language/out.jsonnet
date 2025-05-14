local U2 = import 'ex.jsonnet';
local utils = import 'utils.libsonnet';

{
  // testing visibilibility with ::
  foo: utils.append_foo('F is for '),

  private_add: utils.private_add(1,2),
  public_add: utils.public_add(2,4),
  x: U2,
}
