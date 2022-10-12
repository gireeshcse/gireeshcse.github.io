### React with Typescript

* **Transpilation**
    - Is a method where code from one language is "compiled" or converted into another language.

* **Static Typing**
    - Type of variable is to be defined upfront.

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
type ComplexPerson = {
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

* **readonly** causes a field to become read-only after it has set one time in the constructor

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
```

* Adding routes 

```
npm i react-router-dom

```
* index.tsx
```
import { BrowserRouter as Router } from "react-router-dom";


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <Router>
      <App />
    </Router>
  </React.StrictMode>
);
```
```
import React, { FC } from "react";
const Home: FC = () => {
  return <div>Hello World! Home</div>;
};
export default Home;

//app.tsx
import { Route,Link,Routes } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h3>Sample App</h3>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/about' element={<About />} />
        </Routes>
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
      </header>
    </div>
  );
}
```

* Stages
    - Mounting
        * Simply the instantiation and initialization of a component
        - Methods
            - constructor
            - getDerivedStateFromProps(props, state) - use with caution
            - render
            - componentDidMount
                - If our component rendered state depends on the network call, this is where we can call the network call and update the state
                - Runs only once on the components first load.

    - Updating
        * Rerendering
        - Methods
            - shouldComponentUpdate(nextProps, nextState)
            - getSnapshotBeforeUpdate(prevProps, prevState)
            - componentDidUpdate(prevProps, prevState, snapshot)
                - State needs to change due to prop changes handled here generally
    - Unmounting
        * Removed from the DOM.
        - Methods
            - componentWillUnmount
* Memorization
    - Like caching
    - Checks whether a props value is different from last time and only if it is different then it triggers a state update.
    - React.memo
        - Triggers a re-render only when a child's props change, not when the parent component re-renders

Note: Using small or little local component state as possible is a best practice for React Development.


```
interface GreetingProps {
    name?: string
}

interface GreetingState {
    message:string
}

export default class Greeting extends React.Component<GreetingProps,GreetingState>{
    constructor(props:GreetingProps){
        super(props);
        this.state = {
            message: `Hello from, ${props.name}`
        }
    }
    static getDerivedStateFromProps(props:GreetingProps,state:GreetingState){
        console.log(props,state);
        if(props.name){
            let message = Greeting.getNewMessage(props.name);
            if(message !== state.message){
                const newState = {...state};
                newState.message = message;
                return newState;
            }
        }
        return state;
    }

    static getNewMessage(name: string = "") {
        return `Hello from, ${name}`;
    }

    render(){
        console.log("Rendering Greeting");
        if(!this.props.name){
            return <div>
                No name given
            </div>
        }

        return <div>
            {this.state.message}
        </div>
    }
}

