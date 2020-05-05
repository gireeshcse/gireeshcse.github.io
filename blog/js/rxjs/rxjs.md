# ReactiveX

An API for asynchronous programming with observable streams.

Generally data of application may contain like this

Array is in-memory datastructure.

A stream of data with no limit of data.

## RxJS 

Reactive Extensions For JavaScript

RxJS is a library for composing asynchronous and event-based programs by using observable sequences.

It provides one core type, the **Observable**, satellite types **(Observer, Schedulers, Subjects)** and **operators** inspired by Array#extras (map, filter, reduce, every, etc) to allow handling asynchronous events as collections.

### The essential concepts in RxJS which solve async event management are:

* **Observable:** represents the idea of an invokable collection of future values or events.
* **Observer:** is a collection of callbacks that knows how to listen to values delivered by the Observable.
* **Subscription:** represents the execution of an Observable, is primarily useful for cancelling the execution.
* **Operators:** are pure functions that enable a functional programming style of dealing with collections with operations like map, filter, concat, reduce, etc.
* **Subject:** is the equivalent to an EventEmitter, and the only way of multicasting a value or event to multiple Observers.
* **Schedulers:** are centralized dispatchers to control concurrency, allowing us to coordinate when computation happens on e.g. setTimeout or requestAnimationFrame or others.


**Creating a project**

* npm init
* npm install rxjs --save

To support backward compartibality Such as ES5, we need compiler/transpiler such as **webpack** since it uses ES6 features and **webpack-dev-server** to serve our files to web browser for testing and **typescript** compiler to compile typescript code

* npm install webpack webpack-dev-server webpack-cli typescript typings ts-loader --save-dev

To Install typying for ES6 which rxjs uses

* ./node_modules/.bin/typings install dt~es6-shim --global --save

Sample : main.ts

        import { Observer, from, fromEvent} from 'rxjs';
        import { scan,throttleTime, count } from 'rxjs/operators';

        let numbers = [1,5,9,10];
        let source = from(numbers); // creates Observable from exsting array | promsie | iterable(generator)

        class MyObserver implements Observer<Number>{

            next(value)
            {
                console.log(value);
            }

            error(e)
            {
                console.log(e);
            }

            complete()
            {
                console.log("complete");
            }
        }

        source.subscribe(new MyObserver());
        source.subscribe(new MyObserver());

        // Observable from event
        fromEvent(document, 'click').subscribe(() => console.log('Clicked! on the document'));

        // Pure functions which are less error prone
        // To Isolate state
        fromEvent(document.getElementById("button"),'click')
            .pipe(scan(count => count+1 , 0))
            .subscribe(count => console.log(`clicked ${count} times`));

        // Flow
        // to control the rate of flow of events
        fromEvent(document.getElementById('refresh'),'click')
            .pipe(
                throttleTime(1000),
                scan(count => count + 1, 0)
            )
            .subscribe(count => console.log(`Clicked ${count} times`));


**Simple Observer**

```
import { Observable } from 'rxjs';

let numbers = [1,5,9,10];
let source = Observable.create(observer => {

    for(let n of numbers)
    {
        observer.next(n);
    }

    observer.complete();
});


source.subscribe(
    value => console.log(`value ${value}`),
    e => console.log(`error ${e}`),
    () => console.log("complete")
);


```
Using **from** api

```

import { from } from 'rxjs';

let numbers = [1,5,9,10];
// creates Observable from exsting array | promsie | iterable(generator)
let source = from(numbers); 

source.subscribe(
    value => console.log(`value ${value}`),
    e => console.log(`error ${e}`),
    () => console.log("complete")
);

```

To raise error

```
observer.error("something is wrong")
```
To create asynchronous behaviour which obervables handle

```
import { Observable } from 'rxjs';

let numbers = [1,5,9,10];
let source = Observable.create(observer => {

   let index = 0;

   let produceValue = () => {
       observer.next(numbers[index++]);

       if(index < numbers.length)
       {
           setTimeout(produceValue,2000);
       }else{
           observer.complete();
       }
   };

   produceValue();
});


source.subscribe(
    value => console.log(`value ${value}`),
    e => console.log(`error ${e}`),
    () => console.log("complete")
);
```

```
import { Observable } from 'rxjs';
import { map, filter } from 'rxjs/operators';

let numbers = [1,5,9,10];
let source = Observable.create(observer => {

   let index = 0;

   let produceValue = () => {
       observer.next(numbers[index++]);

       if(index < numbers.length)
       {
           setTimeout(produceValue,2000);
       }else{
           observer.complete();
       }
   };

   produceValue();
}).pipe(
    map((value:number) => sum(value)),
    filter(value => value > 10)
);

function sum(value:number):number
{
    return value * 10;
}

source.subscribe(
    value => console.log(`value ${value}`),
    e => console.log(`error ${e}`),
    () => console.log("complete")
);
```