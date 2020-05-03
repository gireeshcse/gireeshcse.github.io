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

