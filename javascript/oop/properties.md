# Properties

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
* The constructor pattern defines the obejct properties.
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
There is a special `constructor` property located on the obejct instances `duck` and `husky`.

```javascript
let duck = new Bird();
let husky = new Dog();

console.log(duck.constructor === Bird); // true
console.log(husky.constructor === Dog); // true
```

Note that the `constructor` property is a reference to the constructor function that created the instance. The advantage of the `constructor` property is that it is possible to check for this property to find out what kind of object it is. Since the `constructor` property can be overwritten, it is generally better to use the `instanceof` method to check the type of an object.

All objects (wtih the exception of objects created with `Object.create(null)`) will have a `constructor` property. Objects created without the explicit use of a constructor function (such as object- and array-literals) will have a `constructor` property that points to the Fundamental object constructory type for that obejct.

## Good Read
* [Object.prototype.constructor](https://udn.realityripple.com/docs/Web/JavaScript/Reference/Global_Objects/Object/constructor)
* [JavaScript Constructor/Prototype Pattern](https://www.javascripttutorial.net/javascript-constructor-prototype/)
* [Constructor Functions and Prototype Objects in JavaScript](https://codeahoy.com/learn/html/ch19/)