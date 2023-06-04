# Properties

In JavaScript, an object is a collection of properties, where each property is a key-value pair. JavaScript uses prototypal inheritance. Therefore, a property of an object can be either **own** or **inherited**. A property that is defined directly on an object is **own** while property that the object receives from tis prototype is **inherited**.

Objects can inherit features from one another via **prototype**. Every object has its own property called prototype. Because a prototype itself is also another object, the prototype has its own prototype. This creates a something called ***prototype chain***. The prototype chain ends when a prototype has `null` for its own prototype.

### Understand *own properties*
In the following example, the `Account` constructor defines two properties: `name` and `age`:

```javascript
function Account(name, age) {
  this.name = name;
  this.age = age;
}

let luke = new Account('Luke Chiang', 23);
let benson = new Account('Benson Boone', 20);
```

`name` and `age` are called ***own properties***, because they are defined directly on the instance object. That means that `luke` and `benson` each has its own separate copy of those properties. In fact every instance of the `Account` will have its own copy of these properties. The following code adds all the own properties of `luke` to the array `ownProps`:

```javascript
let ownProps = [];

for (let property in luke) {
  if (luke.hasOwnProperty(property)) {
    ownProps.push(property);
  }
}

console.log(ownProps);  // ["name", "age"]
```

### Use `prototype` properties to reduce duplicate code
The `prototype` is an object that is associated with every functions and objects by default in JavaScript, where function's prototype property is accessible and modifiable and object's prototype property (a.k.a attributes) is not visible. The `prototype` object is special type of enumerable object to which additional properties can be attached to it which will be shared across all the instances of it's constructor function.

```javascript
function Bird(name) {
  this.name = name;
}

let duck = new Bird("Donald");
let canary = new Bird("Tweety");
```

Since `numLegs` will probably have the same value for all instances of `Bird`, you essentially have a duplicated variable `numLegs` inside each `Bird` instance. This may not be an issue when there are only two instances, but imagine if there are millions of instances. That would be a lot of duplicated variables.

A better way is to use the `prototype` of `Bird`. Properties in the `prototype` are shared among ALL instances of `Bird`. Now all instances of `Bird` have the `numLegs` property.

```javascript
Bird.prototype.numLegs = 2;

console.log(duck.numLegs);
console.log(canary.numLegs);
```

Since all instances automatically have the properties on the `prototype`, think of a `prototype` as a "recipe" for creating objects. Note that the `prototype` for `duck` and `canary` is part of the `Bird` constructor as `Bird.prototype`.

The combination of the constructor and prototype patterns is the most common way to define custom types in ES5. In this patterns:
* The constructor pattern defines the object properties.
* The prototype pattern defines the object methods.

By using this pattern, all objects of the custom type share the methods defineed in the prototype. Also, each object has its own properties. This constructor/prototype pattern takes the bast parts of both constructor and prototype patterns.

### Iterate over all properties
**Own properties** are defined directly on the object instance itself and **prototype properties** are defined on the `prototype`.

```javascript
function Dog(name) {
  this.name = name; // own property
}

Dog.prototype.legs = 2; // prototype property

let husky = new Dog("Popoy");
```

Here is how you add `husky`'s own properties to the array `ownProps` and `prototype` properties to the array `prototypeProps`:

```javascript
let ownProps = [], prototypeProps = [];

for (let property in husky)
{
  if (husy.hasOwnProperty(property))
  {
    ownProps.push(property);
  }
  else
  {
    prototypeProps.push(property);
  }
}

console.log(ownProps);          // ["name"]
console.log(prototypeProps);    // ["legs"]
```

### Understand the Constructor Property
There is a special `constructor` property located on the object instances `duck` and `husky`.

```javascript
let duck = new Bird();
let husky = new Dog();

console.log(duck.constructor === Bird); // true
console.log(husky.constructor === Dog); // true
```

Note that the `constructor` property is a reference to the constructor function that created the instance. The advantage of the `constructor` property is that it is possible to check for this property to find out what kind of object it is. Since the `constructor` property can be overwritten, it is generally better to use the `instanceof` method to check the type of an object.

All objects (with the exception of objects created with `Object.create(null)`) will have a `constructor` property. Objects created without the explicit use of a constructor function (such as object- and array-literals) will have a `constructor` property that points to the Fundamental object constructory type for that object.