class App extends React.Component{
  state: {name: string}
  constructor(props:any){
    super(props);
    this.state = {
      name:""
    }
    this.onChangeName =  this.onChangeName.bind(this);
  }
  onChangeName(e:React.ChangeEvent<HTMLInputElement>){
    this.setState({
      name:e.target.value
    })
  }
  // onChangeName = (e:React.ChangeEvent<HTMLInputElement>)=>{
  //   this.setState({
  //     name:e.target.value
  //   })
  // }
  render() {
    console.log("rendering app");
    return (
      <div className="App">
        <header className="App-header">
        <input value={this.state.name} onChange={this.onChangeName}/>
        <Greeting name={this.state.name} />
        </header>
      </div>
    )
  }
}
```

```
import React from "react";
interface GreetingProps {
    message: string
}
export default function Greeting(props: GreetingProps) {
    console.log("rendering Greeting")
    return (<div>
            {props.message}
        </div>);    
}
```
#### Hooks

* Hooks ar just Javascript funtions that provide certain capabilities to the component.
    - Creation of State
    - Access to network data
    - anything component needs
* Not component specific and therefore hooks can be used in any component.

* No way of Sharing the logic in our lifecycle event methods to reuse them in other components. Hooks solves this.

* useState
    - Replaces the state and setState
* useEffect
    - Similar to componentDidMount and componentDidUpdate
    - Run before rendering onto the screen happens.
    - Runs after every completed screen rendering.
    - Passing empty array as a parameter, we are forcing it to run only once.
    - Runs asynchronously
* useLayoutEffect
    - This function runs synchronously, allowing us to get certain element values as they currently are on the screen and then do something with them in a synchronous manner. But, of course, this blocks our UI, so we need to only do things that are very fast or the user experience will suffer.
* useCallback
    - Creates an instance of a function once a set of parameters has been changed.
    - Saves the instance in memory otherwise an instance of the function would be recreated on each render.
    - Takes handler funtion as its first parameter and an array of items that may change as its second.If the items don't change, the callback doesn't get new instance.

    ```
    const [startCount, setStartCount] = useState(0);
    const [count, setCount] = useState(0);
    const setCountCallback = useCallback(() => {

         const inc = count + 1 > startCount ? count + 1 :

         Number(count + 1) + startCount;

         setCount(inc);
    }, [count, startCount]);
    ```
    - count is updated each time is changed . if we remove count from the passed in array [startCount]. Then count is updated only once, on the first run, no matter is how many times we call setCountCallBack. even though setCount changes the count value, the count value this function takes only the initial value.So inc will be always computed as 1 only in this case.
* useMemo
    - To save the result of long-running task.(like caching)
    - Only runs if the array of parameters has changed.
    - Similar to useCallback
    - Returns a value that is the result of some heavy computation.
* useReducer
    - Takes two parameters
        - reducer
        - initial state 
    - returns a state object that will be updated by the reducer and dispatcher that receives updated state data, called an action, and passes it to the reducer.
    - reducer
        - acts as a filtering mechanism and determines how action data will be used to update the state.
* useContext
    - a way of having global state data that can be shared across components.
* useRef
    - used to hold any value in its current property.
    - does not trigger a re-render if it changes and values lives as long as the component it was created in lives.
    - To hold a DOM element.

```
import React, { FC, useState, useEffect } from 'react';
const Greeting: FC<GreetingProps> = ({name}:GreetingProps) => {
    const [message,setMessage] = useState("");

    useEffect(() => {
        if(name){
            setMessage(`Hello from, ${name}`)
        }
    },[name])

    return <div>
        {message}
    </div>
}
export default Greeting;
```

```
const reducer = (state: any, action: any) => {

    console.log("enteredNameReducer");
    
    switch(action.type) {
        case "enteredName":
          if(state.enteredName === action.payload) {
            return state;
          }
          
          return { ...state, enteredName: action.payload}
        
        case "message":
          return { ...state, message: `Hello, ${action.payload}` }  
        default:   
            throw new Error("Invalid action type " + action.type);
    }

  }
  
  const initialState = {
    enteredName: "",
    message: "",
  };

function App(){
  const [{message,enteredName},dispatch] = useReducer(reducer,initialState);
  const onChangeName = (e:React.ChangeEvent<HTMLInputElement>) => {
    dispatch({
      "type":"enteredName",
      "payload":e.target.value
    });

    dispatch({
      "type":"message",
      "payload":e.target.value
    });
  }
  console.log("rendering app");
    return (
      <div className="App">
        <header className="App-header">
        <input value={enteredName} onChange={onChangeName}/>
        <Greeting name={message} />
        </header>
      </div>
    )
}
```

```
interface GreetingProps {
    enteredName: string;
    message: string;
    greetingDispatcher: React.Dispatch<{ type: string, payload: string }>;
}

export default function Greeting(props: GreetingProps) {

    console.log("rendering Greeting")
    const onChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {

        props.greetingDispatcher({
            type: "enteredName",
            payload: e.target.value
        });

        props.greetingDispatcher({
            type: "message",
            payload: e.target.value
        });

    }

    return (<div>
            <input value={props.enteredName} onChange={onChangeName} />
            <div>
                {props.message}
            </div>
        </div>);

}

import React, { FC, useEffect, useRef } from 'react';

export interface ListItem {
    id: number;
}

export interface ListItems {
    listItems?: Array<ListItem>;
}

const ListCreator: FC<ListItems> = React.memo(({ listItems }: ListItems) => {

    let renderItems = useRef<Array<JSX.Element> | undefined>();
    useEffect(() => {
        console.log("listItems updated");
        renderItems.current = listItems?.map((item, index) => {
            return <div key={item.id}>
                {item.id}
            </div>;
        });
    }, [listItems]);

    console.log("ListCreator render");
    return (
        <React.Fragment>
            {renderItems.current}
        </React.Fragment>

    );
});

export default ListCreator;


