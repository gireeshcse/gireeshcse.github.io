# React

A JS library to build user interfaces. React makes building complex, interactive and reactive user interfaces simpler.

* Just using the JS
    - It quickly becomes 
        - cumbersome
        - error-prone
        - hard to maintain or edit

* Declarative UI Programming
    - We define the target UI state(s) - not the steps to get there.

* Javascript - Imperative Programming
    - We define step-by-step process for a program's execution.

* React Projects use a Build Process
    - Unprocessed React code won't execute in the browser (JSX code)
    - To optimise the code for the production (minified)

### **Javascript**

```
// File: util.js
export let apiKey = 'ddddhjdfhjkhfj';
export let data = 'ddddhjdfhjkhfj';
export default 'ddddhjdfhjkhfj'; // only one default

// File: app.js
// import {apiKey} from './util.js'
// import apiKey from './util.js'
// import {apiKey as KeyCode } from './util.js'
import * as util from './util.js'

console.log(util.apiKey)
console.log(util.default)

// File: index.html

<script src="app.js" type="module"></script>

```

Different types of values
- Numbers
- Strings
- boolean - true,false

* Variables are data containers.
    - Reusable
    - Readability


```
class User{
    constructor(name,age){
        this.name = name;
        this.age = age;
    }
    greet(){
        console.log("hi");
    }
}

const user1 = new User("Ram",35);
user1.greet();

const hobbies = ["sports","cricket","cooking"];
hobbies.push("reading");
const index = hobbies.findIndex((item)=>{
    return item === "sports"
}); // 0 is returned to index

const index = hobbies.findIndex((item)=> item === "sports" );
const editedHobbies = hobbies.map((item)=> item+"!" );

const newEditedHobbies = hobbies.map((item)=> ({text:item}) );

const userNameData = ["Ram","Chandra"];
const [firstName,lastName] = userNameData; // destructuring
const user = {
    name:"ram",
    age:34
}
const {name:userName, age} = user; // object destructuring
console.log(userName);
const newHobbies = ["writing"];
const mergeHobbies = [...hobbies,...newHobbies];
const extendedUser = {
    isAdmin:true,
    ...user
}

const name = prompt("Your name");
for(const hobby of mergeHobbies){
    console.log(hobby);
}

setTimeout(()=>{
    console.log("Timed out!");
},100);

setTimeout(handleTimeout,1000);
```
* Objects - Reference - Address is stored in the variable.
    ```
    const hobbies = ["writing"];
    hobbies.push("reading");
    ```
* Strings - Primitive

### React Components

All user interfaces in react are made up of components.

- Advantages
    * Reusability
    * Separation of concerns.

- Declarative Approach
    - Define the desired target state(s) and let react figure out the actual javascript DOM instructions.

- Only one root element for the component.

#### Composition

- children props

```
function Card(props){
    return (
        <div className="card "+props.className>
            {props.children}
        </div>
    );
}

function Other(props){
    return (<Card className="active">
        <div>{props.title}</div>
    </Card>)
}
```
#### JSX

* Normal code

```
return (
    <div>
        <h2>Lets get started</h2>
        <Expenses items={expenses} />
    </div>
)
```

* Actual Implementation - Underhood code.

```
import React from 'react';

return React.createElement(
    "div",
    {},
    React.createElement("h2",{},"Lets get started"),
    React.createElement(Expenses,{items:expenses})
)

```
This is the one of the reason we need only one root element for wrapping.

#### Why states are needed

```
function NewComponent(props){
    let title = props.title;

    const changeHandler = (evt) => {
        title = "New title";
        // this will not render the new title.
        // even after the button is clicked
    }

    return (
        <div>
            <div>{title}</div>
            <button onClick={changeHandler}>Change Title</button>
        </div>
    )
}
```

```
import React, {useState} from 'react';

function NewComponent(props){
    const [title,setTitle] = useState(props.title);
    // initial value is considered only once under the hood.
    // it will be not reinitialize with the default value for each rendering.

    const changeHandler = (evt) => {
        let title = "New title";
        setTitle(title); // it schedules the change title. 
        // it will not immediately change the data
        // this will render the new title.
    }

    return (
        <div>
            <div>{title}</div>
            <button onClick={changeHandler}>Change Title</button>
        </div>
    )
}

```

#### Updating states

```
const [userInput,setUserInput] = useState({
    title:"",
    date:"",
    amount:"",
})

const titleChangeHandler = (evt) => {
    setUserInput({
        ...userInput,
        title:evt.target.value
    })
}

const amountChangeHandler = (evt) => {
    setUserInput({
        ...userInput,
        amount:evt.target.value
    })
}
```

If value depends on the previous state. This is needed because react schedules the state updates.

```
setUserInput((prevState)=>{

    return {...prevState,title:evt.target.value}
})
```
#### Lifting State Up

* Child-to-Parent Component communication (Bottom-Up)
* Passing state as props to  components


* <Suspense />

    - Allows react to "suspend" rendering a component subtree
        - Used when a child component is not ready to be rendered.
            - ECMAScript bundle containing the component isn't loaded yet.
            - The data needed for a component to render isn't available yet.
    
    - The "fallback" component will be rendered instead
        - Replaces the complete children component tree

    - Rendering is suspended when a promise is thrown.
        - And resumed when the promise resolves.

