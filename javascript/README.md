![javascript-header](../assets/img/javascript-header.png)

### What is JavaScript
**JavaScript** was initially created to "make web pages alive." The programs in this language are called *scripts*. They can be written right in a web page's HTML and run automatically as the page loads. Scripts are provided and executed as plain text. They don't need special preparation or compilation to run. In this aspect, JavaScript is very different from another language called **Java**.

When JavaScript was created, it initially had another name: "LiveScript". But Java was very popular at that time, so it was decided that positioning a new language as a "younger brother" of Java would help. But as it evovled, JavaScript became a fully independent language with its own specification called [**ECMAScript**](http://en.wikipedia.org/wiki/ECMAScript), and now it has no relation to Java at all.

## Table of Contents

### Basic JavaScript

**JavaScript** is a scripting language you can use to make web pages interactive. It is one of the core technologies of the web, along with HTML and CSS, and is supported by all modern browsers.

* [Comment](basics/comment-variables.md#comment)
* [Variables](basics/comment-variables.md#variables)
* [`var`, `let`, `const`](basics/var-let-const.md)
* [Strings](basics/string.md)
* [Numbers](basics/numbers.md)
* [Arrays](basics/array.md)
* [Functions](basics/functions.md)
* [Boolean](basics/boolean.md)
* [`if` — `else if` — `else` Statements](basics/if-else-if.md)
* [`switch` Statements](basics/swtich.md)
* [Objects](basics/objects.md)
* [Loops](basics/loops.md)
* [Recursion](basics/recursion.md)
* [Generate Random Number](basics/generate-random.md)
* [`parseInt()`](basics/parseInt.md)
* [Conditional Operator](basics/conditional-operator.md)

### ES6

**ECMAScript**, or ES, is a standardized version of JavaScript. Because all major browsers follow this specification, the terms ECMAScript and JavaScript are interchangeable.

ECMAScript is a scripting language specification on which JavaScript is based. [Ecma International](https://www.ecma-international.org/) is in charge of standardizing ECMAScript.

* [`var` and `let` Keywords](es6/var-let.md)
* [`const` Keyword](es6/const.md)
* [Prevent Object Mutation](es6/object-mutation.md)
* [Arrow Functions](es6/arrow-functions.md)
* [Default and Rest Parameters and Spread Operator](es6/functions.md)
* [Destructuring Assignment](es6/destructuring-assignment.md)
* [Template Literals](es6/template-literals.md)
* [Object Literals](es6/object-literals.md)
* [Class](es6/class.md)
* [Modules](es6/modules.md)
* [Promise](es6/promise.md)

### Regular Expressions
**Regular expressions**, often shortened to "regex" or "regexp", are patterns that help programmers match, search, and replace text. Regular expressions are very powerful, but can be hard to read because they use special characters to make more complex, flexible matches.

Regular expressions are patterns used to match character combinations in strings. In JavaScript, regular expressions are also objects.These patterns are used with the [`exec()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/exec), and [`test()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/test) methods of [`RegExp`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp), and with the [`match()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/match), [`matchAll()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/matchAll), [`replace()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replace), [`replaceAll()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replaceAll), [`search()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/search), and [`split()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/split) methods of [`String`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String).

* [`test()` Method](regex/test-method.md)
* [`match()` Method](regex/match-method.md)
* [Capture Groupings](regex/capture-group.md)

### Debugging
**Debugging** is the process of going through your code, finding any issues, and fixing them. Issues in code generally come in three forms: *syntax errors* that prevent your program from running, *runtime errors* where your code has unexpected behavior, or *logical errors* where your code doesn't do what you intended. All modern browsers and most other environments support debugging tools - a special UI in developer tools that makes debugging much easier. It also allows to trace the code step by step to see what exactly is going on.

* [`console`](debugging/console.md)
* [`typeof`](debugging/typeof.md)
* [Syntax Errors](debugging/syntax-errors.md)
* [Runtime Errors](debugging/runtime-errors.md)
* [Logical Errors](debugging/logical-errors.md)

### Basic Data Structures
Data can be stored and accessed in many ways. In computer science, a data structure is **a format to organize, manage and store data** in a way that allows **efficient access and modification**. More precisely, a data structure is a **collection of data values**, the **relationships** among them,and the functions or **operations** that can be applied to that data.

JavaScript has **primitive (built-in)** and **non-primitive (not built-in)** data structures. **Primitive data structures** come by default with the programming language and you can implement them out of the box (like arrays and objects). **Non-primitive data structures** don't come by default and you have to code them up if you want to use them. Different data structures exist because some of them are better suited for certain kind of operations.

* [Array](data-structures/array.md)
* [Objects](data-structures/objects.md)

## Good Read
* [ECMA-6](https://www.ecma-international.org/publications-and-standards/standards/ecma-6/)
* [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
* [ES6 Tutorial](https://www.javascripttutorial.net/es6/)
* [Regular Expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions)
* [The Modern JavaScript Tutorial](https://javascript.info/)
* [JavaScript Algorithms and Data Structures](https://www.freecodecamp.org/learn/javascript-algorithms-and-data-structures/)
* [Data Structures in JavaScript – With Code Examples](https://www.freecodecamp.org/news/data-structures-in-javascript-with-examples/)