### Change the `prototype` to a New Object
Up until now you have been adding properties to the `prototype` inidividually:

```javascript
Bird.prototype.numLegs = 2;
```

This become tedious after more than a few properties:

```javascript
Bird.prototype.eat = function () {
  console.log("nom nom nom");
}

Bird.prototype.describe = function () {
  console.log("My name is " + this.name);
}
```

A more efficient way is to set the `prototype` to a new object that already contains the properties. This way, the properties are added all at ones:

```javascript
Bird.prototype = {
  numLegs: 2,
  eath: function () {
    console.log("nom nom nom");
  },
  describe: function () {
    console.log("My name is " + this.name);
  }
};
```

### Remember to set the `constructor` property when changing the prototype
There is one crucial side effect of manually setting the prototype to a new object. It erases the `constructor` property. This property can be used to check which constructor function created the instance, but since the property has been overwritten, it now gives `false` results:

```javascript
duck.constructor === Bird;    // false
duck.constructor === Object;  // true
duck instanceof Bird;         // true
```

To fix this, whenever a prototype is manually set to a new object, remember to define the `constructor` property:

```javascript
Bird.prototype = {
  constructor: Bird,
  numLegs: 2,
  eat: function () {
    console.log("nom nom nom");
  },
  describe: function () {
    console.log ("My name is " + this.name);
  };
}
```

### Understand where an object's prototype comes from
Just like people inherit genes from their parents, an object inherits its `prototype` directly from the constructor function that created it. For example, here the `Bird` constructor creates the `duck` object:

```javascript
function Bird(name) {
  this.name;
}

let duck = new Bird("Donald");
```

`duck` inherits its `prototype` from the `Bird` constructor function. You can show this relationship with the `isPrototypeOf` method:

```javascript
Bird.prototype.isPrototypeOf(duck);
```

### Understand the `prototype` Chain
All objects in JavaScript (with a few exceptions) have `prototype`. Also, an object's `prototype` itself is an object.

```javascript
function Bird(name) {
  this.name = name;
}
```

Because a `prototype` is an object, a `prototype` can have its own `prototype`. In this case, the `prototype` of `Bird.prototype` is `Object.prototype`.

```javascript
Object.prototype.isPrototypeOf(Bird.prototype);
```

When accessing a property-based language like JavaScript, a dynamic lookup takes places that involves different layers within the object's prototype tree. Since JavaScript functions are objects, they can have properties. A particularly important property that each function has is called `prototype`.

`prototype`, which is itself an object, inherits from its parent's prototype, which inherits from its parent's prototype, and so on. This is often referred to as the *prototype chain*. `Object.prototype`, which is always at the end of the prototype chain, contains methods like `toString()`, `hasProperty`, `isPrototypeOf()`, and so on.


```javascript
let duck = new Bird("Donald");
duck.hasOwnProperty("name");
```

The `hasOwnProperty` method is defined in `Object.prototype`, which can be accessed by `Bird.prototype`, which can then be accessed by `duck`. This is an example of the `prototype` chain. In this `prototype` chain, `Bird` is the `supertype` for `duck`, while `duck` is the `subtype`. `Object` is a `supertype` for both `Bird` and `duck`. `Object` is a `supertype` for all objects in JavaScript. Therefore, any object can use the `hasOwnProperty` method.

## Good Read
* [JavaScript Prototype](https://www.javascripttutorial.net/javascript-prototype/)
* [Prototype in JavaScript](https://www.tutorialsteacher.com/javascript/prototype-in-javascript)
* [Object.prototype.constructor](https://udn.realityripple.com/docs/Web/JavaScript/Reference/Global_Objects/Object/constructor)
* [JavaScript Constructor/Prototype Pattern](https://www.javascripttutorial.net/javascript-constructor-prototype/)
* [Understanding Own Properties of an Object in JavaScript](https://www.javascripttutorial.net/javascript-own-properties/)
* [All you need to know to understand JavaScript’s Prototype](https://www.freecodecamp.org/news/all-you-need-to-know-to-understand-javascripts-prototype-a2bff2d28f03)
* [Constructor Functions and Prototype Objects in JavaScript](https://codeahoy.com/learn/html/ch19/)
* [JavaScript Prototype Chains, Scope Chains, and Performance: What You Need to Know](https://www.toptal.com/javascript/javascript-prototypes-scopes-and-performance-what-you-need-to-know)