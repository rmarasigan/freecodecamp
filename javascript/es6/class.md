# Class

ES6 provides a new syntax to create objects, using the *class* keyword. It should be noted that the `class` syntax is just syntax, and not a full-fledged-class-based implementation of an object-oriented paradigm, unlike in languages such as Java, Python, Ruby, etc.

In ES5, an object can be created by defining a `constructor` function and using the `new` keyword to instantiate the object.

In ES6, a `class` declaration has a `constructor` method that is invoked with the `new` keyword. If the `constructor` method is not explicitly defined, then it is implicitly defined with no arguments. Classes are a template for creating objects. Classes are in fact "special [functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions)", and just as you can define [function expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/function) and [function declarations](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function), the class syntax has two components: [class expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/class) and [class declarations](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/class).

Syntax:
```javascript
// Class Declarations
class class_name {
}

// Class Expressions
let var_name = class {
}
```

Example:
```javascript
// Explicit Constructor
class SpaceShuttle {
   constructor(targetPlanet) {
      this.targetPlanet = targetPlanet;
   }
   takeOff() {
      console.log(`To ${this.targetPlanet}!`);
   }
}

// Implicit Constructor
class Rocket {
   launch() {
      console.log(`To the moon!`);
   }
}

const zeus = new SpaceShuttle('Jupiter');
zeus.takeOff(); // To Jupiter!

const atlas = new Rokcet();
atlas.launch(); // To the moon!
```

It should be noted that the `class` keyword declares a new function, to which a constructor is added. This constructor is invoked when `new` is called to create a new object.

> **Note**: UpperCamelCase should be used by convention for ES6 class names, as in `SpaceShuttle` used above.

The `constructor` method is a special method for creating and initializing an object created with a class.

## Getters and Setters
You can obtain values from an object and set the value of a property within an object. These are classically called *getters* and *setters*. **Getter functions** are meant to simply return (get) the value of an object's private variable to the user without the user directly accessing the private variable. **Setter functions** are meant to modify (set) the value of an object's private variable based on the value passed into the setter function. This change involve calculations, or even overwritting the previous value completely.

Example:
```javascript
class Book {
   constructor(author) {
      this._author = author;
   }

   // Getter
   get writer() {
      return this._author;
   }

   // Setter
   set writer(updatedAuthor) {      
      this._author = updatedAuthor;
   }
}

const novel = new Book('Nicholas Sparks');
console.log(novel.writer);  // Nicholas Sparks

novel.writer = 'Garry Bishop';
console.log(novel.writer);  // Garry Bishop
```

Notice the syntax used to invoke the getter and setter. They do not even look like functions. Getters and setters are important because they hide internal implementation details.

> **Note**: It is convention to precede the name of a private variable with an underscore (_). However, the practice itself does not make a variable private.

## Good Read
* [Classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes)
* [ES6 - Classes](https://www.tutorialspoint.com/es6/es6_classes.htm)
* [JavaScript Class](https://www.javascripttutorial.net/es6/javascript-class/)