function App() {

  const [{ message, enteredName }, dispatch] = useReducer(reducer, initialState);
  const [startCount, setStartCount] = useState(0);
  const [count, setCount] = useState(0);

  const setCountCallback = useCallback(() => {
    const inc = count + 1 > startCount ? count + 1 :
    Number(count + 1) + startCount;
    setCount(inc);
  }, [count, startCount]);

  const [listItems, setListItems] = useState<Array<ListItem>>();

    useEffect(() => {
      const li = [];
      for(let i = 0; i < count; i++) {
        li.push({ id: i });
      }
      setListItems(li);
    }, [count]);

    const onWelcomeBtnClick = () => {
      setCountCallback();
    }

  const onChangeStartCount = (e:React.ChangeEvent<HTMLInputElement>) => {
    setStartCount(Number(e.target.value));
  }

  return (

    <div className="App">
      <header className="App-header">
        <Greeting message={message}
          enteredName={enteredName}
          greetingDispatcher={dispatch} />

          <div style={{marginTop: '10px'}}>
              <label>Enter a number and we'll increment it</label>
              <br/>
              <input value={startCount}onChange={onChangeStartCount}
              style={{width: '.75rem'}} />&nbsp;
              <label>{count}</label>
              <br/>
              <button onClick={onWelcomeBtnClick}>Increment count</button>
          </div>
          <div>
            <ListCreator listItems={listItems} />
        </div>
      </header>
    </div>
  )

}
```
* Code reuse is much easier because the Hooks are not tied to any specific class

* If we are seeing two renders, it may be because we are running in StrictMode for development purposes. so to debug, we can replace it with Fragment.Frament is only used to wrap a set of JSX elements that don't have a parent wrapping element such as div.

* Hook components and React in general priortize componentization over inheritance.
```
<React.Fragment>
    <App />
</React.Fragment>
```
### Testing


'Failed to initialize watch plugin' when running tests for newly created app 

[Issue Solution](https://github.com/facebook/create-react-app/issues/11792)

```
npm run test
```

Example:

Note: **data-testid**  attibute for elements allows us easily find these elements in our tests.
```
import { render, fireEvent, screen } from '@testing-library/react';

test('renders Enter a number link', () => {
  render(<App />);
  const linkElement = screen.getByText(/Enter a number/i);
  expect(linkElement).toBeInTheDocument();
});

// for assertions such as toHaveValue to get value of input.
import "@testing-library/jest-dom/extend-expect";

describe("Test Component",()=>{
    it("renders withcout crashing", ()=>{
        const { baseElement } = render(<DisplayText />);
        // console.log(baseElement.innerHTML)
        expect(baseElement).toBeInTheDocument();
    });

    it("receives input text",()=>{
        const testuser = "testuser";
        const { getByTestId } = render(<DisplayText />);
        const input = getByTestId("user-input");
        fireEvent.change(input, {target: { value: testuser}});
        expect(input).toBeInTheDocument();
        expect(input).toHaveValue(textvalue);
    });

    it("shows welcome message", ()=>{
        const testuser = "testuser";
        const msg = `Welcome to React testing, ${testuser}`;
        const { getByTestId } = render(<DisplayText />);
        const input = getByTestId("user-input");
        const label = getByTestId("final-msg");
        fireEvent.change(input, {target: { value: testuser}});
        const btn = getByTestId("input-submit");
        fireEvent.click(btn);
        expect(label).toBeInTheDocument();
        expect(label.innerHTML).toBe(msg);
    })

    if("matches snapshot", ()=>{
        const { baseElement } = render(<DisplayText />);
        expect(baseElement).toMatchSnapshot();
    });
});
```

* toMatchSnapshot
    - Firstime, it creates `__snapshot__` at the root of our src folder. 
    - It then adds or updates a file that has the same nae as our test file and ends with the extension **.snap** whic contains the emitted HTML elements of our component.
    - Now we have a snapshot and our first test run has succeedded. Now, we have changed our component file but not our test file.If we save our component and our test reruns and fails showing what is not matching which is our newly added code generated html.
    - Now if we wanted this change, we can press **u** character under the **watch** usage  to update the snapshots created.

#### Mocking

* Simply replacing specific funtionality in our test with default values.

* Useful for mocking n/w calls
    - [https://jsonplaceholder.typicode.com/](https://jsonplaceholder.typicode.com/)

```
// wait to handle asynchronous calls.
import { render, fireEvent, cleanup, wait } from   "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

afterEach(cleanup);
describe("test displaytext",()=>{
    const userFullName = "John Reese";
    const getUserFullnameMock = (username:string): [Promise<string>, jest.Mock<Promise<string>,[string]>] => {
        const promise = new Promise<string>((res,rej)=>{
            res(userFullName);
        });
        const getUserFullname = jest.fn(
            async (username: string) : Promise<string> => {
                return promise;
            }
        );
        return [promise, getUserFullname]
    }
})
```