# Object

### Create a Basic JavaScript Object
Think about things people see every day, like cars, shops, and birds. These are all objects: tangible things people can observe and interact with. What are some qualities of these objects? A car has wheels. Shops sell items. Birds have wings.

These qualities, or ***properties***, define what makes up an object. Note that similar objects share the same properties, but may have different values for those properties. For example, all cars have wheels, but not all cars have the same number of wheels.

Objects in JavaScript are used to model real-world objects, giving them properties and behavior just like their real-world counterparts.

Example:

```javascript
let user = {
  name: 'John Doe',
  username: 'j.doe',
  password: 'j.doe123',
  login: function(username, password) {
    if (username === this.username && password === this.password) {
      console.log('Login successfully');
    } else {
      console.log('Authentication failed!');
    }
  }
}
```

In the above example, we are creating one user object having some properties and actions that can be performed.

### Use dont notation (`.`) to access the properties of an object

Dot notation is used on the object name, `user`, followed by the name of the property, `login`, to perform the `login` action.

```javascript
user.login('john', 'john1234');   // Authentication failed
user.login('j.doe', 'j.doe123');  // Login successfully
```

### Create a method on an object
Objects can have a special type of property, called a *method*. Methods are properties that are functions. This adds different behavior to an object.

Example:

```javascript
let user = {
  name: 'John Doe',
  username: 'j.doe',
  password: 'j.doe123',
  login: function(username, password) {
    if (username === this.username && password === this.password) {
      console.log('Login successfully');
    } else {
      console.log('Authentication failed!');
    }
  },
  sayName: function() {
    return `Name: ${user.name}`;
  }
}

console.log(user.sayName());  // Name: John Doe
```

The example adds the `sayName` method, which is a function that returns the `name` of the user.

### Make code more reusable with the `this` keyword

```javascript
sayName: function() {
  return `Name: ${user.name}`;
}
```

While this is a valid way to access the object's property, there is a pitfall here. If the variable name changes, any code referencing the original name would need to be updated as well. In a short object definition, it isn't a problem, but if an object has many references to its properties there is a greater chance for error.

A way to avoid these issues is with the `this` keyword:

```javascript
sayName: function() {
  return `Name: ${this.name}`;
}
```

`this` refers to the object that the method is associated with: `user`. If the object's name is changed to `account`, it is not necessary to find all the references to `user` in the code. It makes the code reusable and easier to read.