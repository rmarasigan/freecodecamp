# `match()` Method

You can also extract the actual matches you found with the `.match()` method. The `.match()` method retrieves the result of matching a string against a [regular expression](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions). To use the `.match()` method, apply the method on a string and pass in the regex inside the parentheses. Here's an example:

```javascript
"Hello, World!".match(/Hello/);   // ["Hello"]

let text = "Regular Expressions";
let textRegex = /Expressions/;
text.match(textRegex);            // ["expressions"]
```

Note that the `.match()` syntax is the "opposite" of the `.test` method you have been using thus far:
```javascript
'string'.match(/regex/);
/regex/.test('string');
```

### More than the first match
To search or extract a pattern more than once, you can use the global search flag: **`g`**.
```javascript
let text = "Repeated, Repeated, Repeated";;
let repeatedRegex = /Repeated/;

text.match(repeatedRegex);
```

And here `match` returns the value `["Repeated", "Repeated", "Repeated"]`.

**NOTE**: You can have multiple flags on your regex liek `/search/gi`.

### Match anything with Wildcard Period
Sometimes you won't (or don't need to) know the exact characters in your patterns. Thinking of all words that match, say, a misspelling would take a long time. Luckily, you can save time using the *wildcard character*: **`.`**

The wildcard character `.` will match any one character. The wildcard is also called *`dot`* and *`period`*. You can use the wildcard character just like any other character in the regex. For example, if you wanted to match `hug`, `huh`, `hut`, and `hum`, you can use the regex `/hu./` to match all four words.

```javascript
let hug = "Bear hug";
let hum = "I'll hum a song";
let textRegex = /hu./;

textRegex.test(hug);
textRegex.text(hum);
```

Both of these `test` calls would return `true`.

### Match single character with multiple possibilities
You can search for a literal pattern with some flexibility with *character classes*. **Character classes** allow you to define a group of characters you wish to match by placing them inside square (`[` and `]`) brackets.

For example, you want to match `bag`, `big`, and `bug` but not `bog`. You can create the regex `/b[aiu]g/` to do this. The `[aiu]` is the character class that will only match the characters `a`, `i`, or `u`.

```javascript
let big = "big";
let bag = "bag";
let bug = "bug";
let bog = "bog";

let textRegex = /b[aiu]g/;
big.match(textRegex); // ["big"]
bag.match(textRegex); // ["bag"]
bug.match(textRegex); // ["bug"]
bog.match(textRegex); // null
```

### Letters of the alphabet
Inside a character set, you can define a range of characters to match using a *hyphen character*: `-`. For example, to match lowercase letters *a* through *e* you would use `[a-e]`.

```javascript
let cat = "cat";
let bat = "bat";
let mat = "mat";

let textRegex = /[a-e]at/;
cat.match(textRegex);   // ["cat"]
bat.match(textRegex);   // ["bat"]
mat.match(textRegex);   // null
```

### Numbers and letters of the alphabet
Using the hyphen (`-`) to match a range of characters is not limited to letters. It also works to match a range of numbers. For example, `[0-5]` matches any number between *0* and *5*, including the *0* and *5*. Also, it is possible to combine a range of letters and numbers in a single character set.

```javascript
let john = "John8675309";
let textRegex = /[a-z0-9]/ig;

john.match(textRegex);
```