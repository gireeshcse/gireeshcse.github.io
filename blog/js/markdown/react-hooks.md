
# Hooks

(Reference)[https://reactjs.org/docs/hooks-overview.html]

* They lets us use state and other React features withiout writing a class.
* Hooks are now available with the release of v16.8.0.
* Hooks allows us to reuse stateful logic without changing our component hierarchy.
* Hooks let us split one component into smaller functions based on what pieces are related (such as setting up a **subscription** or **fetching data**)
* **useState** returns a pair: the **current state value** and a **function** that lets you **update** it.This function is similar to the **this.state** in a class, except it does not merge the old and new state together.
* The only argument to useState is the initial state.  The initial state argument is only used during the first render.
* Hooks are functions that lets us “hook into” React state and lifecycle features from function components. **Hooks don’t work inside classes** — they let us use React without classes.



```
import React, { useState } from 'react';

function Example() {
  // Declare a new state variable, which we'll call "count"
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}

```

**Decalring multiple state variables**

```
function ExampleWithManyStates() {
  // Declare multiple state variables!
  const [age, setAge] = useState(42);
  const [fruit, setFruit] = useState('banana');
  const [todos, setTodos] = useState([{ text: 'Learn Hooks' }]);
  // ...
}
```

## Effect Hook

* **useEffect**, adds the ability to perform side effects from a function component.  
* It serves the same purpose as **componentDidMount, componentDidUpdate, and componentWillUnmount** in React classes, but unified into a single API. 

```
    import React, { useState, useEffect } from 'react';

    // Similar to componentDidMount and componentDidUpdate:
    useEffect(() => {
        // Update the document title using the browser API
        document.title = `You clicked ${count} times`;
    });

```