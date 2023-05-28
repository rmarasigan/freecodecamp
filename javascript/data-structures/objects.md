# Objects
In JavaScript, an **object** is a collection of *key-value pairs*. This data structure is also called **map**, **dictionary** or **hash-table** in other programming languages. A typical JavaScript object looks like this:

```javascript
const obj = {
  key1: "I'm",
  key2: "an",
  key3: "object"
};
```

We use *curly braces* to declare the object. Then declare each key followed by a colon, and the corresponding value. An important thing to mention is thaty *each key has to be unique within the object*. You can't have two keys with the same name.

Objects can store both values and functions. When talking about objects, *values are called properties*, and *functions are called methods*.

```javascript
const obj = {
  key1: "Hello",
  key2: function () { console.log("I'm a property"); }
}
```

To access properties, yo ucan use two different syntaxes, either `object.property` or `object["property"]`. To access methods, we call `object.method()`

```javascript
console.log(obj.key1);      // "Hello!"
console.log(obj["key1"]);   // "Hello!"
obj.key2();                 // "I'm a property"
```

Objects are a good way to group together data that have something in common or are somehow related. Also, thanks to the fact that property names are unique, objects come in handy when we have to separate data based on a unique condition.

### Add key-value pairs to JavaScript Objects
At their most basic, objects are just collections of *key-value* pairs. In other words, they are pieces of data (*values*) mapped to unique identifiers called *properties* (*keys*). Take a look at an example:

```javascript
const tekkenCharacter = {
  human: true,
  player: 'Hwoarang',
  fightingStyle: 'Tae Kwon Doe'
};
```

The above code defines a Tekken video game character object called `tekkenCharacter`. It has three properties, each of which map to a specific value. If you want to add an additional property, such as "origin", it can be done by assigning `origin` to the object:

```javascript
tekkenCharacter.origin = "South Korea"
```

This uses *dot notation*. If you were to observe the `tekkenCharacter` object, it will now include the `origin` property. `Hwoarang` also had distinct orange hair. You can add this property with bracket notation by doing:

```javascript
tekkenCharacter["hair color"] = "dyed orage";
```

*Bracket notation* is required if your property has a space in it or if you want to use a variable to name the property. In the above case, the property is enclosed in quotes to denote it as a string and will be added exactly as shown. Without quotes, it will be evaluated as a variable and the name of the property will be whatever value the variable is. Here's an example with a variable:

```javascript
const eyes = 'eye color';
tekkenCharacter[eyes] = 'brown';
```

The object will look like this:

```javascript
{
  player: 'Hwoarang',
  fightingStyle: 'Tae Kwon Doe',
  human: true,
  origin: 'South Korea',
  'hair color': 'dyed orange',
  'eye color': 'brown'
};
```

### Modify an object nested within an object
Object properties can be nested to an arbitrary depth, and their values can be any type of data supported by JavaScript, including arrays and even other objects. Consider the following:

```javascript
let nestedObject = {
  id: 28802695164,
  date: 'December 31, 2016',
  data: {
    totalUsers: 99,
    online: 80,
    onlineStatus: {
      active: 67,
      away: 13,
      busy: 8
    }
  }
};
```

The `nestedObject` has three properties `id` (value is a number), `date` (value is a string), and `data` (value is an object with its nested structure). While structures can quickly become complex, we can still use the same notaitons to access the information we need. To assign the value `10` to the `busy` property of the nested `onlineStatus` object, we use *dot notation* to reference the property:

```javascript
nestedObject.data.onlineStatus.busy = 10;
```

Nested objects have their own importance that is present inside objects. These obejcts contain properties of objects. The nested objects are utilized to store the object's properties with another object. The dot and square bracket notations are employed to access the properties of objects in JavaScript. Mostly, the object properties are retrieved to use in any piece of code.

### Access property names with bracket notation
For instance, imagine that our `foods` object is being used in a program for a supermarket cash register. We have some function that sets the `selectedFood` and we want to check our `foods` object for the presence of that food. This might look like:

```javascript
let selectedFood = getCurrentFood(scannedItem);
let inventory = foods[selectedFood];
```

This code will evaluate the value stored in the `selectedFood` variable and return the value of that key in the `foods` object, or `undefined` if it is not present. *Bracket notation* is very useful because sometimes obejct properties are not known before runtime or we need to access them in a more dynamic way.

