(Reference)[https://reactjs.org/docs/testing-recipes.html]

### Packages

#### Enzyme

Enzyme is a JavaScript Testing utility for React that makes it easier to test our React Components' output. We can also manipulate, traverse, and in some ways simulate runtime given the output.

(Enzyme Jest)[https://github.com/enzymejs/enzyme/blob/HEAD/docs/guides/jest.md]

```
npm i --save-dev enzyme enzyme-adapter-react-16

import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
 
Enzyme.configure({ adapter: new Adapter() });
```

```
import React from 'react';
import { shallow, mount, render } from 'enzyme';

import Foo from '../Foo';

```

#### Jest

JavaScript test runner that lets us access the DOM via **jsdom**.

#### React Testing Library

Set of helpers that let us test React components without relying on their implementation details.

It does not provide a way to **shallowly** render a component without its children, a test runner like **Jest** lets us do this by mocking.

#### Examples

**beforeEach** and **afterEach** blocks always run and isolate the effects of a test to itself.

```
import { unMountComponentAtNode } from "react-dom";

let container = null;

beforeEach(()=>{
    // setup a DOM element as a render target
    container = document.createElement("div");
    document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

```

**act** makes sure all updates such as rendering, user events or data fetching have been processed and applied to the DOM before we make any assertions.

```
act(() => {
  // render components
});
// make assertions
```

**Rendering**

```
//hello.js

import React from 'React';
export default function Hello(props){
    if(props.name)
        return <h1>Hello, {props.name}!</h1>;
    else
        return <span>Hey, stranger</span>;  
}

// hello.test.js
import React from 'react';
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";

import Hello from "./hello";

let container = null;
beforeEach(() => {
  // setup a DOM element as a render target
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

it("renders with or without a name", () => {
  act(() => {
    render(<Hello />, container);
  });
  expect(container.textContent).toBe("Hey, stranger");

  act(() => {
    render(<Hello name="Ram" />, container);
  });
  expect(container.textContent).toBe("Hello, Ram!");

  act(() => {
    render(<Hello name="Sita" />, container);
  });
  expect(container.textContent).toBe("Hello, Sita!");
});

```

**Data Fetching**

```
// user.js
import React, { useState, useEffect } from 'react';

export default function User(props){
    const [user, setUser] = useState(null);

    async function fetchUserData(id){
        const response = await fetch("/" + id);
        setUser(await response.json());
    }

    useEffect(() => {
        fetchUserData(props.id);
    },[props.id]);

    if(!user)
    {
        return "loading...";
    }

    return (
        <details>
        <summary>{user.name}</summary>
        <strong>{user.age}</strong> years old
        <br />
        lives in {user.address}
        </details>
    );
}

// user.test.js

import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import User from "./user";

let container = null;
beforeEach(() => {
  // setup a DOM element as a render target
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

it("renders user data", async () => {
  const fakeUser = {
    name: "Joni Baez",
    age: "32",
    address: "123, Charming Avenue"
  };
  jest.spyOn(global, "fetch").mockImplementation(() =>
    Promise.resolve({
      json: () => Promise.resolve(fakeUser)
    })
  );

  // Use the asynchronous version of act to apply resolved promises
  await act(async () => {
    render(<User id="123" />, container);
  });

  expect(container.querySelector("summary").textContent).toBe(fakeUser.name);
  expect(container.querySelector("strong").textContent).toBe(fakeUser.age);
  expect(container.textContent).toContain(fakeUser.address);

  // remove the mock to ensure tests are completely isolated
  global.fetch.mockRestore();
});
```
