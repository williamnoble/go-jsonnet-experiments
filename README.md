# jsonnet-experiments

## go-jsonnet
- basics
- skeleton: tiny skeleton used as a POC basis for writing tests
- evaluating:
  - snippets: evaluate both literal and constructed snippets.
  - lib: experimenting with unit-testing single fns as a lib. 

## jsonnet
animals: experimenting importing functions. Output via `jsonnet out.jsonnet`
- animal: constructor visibility +- locals; use of `local` vs `{}`.
- animal2: private `local` constructor exported (slightly better).
- animal3-nice: private `local` constructor, we expose the 'Class', use a fn in a lib.
- animal4: defaults and open params via `{}`. Not very explicit, kinda ugly:(.

externals: passing in vars via a pipeline `jsonnet out.jsonnet --ext-str env=dev`

