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

* **target** is the innermost element in the DOM that triggered the event, while **currentTarget** is the element that the event listener is attached to.


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

#### ReactDOM

```
import ReactDOM from 'react-dom/client';
const root = ReactDOM.createRoot(document.getElementById('root'));
rootElement.render(<Greet />);
```

#### React.Component

[CodePen URL](https://codepen.io/gireeshcse/pen/OJrrKJN)

```
import React from "react";

// Context: Global Variable for our components
// use provider compenent to pass down.
// the parameters here are default values
// To pass down data child components user Consumer
const ThemeContext = React.createContext({
    "fg":"#fff",
    "bg":"#222"
})

function Greet(){
    return (
        <p>
            Have a good day!
        </p>
    )
}

// Context: Consumer
function Header(props){
    return (
        <ThemeContext.Consumer>
        {theme=>(
            <p style={{color:theme.fg,backgroundColor:theme.bg}}>
                {props.message}
            </p>
        )}
        </ThemeContext.Consumer>
    )
}
export default class ClassComponent extends React.Component{

    constructor(props){
        super(props)
        // this.state is private. Updated using only this.setState.
        // good practice to initialize components with 'empty' state
        this.state = {
            "info":"Initial Value",
            "name":props.name === undefined ? "Class Component":props.name,
            products:[],
            theme:{
                "fg":"green",
                "bg":"yellow"
            }
        }
        // In our own custom components methods, we have to manually bind this to the component
        // Otherwise, this.state will not work.
        this.addVote = this.addVote.bind(this);
        this.changeTheme = this.changeTheme.bind(this);

    }
    componentDidMount(){
        //called immediately after a component is mounted. Setting state here will trigger re-rendering.
        this.setState({info:"component mounted"})
        setTimeout(()=>{
            this.setState({info:"component re-rendered"})
            this.forceUpdateInterval = setInterval(()=>{
                console.log("re-Render the component")
                this.forceUpdate() // to render the component
            },100)
        },2000)
    }

    componentWillUnmount(){
        clearInterval(this.forceUpdateInterval)
    }

    // since this is arrow function, this is automatically binded
    handleClick = (evt) => {
        // this.setState is asynchronous it is scheduled by react
        this.setState(state => {
            let i = state.products.length;
            let limit = i+ 1;
            let newProducts = [];
            while(i<=limit){
                newProducts.push({
                    id:i,
                    name:"Product "+i,
                    "votes":0
                })
                i += 1;
            }

            return {products:[...state.products,...newProducts]}
        })
    }
    addVote(evt){
        let productId = evt.currentTarget.dataset.id;
        if(productId !== undefined){
            this.setState((prevState,props)=>{
                //map returns new array
                let newProducts = prevState.products.map(product=>{
                    if(product.id === parseInt(productId)){
                        return Object.assign({},product,{votes:product.votes+1})
                    }
                    return product
                })
                //sort based on votes
                //a,b are subsequent elements
                // result(b - a)
                //result < 0 : a comes first
                // result > 0: b comes first
                newProducts = newProducts.sort((a,b)=>(b.votes - a.votes))
                return {products:newProducts}
            })
        }
    }

    removeProduct = (evt) => {
        let productId = evt.currentTarget.dataset.id;
        if(productId !== undefined){
            this.setState((prevState)=>{
                let newProducts = prevState.products.filter(product=>{
                    if(product.id === parseInt(productId)){
                        return false
                    }
                    return true
                })
                return {products:newProducts}
            })
        }
    }

    changeTheme(){
        this.setState({theme:{
            fg:"red",
            bg:"blue"
        }})
    }

    // this is bound to the component. React binds this to the component for us.
    render(){
        
        // this.props is immutable - 
        // child can read its props, it can't modify. Parent owns it
        let eleFooter = React.createElement("div",{className:"footer"},
            React.createElement("p",{},"Footer Info")
        )

        const eleItems = this.state.products.map(ele=>{
            return (
                <div key={ele.id}>{ele.name} Votes - {ele.votes} 
                    <button data-id={ele.id} onClick={this.addVote}>Add Vote</button>
                    <button data-id={ele.id} onClick={this.removeProduct}>Remove</button>
                </div>
            )
        })

        return (
        <div>
            <ThemeContext.Provider value={this.state.theme}>
                <Header message="React App"/>
                <button onClick={this.changeTheme}>Change Theme</button>
                <div id={this.props.id}>
                    <Greet />
                </div>
                <p>This is a {this.state.name} - {this.state.info}</p>

                <Header message="Products"/><button onClick={this.handleClick}>Add Product</button><hr />
                {eleItems}
                {eleFooter}
            </ThemeContext.Provider>
            <Header message="React App"/>
        </div>
        );
    }
}
```

#### Single Responsibility Principle

- Only responsible for one piece of functionality
- Improves reusability

- Example
    - TimersDashboard
        - EditableTimerList
            - EditableTimer
                - Timer
                - TimerForm
            - EditableTimer
                - Timer
                - TimerForm
            - EditableTimer
                - Timer
                - TimerForm
        - Stats

#### Developing a React App

<ol>
<li>Break the app into components</li>
<li>Build a static version of the app</li>
<li>Determine what should be stateful.</li>
<li>Determine in which component each piece of state should live</li>
<li>Hard-code initial states</li>
<li>Add inverse data flow</li>
<li>Add server communication</li>
</ol>
            
* Bottom-level components
    - Leaf components - contains majority of HTML
* Top Level Components
    - Concerned about the orchestraction.

* **Not a State**
    - Passed in from a parent via props. If it doesn't change over a time.
    - Can we compute it based on any other state or props in our component.

* **For each piece of state**
    - Identify every component that renders something based on that state.
    - Find a common owner component (In Hierarchy)
    - Either the common owner or another component higher up in the hierarchy should own the state.
    - If we can't find a component where it makes sense to own the state, create a new component simply for holding the state and add it some where in the hierarchy above the common owner component.

* Props are states immutable accomplice.
    - State is managed in some select parent components and then that data flows down through children as props.

* React will compare the result of the render to the previous call to render() and if it sees that nothing has changed, it stops there - it won't attempt any DOM manipulation.

* Data flows from top-down through the component tree to leaf components.Leaf components communicate events to state managers by calling prop functions.

* React uses a Virtual DOM
    - Tree of JS objects that represent actual DOM
    - Actual DOM modification leads to
        - Its hard to keep track of changes.
        - it can be slow (modifying DOM on every change can cause significantly degrade the performance)

    - Virtual DOM
        - Use efficient diffing algorithms, in order to know what changed.
        - update subtrees of the DOM simulataneously.
        - Batch updates to the DOM.

    - React's Virtual DOM
        - Tree of React Elements.
    
    - Shadow DOM
        - Ability of the browser to include subtree of DOM elements into rendering of a document, but not into the main document DOM tree.
        - Form of encapsulation on our elements.
        - Video Controls for video tag

    - ReactElement is representation of DOM element in the virtual DOM.
        - Stateless and immutable
        ```
        var boldElement = React.createElement('b');
        ```
* Template literals (Template Strings) - Javascript feature

    ```
    const user = {
        'name': 'ram',
        'isAdmin':true
    }
    const age = 30;

    function userNotification(strings,userExp,ageExp){
        const str0 = strings[0]; // "You "
        const str1 = strings[1]; // " are a "
        const str2 = strings[2]; // "."

        const role = userExp.isAdmin ? " user with admin access": "standard user";

        return `${str0}(${userExp.name})${str1}${role}${str2} whose age is ${ageExp}`;
    }

    const output = myTag`You ${user} are a ${age}.`;

    console.log(output);
    // You(ram) are a user with admin access whose age is 30.

    function recursive(strings, ...values) {
        console.log(strings, values);
        return recursive;
    }
    recursive`Hello``World`;
    // [ 'Hello' ] []
    // [ 'World' ] []
    ```

* Formatter - JS

    ```
    const formatter = new Intl.NumberFormat('en-US',{
        style:'currency',
        currency:'USD',
        minimumFactionDigits:2,
        maximumFactionDigits:2
    });

    fomatter.format(savingsAmount)
    fomatter.format(interestPercentage)
    ```

* **JSX Javascript Extension Syntax**

    - HTML tags starts with lowercase
    - ReactComponents starts with uppercase
    - Transformed into JS by using a pre-processor build tool before we load it with the browser.
    - Attribute Expressions - wrap it in {}
        ```
        const level = 'debug';
        const component = (
            <Alert color={level === 'debug'? 'grey':'red'}> log={true}
        )
        ```
    - Conditional Child Expressions

        ```
        const renderAdminMenu = function(){
            return (
                <MenuLink to="/users">User Accounts</MenuLink>
            )
        }

        const userLevel = this.props.userLevel;
        return(
            <ul>
                <li> Menu</li>
                {userLevel === 'admin' && renderAdminMenu()}
            </ul>
        )

        const Menu = (
            <li>
            {loggedInUser ? <UserMenu />: <LoginLink />}
            </li>
        )
        ```
    - Boolean Attributes
        ```
        <input name="name" disabled={true}>
        ```
    - JSX Comments
        ```
        {/*

        */}
        ```
    - Spread Syntax
        ```
        const props = {msg:"hello",name:"world"}
        <Component {...props}>
        ```
    - Others
        ```
        (
            <div className={classNames.join(' ')}>
                <label className={`form-control ${!isValid ? 'invalid':''}`} htmlFor="name">Name</label>
                {filteredData.length === 0 && <p>No data found.</p>}
                {
                    filteredData.length > 0 && 
                    filteredData.map((item)=>(
                        <CustomComponent key={item.id} title={item.title}>
                    ))
                }
            </div>
        )

        const [userInput,setUserInput] = useState({
            "name":"",
            "age":""
        })

        const inputChangeHandler = (input, value) =>{
            setUserInput((prevInput)=>{
                return {
                    ...prevInput,
                    [input]: value
                }
            })
        }

        <input type="text" value={user["name"]} onClick={(evt)=>{
            inputChangeHandler("name",evt.target.value)
        }} />
        ```
    - **styled-components**
        ```
        import styled from 'styled-components';
        const Button = styled.button`
            color:white;
            width:100%;
            &:focus{
                outline:none
            }
            &:hover{
                cursor:pointer
            }

            @media (min-width:768px){
                width:auto
            }
        `;
        <Button>Click</Button>
        const FormControl = styled.div`

            margin: 0.5px;
            &.invalid input{

            }

            &.invalid label{

            }

            & input:focus{

            }
        `;

        const AdvFormControl = styled.div`
            & input{
                border: 1px solid ${props=> (props.valid ? "#ccc":"red")}
            }
        `;

        <AdvFormControl valid={true}>
            <label>Name</label>
            <input type="text" />
        </AdvFormControl>
        ```

    - CSS Modules
        - Works directly for apps created using create-react-app. otherwise configuration is required

    ```
    File: Button.module.css

    .button{

    }

    .button:focus{
        
    }

    .button:hover,.button:active{
        
    }

    .form-control{

    }

    .form-control.invalid{

    }

    File: Button.js
    import React from 'react';
    import styles from './Button.module.css'

    const Button = props => {
        return (
            <button type={props.type} className={styles.button} onClick={props.onClick}>{props.children}</button>
        )
    }

    File:Other.js

    <div className={`styles['form-control'] ${!isValid} && styles.invalid}`}>
    
    <div>
    ```
        

