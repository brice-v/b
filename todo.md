### TODO

- [ ] Add the rest of the obvious tokens from `b.b` file and test
- [ ] Figure out comments
- [ ] Figure out which operators could be removed
- [ ] Begin work on parser
  - [ ] Still need to figure out what to do with types
  - maybe in the parser phase it will be easy to ignore the types until a sort of type check phase
  - TBD
- [ ] Work on making code easily testable
- [ ] Begin thinking about how exactly things should translate
  - [ ] Sets (Map with bool key)
  - [ ] Objs (Want obj to be like js with properties of go)
- [ ] Document the current lexer properly before moving on to next phase
- [ ] Breakup lexer test into groups according to docs (that need to be created ^)
- [ ] Write test for repl