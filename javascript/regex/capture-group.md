# Capture Groupings

Capturing groups are a way to treat multiple characters as a single unit. They are created by placing the characters to be grouped inside a set of parentheses. Capturing groups are numbered by counting their opening parentheses from the left to the right. In the expression ((A)(B(C))), for example, there are four such groups −
* ((A)(B(C)))
* (A)
* (B(C))
* (C)

### Check For Mixed Grouping of Characters
Sometimes we want to check for groups of characters using a Regular Expression and to achieve that we use parentheses `()`. By placing part of a regular expression inside round brackets or parentheses, you can group that part of the regular expression together. This allows you to apply a quantifier to the entire group or to restrict alternation to part of the regex.

Only parentheses can be used for *grouping*. Square brackets define a *character class*, and curly braces are used by a *quantifier with specific limits*. If you want to find either `Penguin` or `Pumpkin` in a string, you can use the following Regular Expression: `/P(engu|umpk)in/g`

```javascript
let text = "Pumpkin";
let textRegex = /P(engu|umpk)in/;

textRegex.test(text);   // true
```

### Reuse Patterns Using Capture Groups
Capture groups are constructed by enclosing the regex pattern to be captured in parentheses. In this case, the goal is to capture a word consisting of alphanumeric characters so the capture group will be `\w+` enclosed by parentheses: `/(\w+)/`.

The substring matched by the group is saved to a temporary "variable", which can be accessed within the same regex using a backslash and the number of the capture group (e.g. \1). Capture groups are automatically numbered by the position of their opening parentheses (left to right), starting at 1.

```javascript
let text = "row row row your boat";
let textRegex = /(\w+) \1 \1/;

textRegex.test(text);   // true
text.match(textRegex);  // ["row row row", "row"]
```

Using the `.match()` method on a string will return an array with the matched substring, along with its captured groups.

### Use Capture Groups to Search and Replace
Searching is useful. However, you can make searching even more powerful when it also changes (or replaces) the text you match. You can search and replace text in a string using `.replace()` on a string. The inputs for `.replace()` is first the regex pattern you want to search for. The second parameter is the string to replace the match or a function to do something.

```javascript
let text = "The sky is silver.";
let textRegex = /silver/;

text.replace(textRegex, "blue");
```

The `.replace(...)` call would return the string: `The sky is blue`. You can also access capture groups in the replacement string with dollar signs (`$`).

```javascript
"Code Camp".replace(/(\w+)\s(\w+)/, '$2 $1');
```
The `.replace(...)` call would return the string `Camp Code`.


### Remove Whitespace from Start and End
Sometimes whitespace characters around strings are not wanted but are there. Typical processing of strings is to remove the whitespace at the start and end of it.

Example:
```javascript
str = str.replace(/\s+/g, '');
```

`\s` is the regex for "whitespace" and `g` is the "global flag", meaning match all `\s` (whitespaces).