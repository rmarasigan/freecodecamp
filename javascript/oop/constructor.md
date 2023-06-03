# Constructor

Constructors are one of the fundamental concepts in Object-Oriented Programing. A constructor is a function you can use to create an instance of an object. As well as creating a new object, a constructor specifies the properties and behaviors that will belong to it.

### Define a Constructor function
*Constructors* are functions that create new objects. Technically speaking, a constructor function is a regular function with the following convention:
* The name of a constructor function starts with a capital letter like `User`, `Document`, etc.
* A constructor function should be called only with the `new` operator.

They define properties and behaviors that will belong to the new object. Think of them as a blueprint for the creation of new objects. Here is an example of a constructor:

```javascript
function Account() {
  this.firstName = 'John';
  this.lastName = 'Doe';
  this.age = 25;
}
```

This constructor defines an `Account` object with properties `firstName`, `lastName`, and  `age`. Constructors follow a few conventions:
* Constructors are defined with a capitalized name to distinguish them from other functions that are not `constructors`.
* Constructors use the keyword `this` to set properties of the object they will create. Inside the constructor, `this` refers to the new object it will create.
* Constructors define properties and behaviors instead of returning a value as other functions might.

### Use a Constructor to create objects

```javascript
function Account() {
  this.firstName = 'John';
  this.lastName = 'Doe';
  this.age = 25;
}

let user = new Account();
```

**NOTE**: `this` inside the constructor always refers to the object being created.

Notice that the `new` operator is used when calling a constructor. This tells JavaScript to create a new instance of `Acount` called `user`. Without the `new` operator, `this` inside the constructor would not point to the newly created object, giving unexpected results. Now `user` has all the properties defined inside the `Account` constructor:

```javascript
user.firstName;
user.lastName;
user.age;
```

### Extend Constructors to receive arguments
What if you want users with different values for `firstName`, `lastName`, and `age`? It's possible to change the properties of each bird manually but that would be a lot of work:

```javascript
let userJohn = new Account();
userJohn.firstName = 'Luke';
userJohn.lastName = 'Chiang';
userJohn.age = 23
```

Suppose you were writing a program to keep track of hundreds or event thousands of different users in an e-commerce app. It would take a lot of time to create all the user accounts, then change the properties to different values for every one. To more easily create differen `Account` objects, you can design your `Account` constructor to accept parameters:

```javascript
function Account(firstName, lastName, age) {
  this.firstName = firstName;
  this.lastName = lastName;
  this.age = age;
}
```

Then pass in the values as arguments to define each unique users into the `Account` constructor:

```javascript
let userBenson = new Account('Benson', 'Boone', 20);
```

This gives a new instance of `Account` with `firstName`, `lastName`, and `age` properties set to `Benson`, `Boone` and `20`, respectively. The `userBenson` has these properties:

```javascript
userBenson.firstName;
userBenson.lastName;
userBenson.age;
```

The constructor is more flexible. It's now possible to define the properties for each `Account` at the time it is created, which is one way that JavaScript constructors are so usefule. They group objects together based on shared characteristics and behavior and define a blueprint that automates their creation.

### Verify an objects Constructor with `instanceof`
Anytime a constructor function creates a new object, that object is said to be an *instance of* its constructor. JavaScript gives a convenient way to verify this with the `instanceof` operator. `instanceof` allows you to compare an obejct to a constructor, returning `true` or `false` based on whether or not that object was created with the constructor.

```javascript
let Account = function (name, age) {
  this.name = name;
  this.age = age;
}

let john = new Accout('John Doe', 25);
john instanceof Account;
```

This `instanceof` method would return `true`.

If an object is created without using a constructor, `instanceof` will verify that it is not an instance of that constructor:

```javascript
let benson = {
  name: 'Benson Boone',
  age: 20
};

benson instanceof Account
```

This `instanceof` method would return `false`.

## Good Read
* [JavaScript Constructor Function](https://www.javascripttutorial.net/javascript-constructor-function/)
* [Understanding Constructors in JavaScript](https://www.makeuseof.com/javascript-constructors-understanding/)