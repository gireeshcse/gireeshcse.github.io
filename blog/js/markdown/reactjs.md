React is a JavaScript library for building user interfaces.

Building a React app is all about components. An individual React component can be thought of as a UI component in an app.

A component consists of markup, view logic, and often component-specific style is all housed in one place.(Reusable)

React components implement a **render()** method that takes input data and returns what to display.Input data that is passed into the component can be accessed by render() via **this.props**

### Simple Component

```
class HelloMessage extends React.Component {
  render() {
    return (
      <div>
        Hello {this.props.name}
      </div>
    );
  }
}

ReactDOM.render(
  <HelloMessage name="Taylor" />,
  document.getElementById('hello-example')
);

```

### A Stateful Component

 A component can maintain internal state data (accessed via **this.state**). When a component’s state data changes, the rendered markup will be updated by re-invoking **render()**.

 ```
 class Timer extends React.Component {
  constructor(props) {
    super(props);
    this.state = { seconds: 0 };
  }

  tick() {
    this.setState(state => ({
      seconds: state.seconds + 1
    }));
  }

  componentDidMount() {
    this.interval = setInterval(() => this.tick(), 1000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>
        Seconds: {this.state.seconds}
      </div>
    );
  }
}

ReactDOM.render(
  <Timer />,
  document.getElementById('timer-example')
);

```

### Simple Example

```
class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    this.state = { items: [], text: '' };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  render() {
    return (
      <div>
        <h3>TODO</h3>
        <TodoList items={this.state.items} />
        <form onSubmit={this.handleSubmit}>
          <label htmlFor="new-todo">
            What needs to be done?
          </label>
          <input
            id="new-todo"
            onChange={this.handleChange}
            value={this.state.text}
          />
          <button>
            Add #{this.state.items.length + 1}
          </button>
        </form>
      </div>
    );
  }

  handleChange(e) {
    this.setState({ text: e.target.value });
  }

  handleSubmit(e) {
    e.preventDefault();
    if (this.state.text.length === 0) {
      return;
    }
    const newItem = {
      text: this.state.text,
      id: Date.now()
    };
    this.setState(state => ({
      items: state.items.concat(newItem),
      text: ''
    }));
  }
}

class TodoList extends React.Component {
  render() {
    return (
      <ul>
        {this.props.items.map(item => (
          <li key={item.id}>{item.text}</li>
        ))}
      </ul>
    );
  }
}

ReactDOM.render(
  <TodoApp />,
  document.getElementById('todos-example')
);

```

### A Component Using External Plugins

```
class MarkdownEditor extends React.Component {
  constructor(props) {
    super(props);
    this.md = new Remarkable();
    this.handleChange = this.handleChange.bind(this);
    this.state = { value: 'Hello, **world**!' };
  }

  handleChange(e) {
    this.setState({ value: e.target.value });
  }

  getRawMarkup() {
    return { __html: this.md.render(this.state.value) };
  }

  render() {
    return (
      <div className="MarkdownEditor">
        <h3>Input</h3>
        <label htmlFor="markdown-content">
          Enter some markdown
        </label>
        <textarea
          id="markdown-content"
          onChange={this.handleChange}
          defaultValue={this.state.value}
        />
        <h3>Output</h3>
        <div
          className="content"
          dangerouslySetInnerHTML={this.getRawMarkup()}
        />
      </div>
    );
  }
}

ReactDOM.render(
  <MarkdownEditor />,
  document.getElementById('markdown-example')
);

```

### React Component

* **render()** - required
* **props** - input parameters
* **context** - global variable for our components.
* **state** - a way to hold data that is local to a component

#### Creation

```
import React from 'react';
import createReactClass from 'create-react-class';

// React.createClass
const CreateClassApp = createReactClass({
  render: function() {} // required method
});

export default CreateClassApp;

```

```
import React from 'react';

// ES6 class-style
class ComponentApp extends React.Component {
  render() {} // required
}

export default ComponentApp;

```

**PropTypes**

PropTypes are a way to validate the values that are passed in through our props

    import PropTypes from 'prop-types';

    class MapComponent extends React.Component {
      static propTypes = {
        lat: PropTypes.number,
        lng: PropTypes.number,
        zoom: PropTypes.number,
        place: PropTypes.object,
        markers: PropTypes.array
      };

    }

    const MapComponent = createReactClass({
      propTypes: {
      lat: PropTypes.number,
      lng: PropTypes.number
      // ...
      },
    }

Now our component will validate that lat , lng , and zoom are all numbers, while place is an object and marker is an array

**context**

```
import React from 'react';

export const themes = {
  light: {
    foreground: '#222222',
    background: '#e9e9e9'
  },
  dark: {
    foreground: '#fff',
    background: '#222222'
  }
};

export const ThemeContext = React.createContext(themes.dark);

```

### Unit Testing

* Jest - Behavior driven style
* Enzyme - React Component

    const a = { espresso: '60ml' };
    expect(a).toBe({ espresso: '60ml' }) // fail
    expect(1).toBe(1); // pass
    const a = 5;
    expect(a).toBe(5); // pass
    const a = { espresso: '60ml' };
    const b = a;
    expect(a).toBe(b); // pass

    const a = { espresso: '60ml' };
    expect(a).toEqual({ espresso: '60ml' }) // pass

    describe('My test suite', () => {
      it('`true` should be `true`', () => {
        expect(true).toBe(true);
      });
      it('`false` should be `false`', () => {
        expect(false).toBe(false);
      });
    })

npm i --save-dev enzyme react-test-renderer enzyme-adapter-react-16

import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
Enzyme.configure({ adapter: new Adapter() });


### React Router

* Link, Redirect
* Route, Switch