### `delete` keyword to remove object properties
Objects are key-value stores which provide a flexible, intuitive way to structure data and they provide very fast lookup time. Throughout the rest of these challenges, we will describe several common operations you can perform on objects so you can become comfortable applying these useful data structures in your programs.

In JavaScript, the `delete` operator is the only way to remove properties from an object. When you use `delete`, it'll return *`true`* when it removes a property and *`false`* otherwise. It shouldn't be used on predefined JavaScript object properties like `window`, `Math` and `Date` objects as it can crash your application.

If we wanted to remove the `apples` keyword from our `foods` object, we can remove it by using the `delete` keyword:

```javascript
delete foods.apples;
```
#### `delete` operator can't delete a variable
Any property declared with `var` can't be deleted from the global scope or from a function's scope. If you declare a variable without `var`, it can be deleted. The variable declared without the `var` keyword internally stores it as a property of the `window` object.

#### `delete` operator can delete values from an *array*
Since JavaScript *arrays* are objects, elements can be deleted by using `delete`. `delete` will delete the object property, but it will not reindex the array or update its length. This makes it appear as if it's `undefined`. Using the `delete` operator may leave `undefined` holes in the array. Use `pop()`, `shift()` or `splice()` instead.

#### `delete` operator can't delete built-in objects
Deleting built-in objects like `Math`, `Date`, and `window` objects are unsafe, and they can crash your entire application.

#### `delete` operator can delete some non-configurable properties
Object properties, besides a `value`, have three special attributes:

* `writable`: If `true`, the value can be changed, otherwise, it's read-only
* `enumerable`: If `true`, it's listed in loops, otherwise, it's not listed
* `configurable`: If `true`, the property can be deleted or the attributes can be modified, otherwise, it cannot be changed

### Check if an object has a property
What if we just wanted to know if an object has a specific property? JavaScript provides us with two different ways to do this. One uses the `hasOwnProperty()` method and the other uses the `in` keyword. If we have an object `users` with a property of `Alan`, we could check for its presence in either of the following ways:

```javascript
users.hasOwnProperty('Alan');
'Alan' in users;
```

The `hasOwnProperty()` method will check if an object contains a direct property and will return `true` or `false` if it exists or not. *What do we mean by direct property?*

Whenever you create an object in JavaScript, there is a built-in property called a *prototype* and the value is another object. That object will have its own prototype, and this is known as a *prototype chain*. The `hasOwnProperty()` method will only return true for direct properties and not inherited properties from the prototype chain.

Unlike the `hasOwnProperty()` method, the `in` operator will return `true` for both direct and inherited properties that exist in the object.

### Iterate through the keys of an object with a `for...in` statement
The `for...in` loop iterates through properties in the prototype chain. This means that we need to check if the property belongs to the object using the `hasOwnProperty` whenever we loop through an object with the `for...in` loop. This is to make your code safe and to make it always do the right thing which checks that the key is a property of the object itself and not just a prototype.

```javascript
const population = {
  male: 4,
  female: 93,
  others: 10
};

// Iterate through the object
for (const key in population) {
  if (population.hasOwnProperty(key)) {
    console.log(`${key}: ${population[key]}`);
  }
}
```

**NOTE**: Objects do not maintain an ordering to stored keys like arrays do; thus a key's position on an object, or the relative order in which it appears, is irrelevant when referencing or accessing that key.

### Generate an *array* of all object keys with `Object.keys()`
We can also generate an array which contains all the keys stored in an object with the `Object.keys()` method. This method takes an object as the argument and returns an array of stings representing each property in the object. Again, there will be no specific order to the entries in the array.

The **`Object.keys()`** takes the object we want to loop over as an argument and returns an array containing all property names (also known as keys).

## Good Read
* [`delete` operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/delete)
* [How to use Nested Objects in JavaScript](https://linuxhint.com/use-nested-objects-javascript/)
* [5 Things to Know About the JavaScript Delete Operator](https://builtin.com/software-engineering-perspectives/javascript-delete)
* [How to Check if a Property Exists in a JavaScript Object](https://www.freecodecamp.org/news/how-to-check-if-a-property-exists-in-a-javascript-object/)
* [Loop Through an Object in JavaScript – How to Iterate Over an Object in JS](https://www.freecodecamp.org/news/how-to-iterate-over-objects-in-javascript/)