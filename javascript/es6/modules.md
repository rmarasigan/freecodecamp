# Modules

A **module** organizes a related set of JavaScript code. A module can contain variables and functions. A module is nothing more than a chunk of JavaScript code written in a file. By default, variables and functions of a module are not available for use. Variables and functions within a module should be exported so that they can be accessed from within other files. Modules in ES6 work only in ***strict mode***. This means variables or functions declared in a module will not be accessible globally.

### Create a Module Script
JavaScript started with a small role to play on an otherwise mostly HTML web. Today, it’s huge, and some websites are built almost entirely with JavaScript. In order to make JavaScript more modular, clean, and maintainable; ES6 introduced a way to easily share code among JavaScript files. This involves exporting parts of a file for use in one or more other files, and importing the parts you need, where you need them. In order to take advantage of this functionality, you need to create a script in your HTML document with a `type` of `module`. 

Example:
```javascript
<script type="module" src="file_name.js"></script>
```

### Export
Imagine a file called `math.js` that contains several functions related to mathematical operations. One of them is stored in a variable, `multiply`, that takes in two numbers and returns their `product`. You want to use this function in several different JavaScript files. In order to share it with these other files, you first need to `export` it.

```javascript
export const multiply = (x, y) => {
   return x * y;
}
```

The `export` keyword can be used to export components in a module. Exports in a module can be classied as follows:

#### Named Exports
Named exports are distinguished by their names. There can be several named exprots in a module. A module can export selected components using the syntax given below

```javascript
// Using multiple export keyword
export component_1
export component_2
...
export component_n

// Using single export keyword
export { component_1, component_2, ..., component_n }
```

#### Default Exports
Usually you will use this syntax if only one value is being exported from a file. It is also used to create a fallback value for a file or module. Modules that need to export only a single value can use default exports. There can be only one default export per module.

Syntax:
```
export default component_name
```

Example:
```javascript
// Named Function
export default function multiply(x, y) {
   return x * y;
}

// Anonymous Function
export default function(x, y) {
   return x * y;
}
```

Since `export default` is used to declare a fallback value for a module or file, you can only have one value be a default export in each module or file. Additionally, you cannot use `export default` with `var`, `let`, or `const`.

### Import
`import` allows you to choose which parts of a file or module to load. 

Example:
```javascript
import { multiply } from './math.js';
```

Here, `import` will find `multiply` in `math.js`, import just that function for you to use, and ignore the rest. The `./` tells the import to look for the `math.js` file in the same folder as the current file. The relative file path (`./`) and file extension (`.js`) are required when using import in this way. A module can have multiple **import statements**.

Suppose you have a file and you wish to import all of its contents into the current file. This can be done with the `import * as` syntax.

Example:
```javascript
import * as math from './math.js';
```

The above `import` statement will create an object called `math`. This is just a variable name, you can name it anything. The object will contain all of the exports from `math.js` in it, so you can access the functions like you would any other object property.

Example:
```javascript
math.divide(6, 24);
math.multiply(2, 3);
```

#### Named Exports
While importing named exports, the names of the corresponding components must match.

Syntax:
```javascript
import { component_1, component_2, ..., component_n } from module_name

// While importing named exports, they can be renamed using the "as" keyword
import { original_component_name as new_component_name }

// All named exports can be imported onto an object by using the asterisk (*) operator
import * as variable_name from module_name;
```

#### Default Exports
Unlike named exports, a default export can be imported with any name.

Syntax:
```javascript
import any_variable_name from module_name
```

To import a default export, you need to use a different `import` syntax.

Example:
```javascript
import multiply from "./math.js";
```

The syntax differs in one key place. The imported value, `multiply`, is not surrounded by curly braces (`{}`). `multiply` here is simply a variable name for whatever the default export of the `math.js` file is. You can use any name here when importing a default.

## Good Read
* [ES6 - Modules](https://www.tutorialspoint.com/es6/es6_modules.htm)
* [Modules, introduction](https://javascript.info/modules-intro)