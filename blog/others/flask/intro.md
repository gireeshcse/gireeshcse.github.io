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

Decorators are a standard feature of the Python language. A common use of decorators is to register functions as handler functions to be invoked when certain events occur.

    @app.route('/')
    def index():
        return '<h1>Hello World!</h1>'

**add_url_rule()** arguments

* the URL, 
* the endpoint name, 
* and the view function


    def sample():
        return '<h1>Hello World! Sample</h1>'

    app.add_url_rule('/sample', 'sample', sample)

    @app.route('/user/<name>')
    def user(name):
        return '<h1>Hello, {}!</h1>'.format(name)

**Running the app**

`      
export FLASK_APP=app.py
flask run
`
the below is mostlty used in unit testing

`
if __name__ == '__main__':
    app.run()
`

### Debug Mode

* reloader
* debugger (web-based tool)

By default, debug mode is disabled.
To enable it, set a **FLASK_DEBUG=1** environment variable before invoking flask run

To allow web server listen for connections on the public network interface, enabling other computers in the same network to connect as well

    flask run --host 0.0.0.0


