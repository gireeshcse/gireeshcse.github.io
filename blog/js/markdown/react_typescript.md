### React with Typescript

* **Transpilation**
    - Is a method where code from one language is "compiled" or converted into another language.

* **Static Typing**
    - Type of vaiable is to be defined upfront.

* **Four major principles of OOP**
    - Encapsulation
        - Information hiding
        - Containers/Class - protects the data so nothing outside of container can modify or view it
    - Abstraction
        - Hides internal implementation of how data is managed and provide a more simplified interface to the outside code.
    - Inheritence
         - Code reuse
    - Polymorphism

#### Types

* any or unknown 
    - any type cause the compiler to no longer check the type.

```
let val: any = 22;

val = "string value";

val = new Array();

val.push(33);

console.log(val);

// the following will not cause error
val.doesnotexist(33);
```

```
let val: unknown = 22;

val = "string value";

val = new Array();

val.push(33); // gives error

console.log(val);

// the following code has to be used
if (val instanceof Array) {

    val.push(33);

}
```
* & merges two into one. (Intersection Type)
```
let obj: { name: string } & { age: number } = {

    name: 'tom',

    age: 25

}
```

* | or union type

```
let unionObj: null | { name: string } = null;

unionObj = { name: 'jon'};

console.log(unionObj);
```

* Literal types

    - Similar to union types, but they use a set of hardcoded string or number values.

    ```
    let literal: "tom" | "ram" | "sita" | "vedavathi" = "ram";
    literal = "sita";
    literal = "sam"; // compiler error 
    ```

* Type aliases

```
type Points = 20 | 30 | 40 | 50;
let score: Points = 20;

// object literal type declaration.
type ComplexPersion = {
    name: string,
    age: number,
    birthday: Date,
    married: boolean,
    address: string
}
```

#### Functions
* If funtion does not return anything then value it assigned to variable will be undefined
```
function runMore(distance: number): number {

    return distance + 10;

}
```

* Funtions as types

```
type Run = (miles: number) => boolean;

let runner: Run = function (miles: number): boolean {

    if(miles > 10){

        return true;

    }

    return false;

}

console.log(runner(9));

function oldEnough(age: number): never | boolean {

    if(age > 59) {

        throw Error("Too old!");

    }

    if(age <=18){

        return false;

    }

    return true;

}
```

#### Classes

```
class Person {

    constructor() {}

    msg: string;

    speak() {

        console.log(this.msg);

    }

}

const tom = new Person();

tom.msg = "hello";

tom.speak();

// access modifier
class Person {

    constructor(private msg: string) {}

    

    speak() {

        console.log(this.msg);

    }

}

const tom = new Person("hello");

// tom.msg = "hello";

tom.speak();
```

* **readonly** causes a filed to become read-only after it has set one time in the constructor

```
class Person {

    constructor(private readonly msg: string) {}

    

    speak () {

        this.msg = "speak " + this.msg;

        console.log(this.msg);

    }

}

const tom = new Person("hello");

// tom.msg = "hello";

tom.speak();
```