* this.state and this.props are updated asynchronously.
* **The only info we should ever put in state are values that are not computed and do not need to be synced across the app.**
* Minimize the number of components with state. Red flag in our application, if our state gets large or unmanageable.

* Stateless Component
    - React light weight way of building components that only need the render.
    - Funtional components can have performance benefits.
    - Try to pull our state to the parent components.
    - Encourages the reuse.

* Fragments,Portals and Refs

    - Each component should return only one jsx element which will result in <div> soup.
        ```
        <div>
            <div>
                <div>
                    <h2>Some content</h2>
                </div>
            </div>
        </div>
        ```

        In bigger apps, we can easily end up with tons of unnecessary divs (or other elements) which add no semantic meaning or structure to the page but are only there because of React's JSX requirement.

        sol:
        ```
        const Wrapper = (props) => {
            return props.children;
        }
        export default Wrapper;

        return (<Wrapper>
            <Wrapper>
                <h2>Some content</h2>
            </Wrapper>
        </Wrapper>)
        ```

        React Solution:
        ```
        return (
            <React.Fragment>
                <h2>Some content</h2>
                <h2>Some content</h2>
            </React.Fragment>
        )

        //or 
        return (
            <>
                <h2>Some content</h2>
                <h2>Some content</h2>
            </>
        )
        ```

    - React Portals

        - This is used for overlay elements such modals.
        ```
        return (
            <>
                <MyModal>Some content</MyModal>
                <h2>Some content</h2>
            </>
        )
        ```

        - Portals need the locations
        
        ```
        File: index.html

        <body>

            <div id="backdrop-root"></div>
            <div id="modal-root"></div>
            <div id="root"></div>
        </body>
        ```

        - Component logic

        ```
        import ReactDOM from 'react-dom';

        <>
            {ReactDOM.createPortal(<MyModal onClick={props.onClick}>Some content</MyModal>,document.getElementById('modal-root'))}
        </>
        ```

    - Refs

        - Helps to get access to other DOM elements.
        
        ```
        import {useRef} from 'react';

        const nameInputRef = useRef(); // real dom element/node is connected here
        // don't modify the element using this only read properties

        <input type="text" id="name" ref={nameInputRef}> // uncontrolled components

        const submitHandler = (evt) => {
            console.log(nameInputRef.current.value)
        }

        ```

