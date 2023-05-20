# Logical Errors

### Use of assignment operator instead of equality operator
Branching programs, i.e ones that do different things if certain conditions are met, rely on `if`, `else if`, and `else` statements in JavaScript. The condition sometimes takes the form of testing whether a result is equal to a value. This logic is spoken as "if x equals y, then..." which can literally translate into code using the `=`, or assignment operator. This leads to unexpected control flow in your program.

The assignment operator (`=`) in JavaScript assigns a value to a variable name and the `==` and `===` operator check for equality (the triple `===` tests for strict equality, meaning both value and type are the same).