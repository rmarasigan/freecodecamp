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

### Match single characters not specified
You could also create a set of characters that you do not want to match. These types of character sets are called *negated character sets*. To create a negated character set, you place a caret character (`^`) after the opening bracket and before the characters you do not want to match.

For example, `/[^aeiou]/gi` matches all the characters that are not vowel. Note that characters like `.`, `!`, `[`, `@`, `/` and white space are matched - the negated vowel character set only excludes the vowel characters.

```javascript
let sample = "3 blind mice";
let textRegex = /[^aeiou0-5]/gi;

sample.match(textRegex);
```

### Match characters that occur one or more times
Sometimes, you need to match a character (or group of characters) that appears one or more times in a row. This means it occurs at least once, and may be repeated. You can use the **`+`** character to check if that is the case. Remember, the character or pattern has to be present consecutively. That is, the character has to repeat one after the other.

For example, `/a+/g` would find one match in `abc` and return `["a"]`. Because of the `+`, it would also find a single match in `aabc` and return `["aa"]`. If it were instead checking the string `abab`, it would find two matches and return `["a", "a"]` because the `a` characters are not in a row - there is a `b` between them. Finally, since there is no `a` in the string `bcd`, it wouldn't find a match.

### Match characters that occur zero or more times
There's also an option that matches characters that occur zero or more times. The character to do this is the asterisk or star: **`*`**.

```javascript
let soccer = "gooooooooal!";
let phrase = "gut feeling";
let moon = "over the moon";

let textRegex = /go*/;

soccer.match(textRegex);  // ["goooooooo"]
phrase.match(textRegex);  // ["g"]
moon.match(textRegex);    // null
```

### Lazy Matching
In regular expressions, a ***greedy*** match finds the longest possible part of a string that fits the regex pattern and returns it as a match. The alternative is called a ***lazy*** match, which finds the smallest possible part of the string that satisfies the regex pattern.

You can apply the regex `/t[a-z]*i/` to the string `"titanic"`. This regex is basically a pattern that starts with `t`, ends with `i`, and has some letters in between. Regular expressions are by default greedy, so the match would return `["titani"]`. It finds the largest sub-string possible to fit the pattern.

However, you can use the `?` character to change it to lazy matching. `"titanic"` matched against the adjusted regex of `/t[a-z]*?i/` returns `["ti"]`.

**NOTE**: Parsing HTML with regular expressions should be avoided, but pattern matching an HTML string with regular expressions is completely fine.