* Effects, Reducers and Contexts

    - Effect or Side Effect
        - Store data in browser storage
        - Send Http requests to backend servers
        - Set and manage timers/intervals
    
    - useEffect() hook
        - `useEffect(()=>{},[dependencies])`
            - Here the function will be executed after every component evaluation if the specified dependencies changed.
            
        - Without useEffect the following code creates loops

        ```
        funtion App(){
            const [isLoggedIn,setIsLoggedIn] = useState(false);

            const storedUserLoggedInInformation = localstorage.getItem('isLoggedIn');
            if(storedUserLoggedInInformation === '1'){
                setIsLoggedIn(true);
            }

            const loginHandler = (username,password) => {
                // API CALL

                localstorage.setItem('isLoggedIn','1');
            }
        }
        ```

        - Modified version

        ```
        import {useEffect} from 'react';

        funtion App(){
            const [isLoggedIn,setIsLoggedIn] = useState(false);

            useEffect(()=>{
                const storedUserLoggedInInformation = localstorage.getItem('isLoggedIn');
                if(storedUserLoggedInInformation === '1'){
                    setIsLoggedIn(true);
                }
            },[]) // runs only once when app starts up

            const loginHandler = (username,password) => {
                // API CALL

                localstorage.setItem('isLoggedIn','1');
            }
        }
        ```
        - Runs only if email or password is changed.
        ```
        useEffect(()=>{
            setFormIsValid(
                email.includes('@') && password.trim().length > 6
            )
        },[email,password])
        ```

        - Clean up function

        ```

         useEffect(()=>{
            const formValidator = setTimeout(()=>{
                setFormIsValid(
                    email.includes('@') && password.trim().length > 6
                )
            },500) // called after 500ms

            // called before every useEffect call except initial call
            // also called before unmounting of component
            return ()=>{
                clearTimeout(formValidator);
            }
            
        },[email,password])
        // if email/password is changed 15 times continously, cleanup is called 15 times but timeout function is called once. 

         useEffect(()=>{
            console.log("RUNS everytime")
         })

        ```
    - `useReducer()`
        - Sometimes, we have more complex state - for example if it got multiple states, multiple ways of changing it or dependencies to other states. useState() then often becomes hard or error prone to use. its easy to write bad, inefficient or buggy code in such schenarios. 
        - useReducer() can be used as a replacement for useState() if we need 'more powerful state management.'

        ```
        const [state,dispatchFn] = useReducer(reducerFn,initialState,initFn);
        
        // reducerFn =(prevState,action) => newState
        ```

        ```

        import {useReducer} from 'react';
        const userReducer = (state,action) => {
            if(action.type === 'USER_INPUT"){
                let val = {};
                if(action.payload.fieldName === "email"){
                    val["email"] = action.payload.val
                    val["vaildEmail"] = action.payload.val.includes('@')
                }else{
                    val["password"] = action.payload.val
                    val["validPassword"] = action.payload.val.includes('@')
                }
                return {
                    ...state,...val
                }
            }

            return state;

        }

        const initialState = {
            "email":"",
            "password":"",
            "vaildEmail":false,
            "validPassword":false
        }

        function Login(){
            const [userState,dispatchUser] = useReducer(userReducer,initialState);
        
            const [formValid,setFormValid] = useState(false);

            useEffect(()=>{
                setFormValid(userState.validEmail && userState.validPassword)
            },[userState.validEmail,userState.validPassword])

            const emailChangeHandler = (evt) => {
                dispathchUser({
                   "type":"USER_INPUT",
                    payload:{
                        "val":evt.current.value,
                        "fieldName":"username"
                    }
                })
            }
        }

        ```

        - useState()
            - Main state management tool
            - Great for independent piecies of state/data
            - Great if state updates are easy and limited to a few kinds of updates
        - useReducer()
            - Should be considered if we have related piecies of state/data
            - Can be helpful if we have more complex state updates


    - Context

    ```
    //auth-context.js
    const AuthContext = React.createContext({
        isLoggedIn:false,
        onLogout:()=>{}
    })

    export default AuthContext;


    //JSX Code
    const [isLoggedIn,setIsLoggedIn] = useState(false);

    <AuthContext.Provider value={{isLoggedIn:isLoggedIn,onLogout:logoutHandler}}>
        <Nav />
        <Login />
    <AuthContext.Provider>


    //JSX Code
    <AuthContext.Consumer>
    {(ctx)=>{
        return (
            <div>
            {
                ctx.isLoggedIn ? "User Logged In": "Not Logged In"
            }
            </div>
        )
    }}
    </AuthContext.Consumer>
    ```
    Note: The default value will be used if we use hook. for provider we need to provide the value

    - Using hook

    ```
    const ctx = useContext(AuthContext);
    ```



#### ```<Suspense />```

    - Allows react to "suspend" rendering a component subtree
        - Used when a child component is not ready to be rendered.
            - ECMAScript bundle containing the component isn't loaded yet.
            - The data needed for a component to render isn't available yet.
    
    - The "fallback" component will be rendered instead
        - Replaces the complete children component tree

    - Rendering is suspended when a promise is thrown.
        - And resumed when the promise resolves.
