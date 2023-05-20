# Syntax Errors

### Misspelled variable and function names
Transposed, missing, or mis-capitalzied characters in a variable or function name will have the browser looking for an object that doesn't exist - and complain in the form of a reference error. JavaScript variable and function name are case-sensitive.

### Unclosed parentheses, brackets, braces, and quotes
Another syntax error to be aware of is that all opening parentheses, brackets, curly braces, and quotes have a closing pair. Forgetting a piece tends to happen when you're editing existing code and inserting items with one of the pair types. Also, take care when nesting code blocks into othes, such as adding a callback function as an argument to a method.

One way to avoid this mistake is as soon as the opening character is typed, immediately include the closing match, then move the cursor back between them and continue coding. Fortunately, most modern code editors generate the second half of the pair automatically.

### Mixed usage of single and double quotes
JavaScript allows the use of both single (`'`) and double (`"`) quotes to declare a string. Deciding which one to use generally comes down to personal preference, with some expectations. Having two choices is great when a string has contractions or another piece of text that's in quotes. Just be careful that you don't close the string too early, which causes a syntax error.