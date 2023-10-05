# React

A library to build user interfaces.

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
```
* <Suspense />

    - Allows react to "suspend" rendering a component subtree
        - Used when a child component is not ready to be rendered.
            - ECMAScript bundle containing the component isn't loaded yet.
            - The data needed for a component to render isn't available yet.
    
    - The "fallback" component will be rendered instead
        - Replaces the complete children component tree

    - Rendering is suspended when a promise is thrown.
        - And resumed when the promise resolves.

