from flask import Flask
from flask import request

app = Flask(__name__)

@app.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    return '<h1>Hello World! {}</h1>'.format(user_agent)

def sample():
    return '<h1>Hello World! Sample</h1>'

def sampleHello():
    return '<h1>Hello World! Sample</h1>'

@app.route('/user/<name>')
def user(name):
    return '<h1>Hello, {}!</h1>'.format(name)

app.add_url_rule('/sample', '', sample)