# Flask

* Micro-framework
* Designed as a extensible framework
* It provides a solid core with the basic services, while extensions provide the rest
* Flask has three main dependencies. 
    - Routing, 
    - Debugging,
    - Web Server Gateway Interface (WSGI)

### Creating Virtual Environment

    sudo apt-get install python3-venv
    python3 -m venv virtual-environment-name

    python3 -m venv demo
    source demo/bin/activate
    pip install flask
    pip freeze

        click==7.1.2
        Flask==1.1.2
        itsdangerous==1.1.0
        Jinja2==2.11.2
        MarkupSafe==1.1.1
        Werkzeug==1.0.1

    pip freeze > requirements.txt

    pip install -r requirements.txt

Decorators are a standard feature of the Python language. A common use of decorators is to register functions as handler functions to be invoked when certain events occur.

    @app.route('/')
    def index():
        return '<h1>Hello World!</h1>'

**add_url_rule()** arguments

* the URL, 
* the endpoint name, 
* and the view function

```
def sample():
    return '<h1>Hello World! Sample</h1>'

app.add_url_rule('/sample', 'sample', sample)

@app.route('/user/<name>')
def user(name):
    return '<h1>Hello, {}!</h1>'.format(name)
```

**Running the app**

```      
export FLASK_APP=app.py
flask run
```
the below is mostlty used in unit testing

```
if __name__ == '__main__':
    app.run()
```

### Debug Mode

* reloader
* debugger (web-based tool)

By default, debug mode is disabled.
To enable it, set a **FLASK_DEBUG=1** environment variable before invoking flask run

To allow web server listen for connections on the public network interface, enabling other computers in the same network to connect as well

    flask run --host 0.0.0.0

### The Request-Response Cycle

* Flask uses contexts to temporarily make certain objects globally accessible
* **request object** 

```
from flask import request
@app.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    return '<p>Your browser is {}</p>'.format(user_agent)
```

Contexts enable Flask to make certain variables globally accessible to a thread without interfering with the other threads.

* application context

    - current_app

        The application instance for the active application.

            from flask import current_app
            current_app.name
    - g

        An object that the application can use for temporary storage during the handling of a request. This variable is reset with each request.

* request context

    - request

        The request object, which encapsulates the contents of an HTTP request sent by the client.
        contains all the information that the client included in the HTTP request
    - session

        The user session, a dictionary that the application can use to store values that are “remembered” between requests.

Flask activates (or pushes) the application and request contexts before dispatching a request to the application, and removes them after the request is handled

#### Request Object

* form
    - A dictionary with all the form fields submitted with the request.
* args 
    - A dictionary with all the arguments passed in the query string of the URL.
* values 
    - A dictionary that combines the values in form and args .
* cookies 
    - A dictionary with all the cookies included in the request.
* headers 
    - A dictionary with all the HTTP headers included in the request.
* files 
    - A dictionary with all the file uploads included with the request.
* get_data()
    - Returns the buffered data from the request body.
* get_json() 
    - Returns a Python dictionary with the parsed JSON included in the body of the request.
* blueprint 
    - The name of the Flask blueprint that is handling the request. 
* endpoint 
    - The name of the Flask endpoint that is handling the request. Flask uses the name of the view function as the endpoint name for a route.
* method 
    - The HTTP request method, such as GET or POST .
* scheme 
    - The URL scheme ( http or https ).
* is_secure() 
    - Returns True if the request came through a secure (HTTPS) connection.
* host 
    - The host defined in the request, including the port number if given by the client.
* path 
    - The path portion of the URL.
* query_string 
    - The query string portion of the URL, as a raw binary value.
* full_path 
    - The path and query string portions of the URL.
* url 
    - The complete URL requested by the client.
* base_url 
    - Same as url , but without the query string component.
* remote_addr 
    - The IP address of the client.
* environ 
    - The raw WSGI environment dictionary for the request.

#### Request Hooks

* before_request

    Registers a function to run before each request.
* before_first_request

    Registers a function to run only before the first request is handled. This can be a convenient way to add server initialization tasks.

* after_request

    Registers a function to run after each request, but only if no unhandled exceptions occurred.
* teardown_request

    Registers a function to run after each request, even if unhandled exceptions occurred.

A common pattern to share data between request hook functions and view functions is to use the **g** context global as storage.


#### Responses

```
@app.route('/')
def index():
    return '<h1>Bad Request</h1>', 400
```

**third argument**, a dictionary of headers that are added to the HTTP response.

```
from flask import make_response
@app.route('/')
def index():
    response = make_response('<h1>This document carries a cookie!</h1>')
    response.set_cookie('answer', '42')
    return response
```

* status_code 

    The numeric HTTP status code
* headers 

    A dictionary-like object with all the headers that will be sent with the response
