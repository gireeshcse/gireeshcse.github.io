from flask import Flask, render_template
from flask import request
from flask_bootstrap import Bootstrap
from flask_moment import Moment
from datetime import datetime

app = Flask(__name__)
bootstrap = Bootstrap(app)
moment = Moment(app)
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


@app.route('/home')
def home():
    return render_template("index.html",current_time=datetime.utcnow())

@app.route('/home/user/<name>')
def homeUser(name):
    return render_template("user.html",name=name,current_time=datetime.utcnow())

app.add_url_rule('/sample', '', sample)