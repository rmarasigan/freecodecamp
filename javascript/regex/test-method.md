# `test()` Method

Regular expressions are used in programming languages to match parts of strings. You create patterns to help you do that matching. If you want to find the word `the` in the string `The dog chased the cat`, you could use the following regular expression: `/the/`. Notice that quote marks are not required within the regular expression.

JavaScript has multiple ways to use regexes. One way to test a regex is using the `.test()` method. The `.test()` method takes the regex, applies it to a string (which is placed inside the parentheses), and returns `true` or `false` if your pattern finds something or not.

```javascript
let title = "freeCodeCamp";
let titleRegex = /Code/;

titleRegex.test(title);
```

The `test` method here returns `true`.

Here's another example searching for a literal match of the string `John`:
```javascript
let greetings = "Hi, my name is John";
let greetingsRegex = /John/;

greetingsRegex.test(greetings);
```

This `test` call will return `true`. Any other forms of `John` will not match. For example, the regex `John` will not match `john` or `JOHN`.

```javascript
let wrongRegex = /john/;
wrongRegex.test(greetings);
```
This `test` call will return `false`.

JavaScript RegExp objects are **stateful** when they have the [`global`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/global) or [`sticky`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/sticky) flags set (e.g., `/foo/g` or `/foo/y`). They store a [`lastIndex`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/lastIndex) from the previous match. Using this internally, `test()` can be used to iterate over multiple matches in a string of text (with capture groups).

### Match a literal string with different possibilities
Using regexes liek `/coding/`, you can look for the pattern `coding` in another string. This is powerful to search single strings, but it's limited to only one pattern. You can search for multiple patterns using the **`alternation`** or **`OR`** operator: **`|`**.

This operator matches patterns either before or after it. For example, if you wanted to match the strings `yes` or `no`, the regex you want is `yes|no`. You can also search for more than just two patterns. You can do this by adding more patterns with more `OR` operators separating them, like `/yes|no|maybe/`.

### Ignore case while matching
Up until now, you've looked at regexes to do literal matches of strings. But sometimes, you might want to also match case differences. Case (or sometimes letter case) is the difference between uppercase letters and lowercase letters. Examples of uppercase are `A`, `B`, and `C`. Examples of lowercase are `a`, `b`, and `c`.

You can match both cases using what is called a ***flag***. There are other flags but here you'll focus on the flag that ignores case - the **`i`** flag. You can use it by appending it to the regex. An example of using this flag is `/ignorecase/i`. This regex can match the strings `ignorecase`, `igNoreCase`, and `IgnoreCase`.

### [Quantifiers](https://learn.microsoft.com/en-us/dotnet/standard/base-types/quantifiers-in-regular-expressions)
Quantifiers specify how many instance of a character, group, or character class must be present in the input for a match to be found.

<table>
  <thead>
    <tr>
      <td> Greedy quantifier </td>
      <td> Lazy quantifier </td>
      <td> Description </td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>?</code></td>
      <td><code>??</code></td>
      <td>Matches 0 or one time</td>
    </tr>
    <tr>
      <td><code>{</code> n <code>}</code></td>
      <td><code>{</code> n <code>}?</code></td>
      <td>Matches exactly <i>n</i> times</td>
    </tr>
    <tr>
      <td> <code>{</code> n <code>,}</code> </td>
      <td> <code>{</code> n <code>,}?</code> </td>
      <td>Matches at least <i>n</i> times</td>
    </tr>
    <tr>
      <td> <code>{</code> n <code>,</code> m <code>}</code> </td>
      <td> <code>{</code> n <code>,</code> m <code>}?</code> </td>
      <td>Matches from <i>n</i> to <i>m</i> times</td>
    </tr>
  </tbody>
</table>

The quantities `n` and `m` are integer constants. Ordinarily, quantifiers are greedy. They cause the regular expression engine to match as many occurrences of particular patterns as possible. Appending the `?` character to a quantifier makes it lazy.

### Specify Upper and Lower Number of matches
You can specify the lower and upper number of patterns with *quantity specifiers*. Quantity specifiers are used with curly brackets (`{}`). You put two numbers between the curly brackets - for the lower and upper number of patterns.

```javascript
let a = "waaaaah";
let multipleARegex = /wa{1,6}h/;
multipleARegex.test(a);   // true
```

### Specify only the Lower Number of matches
You can specify the lower and upper number of patterns with quantity specifiers using curly brackets. Sometimes you only want to specify the lower number of patterns with no upper limit. To only specify the lower number of patterns, keep the first number followed by a comma.

```javascript
let a1 = "waaaaah";
let a2 = "waah";

let multipleARegex = /wa{3,}h/;
multipleARegex.test(a1);   // true
multipleARegex.test(a2);   // false
```

### Specify Exact Number of Matches
You can specify the lower and upper number of patterns with quantity specifiers using curly brackets. Sometimes you only want a specific number of matches. To specify a certain number of patterns, just have that one number between the curly brackets.

```javascript
let a1 = "waaaaah";
let a2 = "waaah";

let multipleARegex = /wa{3}h/;
multipleARegex.test(a1);   // false
multipleARegex.test(a2);   // true
```

### Check for All or None
Sometimes the patterns you want to search for may have parts of it that may or may not exist. However, it may be important to check for them nonetheless. You can specify the possible existence of an element with a question mark, `?`. This checks for zero or one of the preceding element. You can think of this symbol as saying the previous element is optional.

```javascript
let en_american = "color";
let en_british = "colour";
let enRegex = /colou?r/;

enRegex.test(en_american);    // true
enRegex.test(en_british);     // true
```