* set_cookie() 

    Adds a cookie to the response
* delete_cookie() 

    Removes a cookie
* content_length 

    The length of the response body
* content_type 

    The media type of the response body
* set_data() 

    Sets the response body as a string or bytes value
* get_data() 

    Gets the response body

```
from flask import redirect
@app.route('/')
def index():
    return redirect('http://www.example.com')
```

```
from flask import abort
@app.route('/user/<id>')
def get_user(id):
    user = load_user(id)
    if not user:
        abort(404)
    return '<h1>Hello, {}</h1>'.format(user.name)
```

### Templates

Mixing business and presentation logic leads to code that is hard to understand and maintain.

A template is a file that contains the text of a response, with placeholder variables for the dynamic parts that will be known only in the context of a request. The process that replaces the variables with actual values and returns a final response string is called rendering. For the task of rendering templates, Flask uses a powerful **template engine** called **Jinja2**.

```
<h1>Hello, {{ name }}!</h1>
```

By default Flask looks for templates in a **templates** subdirectory located inside the main application directory.

```
from flask import Flask, render_template
app = Flask(__name__)

@app.route('/home')
def home():
    return render_template("index.html")

@app.route('/home/user/<name>')
def homeUser(name):
    return render_template("user.html",name=name)
```

Variables can be modified with filters, which are added after the variable name with a pipe character as separator.

```
<p>Hello {{ name | capitalize }}!</p>
```

#### Filters

* safe 
    - Renders the value without applying escaping
* capitalize 
    -  Converts the first character of the value to uppercase and the rest to lowercase
* lower  
    - Converts the value to lowercase characters
* upper  
    - Converts the value to uppercase characters
* title  
    - Capitalizes each word in the value
* trim  
    - Removes leading and trailing whitespace from the value
* striptags  
    - Removes any HTML tags from the value before rendering

#### Control Structures

```
{% if user %}
    Hello, {{ user }}!
{% else %}
    Hello, Stranger!
{% endif %}

<ul>
    {% for comment in comments %}
        <li>{{ comment }}</li>
    {% endfor %}
</ul>

{% macro render_comment(comment) %}
    <li>{{ comment }}</li>
{% endmacro %}

<ul>
    {% for comment in comments %}
        {{ render_comment(comment) }}
    {% endfor %}
</ul>

{% import 'macros.html' as macros %}
<ul>
    {% for comment in comments %}
        {{ macros.render_comment(comment) }}
    {% endfor %}
</ul>

{% include 'common.html' %}

```

#### template inheritance

base.html

```
<html>
<head>
    {% block head %}
        <title>{% block title %}{% endblock %} - My Application</title>
    {% endblock %}
</head>
<body>
    {% block body %}
    {% endblock %}
</body>
</html>
```

Base templates define blocks that can be overridden by derived templates. The Jinja2 **block** and **endblock** directives define blocks of content that are added to the base template.

```
{% extends "base.html" %}
    {% block title %}Index{% endblock %}
    {% block head %}
    {{ super() }}
<style>
</style>
    {% endblock %}
{% block body %}
    <h1>Hello, World!</h1>
{% endblock %}
```

#### Bootstrap Extension

```
pip install flask-bootstrap
from flask_bootstrap import Bootstrap
# ...
bootstrap = Bootstrap(app)
```

```
{% extends "bootstrap/base.html" %}
{% block title %}User Home Page{% endblock %}

{% block content %}
<div class="container">
    <div class="page-header">
        <h3>User Page</h3>
        <h1>Hello, {{ name | capitalize  }}!</h1>
    </div>
</div>
{% endblock %}
```

#### Error Handlers

```
@app.errorhandler(404)
def page_not_found(e):
    return render_template('404.html'), 404

@app.errorhandler(500)
def internal_server_error(e):
     render_template('500.html'), 500
```

```
{% extends "base.html" %}
{% block title %}Flasky - Page Not Found{% endblock %}
{% block page_content %}
    <div class="page-header">
        <h1>Not Found</h1>
    </div>
{% endblock %}
```

#### url_for() helper function

This function takes the view function name (or endpoint name for routes defined with app.**add_url_route()** ) as its single argument and returns its URL.

url_for('index') => /

url_for('index',_external=True) => http://localhost:5000/ 

url_for('user',name='john',_external=True) would return http://localhost:5000/user/john

url_for('user', name='john', page=2, version=1) would return /user/john?page=2&version=1

#### Static Files

Flask automatically supports static files by adding a special route to the application defined as /static/<filename>

url_for('static', filename='css/styles.css', _external=True) would return http://localhost:5000/static/css/styles.css

In its default configuration, Flask looks for static files in a subdirectory called static located in the application’s root folder.

```
<link rel="shortcut icon" href="{{ url_for('static', filename='favicon.ico') }}" type="image/x-icon">
```

