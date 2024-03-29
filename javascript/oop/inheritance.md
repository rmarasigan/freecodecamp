# Inheritance

JavaScript is a **prototype-based language**, meaning object properties and methods can be shared through generalized objects that have the ability to be cloned and extended. When it comes to inheritance, JavaScript only has one structure: objects.

Each object has a private property (referred to as its `[[Prototype]]]`) that maintains a link to another object called its prototype. That prototype object has its own prototype, and so on until an object whose prototype is `null` is reached.

By definition, `null` has no prototype, and acts as the final link this chain of prototypes. This is known as prototypical inheritance and differes from class inheritance. Among popular object-oriented programming languages, JavaScript is relatively unique, as other prominent languages such as PHP, Python, and Java are class-based languages, which instead define classes as blueprints for objects.

### Use inheritance so you don't repeat yourself
There's a principle in programming called *Don't Repeat Yourself (DRY)*. The reason repeated code is a problem is because any change requires fixing code in multiple places. This usually means more work for programmers and more room for errors.

Notice in the example below that the `describe` method is shared by `Bird` and `Dog`:

```javascript
Bird.prototype = {
  constructor: Bird,
  describe: function () {
    console.log ("My name is " + this.name);
  };
};

Dog.prototype = {
  constructor: Dog,
  describe: function () {
    console.log ("My name is " + this.name);
  };
};
```

The `describe` method is repeated in two places. The code can be edited to follow the DRY principle by creating a `supertype` (or parent) called `Animal`:

```javascript
function Animal () {
  Animal.prototype = {
    constructor: Animal,
    describe: function () {
      console.log("My name is ", this.name);
    }
  }
}
```

Since `Animal` includes the `describe` method, you can remove it from `Bird` and `Dog`:

```javascript
Bird.prototype = {
  constructor: Bird
};

Dog.prototype = {
  constructor: Dog
};
```

### Inherit behaviors from a `supertype`
When it comes to inheritance, JavaScript only has one construct: objects. Each object has a private property which holds a link to another object called its **prototype**. That prototype object has a prototype of its own, and so on until an object is reached with `null` as its prototype. By definition, `null` has no prototype, and acts as the final link in this **prototype chain**. It is possible to mutate any member of the prototype chain or even swap out the prototype at runtime, so concepts like static dispatching do not exist in JavaScript.

Now, we are going to reuse the methods of `Animal` inside `Bird` and `Dog` without defining them again. It uses a technique called inheritance. In here, we are going to make an instance of the `supertype` (or parent).  You already know one way to create an instance of `Animal` using the `new` operator.

```javascript
let animal = new Animal();
```

There are some disadvantages when using this syntax for inheritance. Instead , here's an alternative without those disadvantages:

```javascript
let animal = Object.create(Animal.prototype);
```

`Object.create(obj)` creates a new object, and sets `obj` as the new object's `prototype`. Recall that the `prototype` is like the "recipe" for creating an object. By setting the `prototype` of `animal` to be the `prototype` of `Animal`, you are effectively giving the `animal` instance the same "recipe" as any other instance of `Animal`.

```javascript
animal.eat();
animal instanceof Animal; // true
```

### Set the child's prototype to an instance of the parent
In here, we are going to set the `prototype` of the subtype (or child)—in this case, `Bird`—to be an instance of `Animal`.

```javascript
Bird.prototype = Object.create(Animal.prototype);
```

Remember that the `prototype` is like the "recipe" for creating an object. In a way, the recipe for `Bird` now includes all the key "ingredients" from `Animal`.

```javascript
let duck = new Bird("Donald");
duck.eat();
```

`duck` inherits all of `Animal`'s properties, including the `eat` method.

### Reset an inherited constructor property
When an object inherits its `prototype` from another object, it also inherits the supertype's constructor property.

Here's an example:

```javascript
function Bird() {}
Bird.prototype = Object.create(Animal.prototype);

let duck = new Bird();
duck.constructor
```

But `duck` and all instances of `Bird` should show that they were constructed by `Bird` and not `Animal`. To do so, you can manuallys et the constructor property of `Bird` to the `Bird` object.

```javascript
Bird.prototype.constructor = Bird;
duck.constructor
```

### Add methods after inheritance
A constructor function that inherit its `prototype` object from a supertype constructor functions can still have its own methods in addition to inherited methods. For example, `Bird` is a constructor that inherits its `prototype` from `Animal`:

```javascript
function Animal() {}
Animal.prototype.eat = function () {
  console.log("nom nom nom");
}

function Bird() {}
Bird.prototype = Object.create(Animal.prototype);
Bird.prototype.constructor = Bird;
```

In addition to what is inherited from `Animal`, you want to add behavior that is unique to `Bird` objects. Here, `Bird` will get a `fly()` function. Functions are added to `Bird`'s prototype the same way as any constructor function:

```javascript
Bird.prototype.fly = function () {
  console.log("I'm flying!");
}
```

Now instances of `Bird` wil have both `eat()` and `fly()` methods:

```javascript
let duck = new Bird();
duck.eat; // nom nom nom
duck.fly; // I'm flying!
```

### Override Inherited Methods
Now, you learned that an object can inherit its behavior (methods) from another obejct by referencing its `prototype` object:

```javascript
child_object.prototype = Object.create(parent_object.prototype);
```

Then the `child_object` received its own methods by chaining them onto its `prototype`:

```javascript
child_object.prototype.method_name = function () {...};
```

It is possible to override an inherited method. It is done the same way - by adding a method to `child_object.prototype` using the same method name as the one to override. Here's an example of `Bird` overriding the `eat()` method inherited from `Animal`:

```javascript
function Animal() { }
Animal.prototype.eat = function() {
  return "nom nom nom";
};

function Bird() { }
Bird.prototype = Object.create(Animal.prototype);

Bird.prototype.eat = function() {
  return "peck peck peck";
};
```

If you have an instance `let duck = new Bird();` and you call `duck.eat()`, this is how JavaScript looks for the method on the `prototype` chain of `duck`:

1. `duck` => Is `eat()` defined here? No.
2. `Bird` => Is `eat()` defined here? => Yes. Execute it and stop searching.
3. `Animal` => `eat()` is also defined, but JavaScript stopped searching before reaching this level.
4. Object => JavaScript stopped searching before reaching this level.

## Good Read
* [Inheritance and the prototype chain](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Inheritance_and_the_prototype_chain)
* [JavaScript Prototypes and Inheritance – and Why They Say Everything in JS is an Object](https://www.freecodecamp.org/news/prototypes-and-inheritance-in-javascript/)