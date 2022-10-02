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
* Getter
    - A property that allows modification or validation of a related field before returning it.

* Setter
    - A property that allows modification or computation of a value before setting to a related field

```
class Speaker {

    private message: string;

    constructor(private name: string) {}

    get Message() {
        if(!this.message.includes(this.name)){
            throw Error("message is missing speaker's name");
        }
        return this.message;
    }
    set Message(val: string) {
        let tmpMessage = val;
        if(!val.includes(this.name)){
            tmpMessage = this.name + " " + val;
        }
        this.message = tmpMessage;
    }
}

const speaker = new Speaker("john");
speaker.Message = "hello";
console.log(speaker.Message);
```
Note: includes - available in es5 and es6

```
tsc --target "ES6" getSet
```
* Static properties and methods
```
class Runner {    

    static lastRunTypeName: string;

    constructor(private typeName: string) {}

    

    run() {        

        Runner.lastRunTypeName = this.typeName;

    }

}

const a = new Runner("a");

const b = new Runner("b");

b.run();

a.run();

// determines the last class instance that has called the run.
console.log(Runner.lastRunTypeName); // a 
```

### Interfaces

* Interface is like contract which specifies types of parameters and return types which enforces the expectations between both the user and creator of the interface. 
* Strict rules about what can come out of and go into a type instance.

```
interface Employee {
    name: string;
    id: number;
    isManager: boolean;
    getUniqueId: () => string;
}
const linda: Employee = {
    name: "linda",
    id: 2,
    isManager: false,
    getUniqueId: (): string => {
        let uniqueId = linda.id + "-" + linda.name;
        if(!linda.isManager) {
            return "emp-" + uniqueId;
        }
        return uniqueId;
    }
}

console.log(linda.getUniqueId());

const pam: Employee = {
    name: "pam",
    id: 1,
    isManager: true,
    getUniqueId: (): string => {
        let uniqueId = pam.id + "-" + pam.name;
        if(pam.isManager) {
            return "mgr-" + uniqueId;
        }
        return uniqueId;
    }
}
console.log(pam.getUniqueId());
```

* The main use of interfaces is to allow for a single structure across objects but to enable different implementations.

### Inheritance

* private
```
class Vehicle {
    constructor(private wheelCount: number) {}
    showNumberOfWheels() {
        console.log(`moved ${this.wheelCount} miles`);
    }
}

class Motorcycle extends Vehicle {
    constructor() {
        super(2);
    }
}

class Automobile extends Vehicle {
    constructor() {
        super(4);
    }
}

const motorCycle = new Motorcycle();
motorCycle.showNumberOfWheels();

const autoMobile = new Automobile();
autoMobile.showNumberOfWheels();
```

* protected
```
class Vehicle {
    constructor(protected wheelCount: number) {}
    showNumberOfWheels() {
        console.log(`moved ${this.wheelCount} miles`);
    }
}

class Motorcycle extends Vehicle {
    constructor() {
        super(2);
    }
}

class Automobile extends Vehicle {
    constructor() {
        super(4);
    }
    updateWheelCount(newWheelCount: number){
        // can set because it is protected variable 
        this.wheelCount = newWheelCount;
    }
}

const motorCycle = new Motorcycle();
motorCycle.showNumberOfWheels();

const autoMobile = new Automobile();
autoMobile.showNumberOfWheels();
```

* **namespaces**
    - Creates scoping containers and separates one set of code from another.

```
namespace A {
    class FirstClass {}
}

namespace B {
    class SecondClass {}
    // creates error since it is not available to this namespace
    const test = new FirstClass();
}
```

* **Abstract classes**

* An abstract class can act both as a regular class, providing member implementations, and as an interface, providing only the rules for implementation for a child class.
* We cannot instantiate an abstract class.

```
namespace AbstractNamespace {
    abstract class Vehicle {
        constructor(protected wheelCount: number) {}

        abstract updateWheelCount(newWheelCount: number): void;

        showNumberOfWheels() {
            console.log(`moved ${this.wheelCount} miles`);
        }
    }

    class Automobile extends Vehicle {
        constructor() {
            super(4);
        }
        updateWheelCount(newWheelCount: number){
            // can set because it is protected variable 
            this.wheelCount = newWheelCount;
        }
    }
}

```

* **Interface**

