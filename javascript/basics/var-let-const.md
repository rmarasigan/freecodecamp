# `var`, `let`, `const`

### Differences between `var` and `let`
One of the biggest problems with declaring variables with the `var` keyword is that you can easily overwrite variable declarations.

```javascript
var x = 1;
var x = 2;

console.log(x);
```

As your codebase becomes larger, you might accidentally overwrite a variable that you did not intend to. Because this behavior does not throw an error, searching for and fixing bugs becomes more difficult.

A keyword called `let` was introduced in ES6, a major update to JavaScript, to solve this potential issue with the `var` keyword. So unlike `var`, when you use `let`, a variable with the same name can only be declared once.

### Read-Only Variable with the `const` Keyword
**`const`** has all the awesome features that `let` has, with the added bonus that variables declared using `const` are read-only. They are a constant value, which means that once a variable is assigned with `const`, it cannot be reassigned.

```javascript
const FAVE_PET =  "Dogs";
FAVE_PET = "Cats";         // TypeError: "FAVE_PET" is read-only
```

**NOTE**: It is common for developers to use uppercase variable identifiers for immutable values and lowercase or camelCase for mutable values (objects and arrays). 