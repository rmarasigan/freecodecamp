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

### Match beginning string patterns
Outside of a character set, the caret is used to search for patterns at the beginning of strings.

```javascript
let rickyRegex = "/^Ricky/";
let notFirst = "You can't find Ricky now."
let ricky = "Ricky is first and can be found.";

rickyRegex.test(ricky);     // true
rickyRegex.test(notFirst);  // false
```

### Match ending string patterns
There is also a way to search for patterns at the end of strings. You can search the end of strings using the dollar sign character `$` at the end of the regex.

```javascript
let textRegex = /story$/;
let ending = "This is a never ending story";
let noEnding = "Sometimes a story will have to end";

textRegex.test(ending);   // true
textRegex.test(noEnding); // false
```

### Match all letters and numbers
Using character classes, you were able to search for all letters of the alphabet with `[a-z]`. This kind of character class is common enough that there is a shortcut for it, although it includes a few extra characters as well.

The closest character class in JavaScript to match the alphabet is `\w`. This shortcut is equal to `[A-Za-z0-9_]`. This character class matches upper and lowercase letters plus numbers. Note, this character class also includes the underscore character (`_`).

```javascript
let shortHand = /\w+/;
let longHand = /[A-Za-z0-9_]+/;

let numbers = "42";
let names = "important_var";

longHand.test(numbers);   // true
shortHand.test(numbers);  // true

longHand.test(names);   // true
shortHand.test(names);  // true
```

These shortcut character classes are also known as *shorthand character classes*.

### Match everything but letters and numbers
A natural pattern you might want to search for is the opposite of alphanumerics. You can search for the opposite of the `\w` with *`\W`*. Note, the opposite pattern uses a capital letter. This shortcut is the same as `[^A-Za-z0-9_]`.

```javascript
let shortHand = /\W/;

let numbers = "42%";
let phrase = "Coding!";

numbers.match(shortHand); // ["%"]
phrase.match(shortHand);  // ["!"]
```

### Match all numbers
Another common pattern is looking for just digits or numbers. The shortcut to look for digit characters is `\d`, with a lowercase `d`. This is equal to the character class `[0-9]`, which looks for a single character of any number between zero and nine.

### Match all non-numbers
You can also search for non-digits using a similar shortcut that uses an uppercase `D` instead. The shortcut to look for non-digit characters is `\D`. This is equal to the character class `[^0-9]`, which looks for a single character that is not a number between zero and nine.

### Match whitespace
You can also match the whitespace or spaces between letters. You can search for whitespace using `\s`, which is a lowercase **`s`**. This pattern not only matches whitespace, but also carriage return, tab, form feed, and new line characters. You can think of it as similar to the character class `[ \r\t\f\n\v]`.

```javascript
let spaceRegex = /\s/g;
let whitespace = "Whitespace. Whitespace everywhere!";
whitespace.match(spaceRegex); // [" ", " "]
```

### Match non-whitespace characters
Search for non-whitespace using `\S`, which is an uppercase `s`. This pattern will not match whitespace, carriage return, tab, form feed, and new line characters. You can think of it being similar to the character class `[^ \r\t\f\n\v]`.

```javascript
let spaceRegex = /\S/g;
let whitespace = "Whitespace. Whitespace everywhere!";
whitespace.match(spaceRegex).length; // 32
```

### Positive and Negative Lookahead
*Lookaheads* are patterns that tell JavaScript to look-ahead in your string to check for patterns further along. This can be useful when you want to search for multiple patterns over the same string. There are two kinds of lookaheads: positive lookahead and negative lookahead.

A ***positive lookahead*** will look to make sure the element in the search pattern is there, but won't actually match it. A positive lookahead is used as (`?=...`) where the `...` is the required part that is not matched.

Syntax:
```javascript
/match(?=element)/
```

If `element` follows `match` then it will be a match otherwise will technically not be a match and will not be declared as a match. The positive lookahead is a sort of group with parenthesis around it. Within this group the expression starts with a question mark immediately followed by equal sign and then the element to look ahead.

Example:
```
This is a car

Regex: /a(?=r)/
```
The regex is going to match that `a` which is immediately followed by an `r`.

On the other hand, a ***negative lookahead*** will look to make sure the element in the search pattern is not there. A negative lookahead is used as (`?!...`) where the `...` is the pattern that you do not want to be there. The rest of the pattern is returned if the negative lookahead part is not present.

Syntax:
```javascript
/match(?!element)/gm
```

Where `match` is the item to match and `element` is the item which should not immediately follow for a successful match. The match will be declared a match if it is not followed by a given element. Thus this pattern helps in matching those items which have a condition of not being immediately followed by a certain character, group of characters or a regex group.

Example:
```
ad, ae, az, ab

Regex: /a(?!b)/
```
The regex is going to match `ad`, `ae`, and `az` but will not match `ab`.


```javascript
let quit = "qu";
let no_quit = "qt";

let quitRegex = /q(?=u)/;
let noQuitRegex = /q(?!u)/;

quit.match(quRegex);          // ["q"]
no_quit.match(noQuitRegex);   // ["q"]
```

A more practical use of lookaheads is to check two or more patterns in one string. Here is a (naively) simple password checker that looks for between 3 and 6 characters and at least one number:
```javascript
let password = "abc123";
let checkPass = /(?=\w{3,6})(?=\D*\d)/;

checkPass.test(password);
```