from flask import Flask, render_template, session,redirect, url_for, flash
from flask import request
from flask_bootstrap import Bootstrap
from flask_moment import Moment
from datetime import datetime

from forms import NameForm

app = Flask(__name__)
bootstrap = Bootstrap(app)
moment = Moment(app)
app.config['SECRET_KEY'] = 'for csrf string'

@app.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    name = session.get('name') if session.get('name') else 'Stranger'
    return render_template('home.html',name=name,user_agent=user_agent)

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

@app.route('/register',methods=['GET','POST'])
def register():
    form = NameForm()
    if form.validate_on_submit():
        session['name'] = form.name.data 
        flash('Looks like you have changed your name!')
        return redirect(url_for('index'))
    return render_template('register.html',form=form,name=session.get('name'))

app.add_url_rule('/sample', '', sample)