#### Localization of Dates and Times with Flask-Moment

The server needs uniform time units that are independent of the location of each user, so typically Coordinated Universal Time (UTC) is used.

Flask-Moment is an extension for Flask applications that makes the integration of Moment.js into Jinja2 templates very easy.

```
pip install flask-moment
from flask_moment import Moment
moment = Moment(app)
```

```
{%block scripts %}
{{ super() }}
{{moment.include_moment() }}
{%endblock %}
```

```
from datetime import datetime
@app.route('/')
def index():
    return render_template('index.html',current_time=datetime.utcnow())
```

```
<p>The local date and time is {{ moment(current_time).format('LLL') }}.</p>
<p>That was {{ moment(current_time).fromNow(refresh=True) }}</p>
```

### Web Forms

```
pip install flask-wtf

app = Flask(__name__)
app.config['SECRET_KEY']='hard to guess string which is used as an encryption or signing key to improve the security of application';

```

**app.config** dictionary is a general purpose place to store configuration variables used by Flask.
**Flask-WTF** requires secret key to be configures to protect all forms against CSRF attacks

```
from flask_wtf import FlaskForm
from wtforms import StringField, SubmitFeild
from wtforms.validators import DataRequired

class NameForm(FlaskForm):
    name = StringField('Name',validators=[DataRequired()])
    submit = SubmitField('Submit')
```

* BooleanField 

    Checkbox with True and False values
* DateField
    
     Text field that accepts a datetime.date value in a given format
* DateTimeField
    
     Text field that accepts a datetime.datetime value in a given format
* DecimalField 
    
    Text field that accepts a decimal.Decimal value
* FileField 
    
    File upload field
* HiddenField

     Hidden text field
* MultipleFileField
    
     Multiple file upload field
* FieldList

     List of fields of a given type

* FloatField
    
    Text field that accepts a floating-point value
* FormField

     Form embedded as a field in a container form
* IntegerField

     Text field that accepts an integer value
* PasswordField

     Password text field
* RadioField
    
     List of radio buttons
* SelectField
    
     Drop-down list of choices
* SelectMultipleField 
    
    Drop-down list of choices with multiple selection
* SubmitField
    
    Form submission button
* StringField 
    
    Text field
* TextAreaField

     Multiple-line text field

#### WTForms validators

DataRequired 

    Validates that the field contains data after type conversion
Email

     Validates an email address
EqualTo 
    
    Compares the values of two fields; useful when requesting a password to be entered twice for confirmation
InputRequired
    
     Validates that the field contains data before type conversion
IPAddress
    
     Validates an IPv4 network address
Length

     Validates the length of the string entered
MacAddress

     Validates a MAC address
NumberRange

     Validates that the value entered is within a numeric range
Optional
    
     Allows empty input in the field, skipping additional validators
Regexp
    
     Validates the input against a regular expression
URL
    
     Validates a URL
UUID
    
     Validates a UUID
AnyOf
    
     Validates that the input is one of a list of possible values
NoneOf

     Validates that the input is none of a list of possible values

```
{% extends "bootstrap/base.html" %}
{% import "bootstrap/wtf.html" as wtf %}

{%block scripts %}
{{ super() }}
{%endblock %}

{% block content %}

<div class="page-header">
    <h1>Hello, {% if name %}{{ name }}{% else %}Stranger{% endif %}!</h1>
</div>

<form method="POST">
    {{ form.hidden_tag() }}
    {{ form.name.label }} {{ form.name(id='my-text-field') }}
    {{ form.submit() }}
</form>
{{ wtf.quick_form(form) }}
{%endblock %}
```

```
app = Flask(__name__)
app.config['SECRET_KEY'] = 'for csrf string'

@app.route('/register',methods=['GET','POST'])
def register():
    name = None
    form = NameForm()
    if form.validate_on_submit():
        name = form.name.data 
        form.name.data = ''
    return render_template('register.html',form=form,name=name)
```

```

from flask import Flask, render_template, session, redirect, url_for

@app.route('/register',methods=['GET','POST'])
def register():
    form = NameForm()
    if form.validate_on_submit():
        session['name'] = form.name.data 
        return redirect(url_for('index'))
    return render_template('register.html',form=form,name=session.get('name'))

@app.route('/')
def index():
    name = session.get('name') if session.get('name') else 'Stranger'
    return '<h1>Hello {}! {}</h1>'.format(name)
```

#### Flash Messages

```
flash('Looks like you have changed your name!')

{% block content %}
<div class="container">
    {% for message in get_flashed_messages() %}
        <div class="alert alert-warning">
            <button type="button" class="close" data-dismiss="alert">&times;</button>
            {{ message }}
        </div>
    {% endfor %}
    
    {% block page_content %}{% endblock %}
</div>
{% endblock %}
```
