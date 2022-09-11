# Strings

```javascript
var variable_name = "some text here";
```

`"some text here"` is called a *string literal*. A string literal, or string, is a series of zero or more characters enclosed in single or double quotes.

## Escaping Literal Quotes
When you are defining a string you must start and end with a single or double quote. In JavaScript, you can escape a quote from considering it as an end of string quote by placing a *backslash (`\`)* in front of the quote. This signals to JavaScript that the following quote is not the end of the string, but should instead appear inside the string.

```javascript
const message = "Joh said \"Peter is learning JavaScript\".";
```

## Single Quotes
*String* values in JavaScript may be written with single or double quotes, as long as you start and end with the same type of quote. Unlike some other programming languages, single and double quotes work the same in JavaScript.

```javascript
const doubleQuote = "This is a string"; 
const singleQuote = 'This is also a string';
```

## Escape Sequences
Quotes are not the only characters that can be escaped inside a string. There are two reasons to use escaping characters:

1. To allow you to use characters you may not otherwise be able to type out, such as a carriage return.
2. To allow you to represent multiple quotes in a string without JavaScript misinterpreting what you mean.

<table>
   <tr>
      <th>Code</th>
      <th>Output</th>
   </tr>
   <tr>
      <td>\'</td>
      <td>Single quote</td>
   </tr>
   <tr>
      <td>\"</td>
      <td>Double quote</td>
   </tr>
   <tr>
      <td>\\</td>
      <td>Backslash</td>
   </tr>
   <tr>
      <td>\n</td>
      <td>New line</td>
   </tr>
   <tr>
      <td>\r</td>
      <td>Carriage return</td>
   </tr>
   <tr>
      <td>\t</td>
      <td>Tab</td>
   </tr>
   <tr>
      <td>\b</td>
      <td>Word boundary</td>
   </tr>
   <tr>
      <td>\f</td>
      <td>Form feed</td>
   </tr>
<table>

## Concatenating Strings
In JavaScript, when the `+` operator is used with a String value, it is called the *concatenation* operator. You can build a new string out of other strings by concatenating them together. Concatenation does not add spaces between concatenated strings, so you'll need to add them yourself.

```javascript
let string = "This is the first. " + "This is the end.";
```

We can also use the `+=` operator to concatenate a string onto the end of an existing string variable. This can be very helpful to break a long string over several lines.

```javascript
let string = "This is the first.";
string += " This is the end.";
```

Sometimes you will need to build a string. By using the concatenation operator (`+`), you can insert one or more variables into a string you're building.

```javascript
let name = "feeCodeCamp";
let string = "Hello, this is " + name + ". How are you doin'?";
```

## Appending Variables
Just as we can build a string over multiple lines out of string literals, we can also append variables to a string using the plus equals (`+=`) operator.

```javascript
const adjective = "awesome!";
const message = "freeCodeCampe is "

message += adjective;
```

## Length of a String
You can find the length of a String value by writing `.length` after the string variable or string literal.

```javascript
console.log("This is a message".length);
```

**NOTE**: The space character is also counted.

## Bracket Notation
**Bracket notation** is a way to get a character at a specific index within a string.  Most modern programming languages, like JavaScript, don't start counting at 1 like humans do. They start at 0. This is referred to as *Zero-based* indexing.

```javascript
const name = "John";
const firstLetter = name[0];  // J
const secondLetter = name[1];  // o
const thirdLetter = name[2];  // h
const fourthLetter = name[name.length - 1];  // n
```

## Immutability
In JavaScript, String values are immutable, which means that they cannot be altered once created.

```javascript
const string = "Rate";
string[0] = "L";
```

It cannot change the value of `string` to `Late`, because the contents of `string` cannot be altered. This does not mean that `string` cannot be changed, just that the individual characters of sa string literal cannot be changed.