```
namespace InterfaceNamespace {

    interface Thing {
        name: string;
        getFullName: () => string;
    }

    interface Vehicle extends Thing {
        wheelCount: number;
        updateWheelCount: (newWheelCount: number) => void;
        showNumberOfWheels: () => void;
    }

    class Motorcycle implements Vehicle {
            name: string;
            wheelCount: number;
            constructor(name: string) {
                // no super for interfaces
                this.name = name;
            }

            updateWheelCount(newWheelCount: number){
                this.wheelCount = newWheelCount;
                console.log(`Automobile has ${this.wheelCount}`);
            }

            showNumberOfWheels() {
                console.log(`moved Automobile ${this.wheelCount} miles`);
            }

            getFullName() {
                return "MC-" + this.name;
            }

        }

        const moto = new Motorcycle("beginner-cycle");
        console.log(moto.getFullName());

    }
}
```

### Generics

```
function getLength<T>(arg: T): number {
    if(arg.hasOwnProperty("length")) {
        return arg["length"];
    }
    return 0;
}

console.log(getLength<number>(22));
console.log(getLength("Hello world."));
```

### Optional Chaining

```
namespace OptionalChainingNS {

    interface Wheels {
        count?: number;
    }

    interface Vehicle {
        wheels?: Wheels;
    }

    class Automobile implements Vehicle {
        constructor(public wheels?: Wheels) {}
    }

    const car: Automobile | null = new Automobile({
        count: undefined
    });

    console.log("car ", car);
    console.log("wheels ", car?.wheels);
    console.log("count ", car?.wheels?.count);
}

const count = !car ? 0
    : !car.wheels ? 0
    : !car.wheels.count ? 0
    : car.wheels.count;
```

### Typescript configuration

```
tsc tsfile.ts –lib 'es5, dom'
```

Above command ignores the tsconfig.json file

* tsconfig.json
    - --lib: This is used to indicate which JavaScript version you will use during development.
    - --target: This indicates which version of JavaScript you want to emit out into .js files.
    - --noImplicitAny: Does not allow the any type, without explicitly declaring it.
    - --outDir: This is the directory where JavaScript files will be saved to.
    - --outFile: This is the final JavaScript filename.
    - --rootDirs: This is an array that stores the .ts file source code.
    - --exclude: This is an array of folders and files to exclude from compilation.
    - --include: This is an array of folders and files to include in compilation.

### Others

```
function doSomething(a, ...others) {

    console.log(a, others, others[others.length - 1]);

}

doSomething(1,2,3,4,5,6,7);
// 1 [2,3,4,5,6,7] 7
````

```
const items = [
    { name: 'jon', age: 20 },
    { name: 'linda', age: 22 },
    { name: 'jon', age: 40}
];

const jon = items.find((item) => {
    return item.name === 'jon'
});
console.log(jon);
// { name: 'jon', age: 20 }
```

```
const widgets = [
    { id: 1, color: 'blue' },
    { id: 2, color: 'yellow' },
    { id: 3, color: 'orange' },
    { id: 4, color: 'blue' },
]

console.log('some are blue', widgets.some(item => {
    return item.color === 'blue';
}));
// some are blue true

console.log('every one is blue', widgets.every(item => {
    return item.color === 'blue';
}));
// every one is blue false
```

```
const userIds = [
    1,2,1,3
];
const uniqueIds = new Set(userIds);
console.log(uniqueIds);
uniqueIds.add(10);
console.log('has', uniqueIds.has(3)); // true
console.log('size', uniqueIds.size); // 4
```

* Map
    - Collection of key-value pairs

```
const mappedEmp = new Map();
mappedEmp.set('linda', { fullName: 'Linda Johnson', id: 1 });
mappedEmp.set('jim', { fullName: 'Jim Thomson', id: 2 });
mappedEmp.set('pam', { fullName: 'Pam Dryer', id: 4 });

console.log(mappedEmp);
console.log('get', mappedEmp.get('jim'));
console.log('size', mappedEmp.size); // 3

for(let [key, val] of mappedEmp) {
    console.log('iterate', key, val);
}
```

### async and await

```
setTimeout(() => {
    console.log('I waited and am done now.');
}, 3000);
```

* A **Promise** is an object with a delayed completion at some indeterminate future time.

```
const myPromise = new Promise((resolve, reject) => {
    setTimeout(() => {
        //resolve('I completed successfully');
        reject('I failed');
    }, 500);
});

myPromise
.then(done => {
    console.log(done);
})
.catch(err => {
    console.log(err);
});
```

```
async function delayedResult() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve('I completed successfully');
        }, 500);
    });
}
(async function execAsyncFunc() {
    const result = await delayedResult();
    console.log(result);
})();
```

```
// natively not available in nodejs
const fetch = require('node-fetch');

(async function getData() {

    const response = await fetch('https://pokeapi.co/api/v2/pokemon/ditto/');

    if(response.ok) {
        const result = await response.json();
        console.log(result);
    } else {
        console.log('Failed to get anything');
    }

})();
```

### React

```
npx create-react-app try-react --template typescript
``