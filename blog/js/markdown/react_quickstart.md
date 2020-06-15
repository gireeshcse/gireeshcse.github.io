# React Quick Start

### Routing

    npm install react-router-dom --save

#### Step 1

```
import React from 'react';
import ReactDOM from 'react-dom';

import { BrowserRouter as Router } from "react-router-dom";
import App from './App';

ReactDOM.render(
    <Router>
        <App />
    </Router>,
    document.getElementById('root')
);

```
#### Step 2

```
import { Redirect, Switch, Route } from 'react-router-dom';

...

render(){
    return (
      <Switch>
        <Route path='/login' component={Login} />
        <Route path='/register' component={Register} />
        <Route exact path='/' render={() => (
          <Redirect to='/login' />
        )} />
        
        <Route component={NoMatch} />
      </Switch>
    );
}

```

#### Step 4

```
import { Link } from 'react-router-dom';

....

<Link to="/login" id="login-button" className="btn btn-facebook btn-user btn-block">
    <i className="fas fa-heading"></i> Login
</Link>

```

### Forms

    npm install --save prop-types
    npm install --save validator

