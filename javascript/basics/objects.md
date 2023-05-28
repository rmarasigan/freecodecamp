# Objects
AN **object** is a data type that can take in collections of key-value pairs. A major difference between an object and other data types such as strings and numbers in JavaScript is that an objects can store different types of data as its values. On the other hand, primitive data types such as numbers and stirngs can store only numbers and strings, respectively, as their values.

The **key**, also known as the property name, is usually a string. If any other data type is used as a property name other than strings, it would be converted into a string.

The most recognizable feature of an object is the brackets which contain the key-value pair.

```javascript
const object = {};
console.log(typeof object); // 'object'
```

**Objects** are similar to `arrays`, except that instead of using indexes to access and modify their data, you access the data in objects through what are called *properties*. Objects are useful for storing data in a structured way, and can represent real world objects.

Example:
```javascript
const cat = {
   "name: "Whiskers",
   "legs": 4,
   "tails": 1,
   "enemies": ["Water", "Dogs"]
};
```

## Accessing object properties with `.` (dot) notation
There are two ways to access the properties of an object: *dot notation* (`.`) and *bracket notation* (`[]`), similar to an array. Dot notation is what you use when you know the name of the property you're trying to access ahead of time.

```javascript
const object = {
   property1: "value_1",
   property2: "value_2"
};

const property1 = object.property1; // value_1
const property2 = object.property2; // value_2
```

## Accessing object properties with `[]` (bracket) notation
The second way to access the properties of an object is *bracket notation* (`[]`). If the property of the object you are trying to access has a space in its name, you will need to use bracket notation. However, you can still use bracket notation on object properties without spaces.

```javascript
const object = {
   "Space Name": "Kirk",
   "More Space": "Spock",
   "NoSpace": "USS Enterprise"
};

object["Space Name"]; // Kirk
object['More Space']; // Spock
object["NoSpace"];    // USS Enterprise
```

## Accessing object properties with variables
Another use of bracket notation on objects is to access a property which is stored as the value of a variable. This can be very useful for iterating through an object's properties or when accessing a lookup table.

```javascript
const dogs = {
   Fido: "Mutt",
   Snoopie: "Beagle",
   Hunter: "Doberman",
};

const myDog = "Hunter";
const breed = dogs[myDog];
console.log(breed); // Doberman
```

## Updating Object Properties
After you've created a JavaScript object, you can update its properties at any time just like you would update any other variable. You can use either dot or bracket notation to update.

```javascript
const dog = {
   name: "Popoy",
   legs: 4,
   tail: 1
};

dog.name = "Basha";
```

## Add new properties
You can add new properties to existing JavaScript objects the same way you would modify them.

```javascript
// Using the dot notation
dog.bark = "bow-wow";

// Using the bracket notation
dog["bark"] = "bow-wow";
```

## Delete properties
We can also delete properties from objects.

Syntax:
```javascript
delete object.property_name;
```

Example:
```javascript
delete dog.bark;
console.log(dog);
```

Result:
```json
{
   "name": "Basha",
   "legs": 4,
   "tail": 1
}
```

## Using objects for lookups
Objects can be thought of as a key/value storage, like a dictionary. If you have tabular data, you can use an object to lookup values rather than a switch statement or an if/else chain. This is most useful when you know that your input data is limited to a certain range.

## Testing objects
Sometimes it is useful to check if the property of a given object exists or not. We can use the `.hasOwnProperty(propname)` method of objects to determine if that object has the given property name. `.hasOwnProperty()` returns `true` or `false` if the property is found or not.

```javascript
dog.hasOwnProperty("name"); // true
dog.hasOwProperty("bark");  // false
```

## Manipulating Complex objects
Sometimes you may want to store data in a flexible *Data Structure*. A JavaScript object is one way to handle flexible data. They allow for arbitrary combinations of *strings*, *numbers*, *booleans*, *arrays*, *functions*, and *objects*.

## Accessing Nested Objects
The sub-properties of objects can be accessed by chaining together the dot or bracket notation.

Example:
```javascript
const storage = {
   "desk": {
      "drawer": "stapler"
   },
   "cabinet": {
      "top drawer": {
         "folder_1": "a file",
         "folder_2": "secrets"
      },
      "bottom drawer": "soda"
   }
};

storage.cabinet["top drawer"].folder_2; // secrets
storage.desk.drawer;                    // stapler
```

## Accessing Nested Arrays
Objects can contain both nested objects and nested arrays. Similar to accessing nested objects, array bracket notation can be chained to access nested arrays.

Example:
```javascript
const pets = [
   {
      type: "cat",
      names: [
         "Meowzer",
         "Fluffy",
         "KitCat"
      ]
   },
   {
      type: "dog",
      names: [
         "Spot",
         "Bowser",
         "Frankie"
      ]
   }
];

pets[0].names[1];   // Fluffy
pets[1].names[0];   // Spot
```

## Great Read
* [Objects in JavaScript – A Beginner's Guide](https://www.freecodecamp.org/news/objects-in-javascript-for-beginners/)
* [What is a JavaScript Object? Key Value Pairs and Dot Notation Explained](https://www.freecodecamp.org/news/what-is-a-javascript-object/)