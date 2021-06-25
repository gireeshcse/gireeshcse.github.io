* Any value that is not a number, a string, a boolean , a symbol, null or undefined is an object
* Objects are mutable and its primitive types are immutable.
* Object values are references.
* Some global values,properties and functions

        Number.POSITIVE_INFINITY        parseInt()                  Number.MAX_SAFE_INTEGER     Date()
        Infinity                        parseFloat()                isNaN()                     RegExp()
        Number.MAX_VALUE                Number.isNaN()              eval()                      String()
        NaN                             Number.isInteger()          Math                        Object()
        Number.NaN                      Number.EPSILON              JSON


* Floating-point approximations are problematic for our programs

        let x = .3 - .2 
        let y = .2 - .1
        x === y  //false
        x === .1 //false

        let n = 123456.789;
        n.toFixed(0) //"123457"
        n.toFixed(2) //"123456.79"
        n.toFixed(5) //"123456.78900"
        n.toExponential(1) //"1.2e+5"
        n.toExponential(3) //"1.235e+5"
        n.toPrecision(4) //"1.235e+5"
        n.toPrecision(7) //"123456.8"
        n.toPrecision(10) //"123456.7890"

* Dates and times

        let timestamp = Date.now(); // current time as timestamp (number)
        let now = new  Date(); // Current time as a Date object
        let ms = now.getTime();
        let iso = now.toISOString();

* Strings

        let s = "Hello, world";
        s.substring(1,4); // ell
        s.slice(1,4); // ell
        s.slice(-3); // rld
        s.split(", "); // ["Hello", "world"]

        // Searching a string
        s.indexOf("l") // 2: position of first letter l
        s.indexOf("l", 3) // 3: position of first "l" at or after 3
        s.indexOf("zz")  // -1: s does not include the substring "zz"
        s.lastIndexOf("l") //  10: position of last letter l

        // Boolean searching functions in ES6 and later
        s.startsWith("Hell") // => true: the string starts with these
        s.endsWith("!") // => false: s does not end with that
        s.includes("or") // => true: s includes substring "or"

        // Creating modified versions of a string
        s.replace("llo", "ya") // => "Heya, world"
        s.toLowerCase() // => "hello, world"
        s.toUpperCase() // => "HELLO, WORLD"

        // Inspecting individual (16-bit) characters of a string
        s.charAt(0) // => "H": the first character
        s.charAt(s.length-1) // => "d": the last character
        s.charCodeAt(0) // => 72: 16-bit number at the specified position
        s.codePointAt(0) // => 72: ES6, works for codepoints > 16 bits

        // String padding functions in ES2017
        "x".padStart(3) // => " x": add spaces on the left to a length of 3
        "x".padEnd(3) // => "x ": add spaces on the right to a length of 3
        "x".padStart(3, "*") // => "**x": add stars on the left to a length of 3
        "x".padEnd(3, "-") // => "x--": add dashes on the right to a length of 3
        
        
        
        
        // Space trimming functions. trim() is ES5; others ES2019
        " test ".trim() // => "test": remove spaces at start and end
        " test ".trimStart() // => "test ": remove spaces on left. Also trimLeft
        " test ".trimEnd() // => " test": remove spaces at right. Also trimRight

        // Miscellaneous string methods
        s.concat("!") // => "Hello, world!": just use + operator instead
        "<>".repeat(5) // => "<><><><><>": concatenate n copies. ES6

        let name = "Ram";
        let greeting = `Hello ${name}.`; // Hello Ram.

        // pattern matching
        /^HTML/;
        /[1-9][0-9]*/;
        /\bjavascript\b/i; // Match "javascript" as a word, case-insensitive
        let text = "testing: 1, 2, 3";
        let pattern = /\d+/g;
        pattern.test(text); // => true: a match exists
        text.search(pattern) // => 9: position of first match
        text.match(pattern) // ["1", "2", "3"]: array of all matches
        text.replace(patten, "#") // "testing, #, #, #"
        text.split(/\D+/) // ["", "1", "2", "3"] : split on non digits
        text.split(/\d+/) // ["testing: ", ", ", ", ", ""]

* false values : undefined, null, 0, -0, NaN, ""

### Symbols

* Introduced in ES6 
* To serve as non-string property names. Generally, object property names are typically strings.
    
        let property = "name";
        let symname = Symbol("name");
        let obj = {};
        obj[property] = "old";
        obj[symname] = "new"; 
        obj[property] // "old"
        obj[symname] // "new"

* To obtain a Symbol value, we call the **Symbol()** function. This function **never returns the same value twice, even when called with the same argument.**
* These can be used to add a new property to an object and do not need to worry that we might be overwriting an existing property with same name accidentally.

Javascript defines a global Symbol registry.The **Symbol.for()** function takes a string argument and returns a Symbol value that is associated with the string you pass.If no Symbol is already associated with that string, a new one is created and returned; otherwise, the already existing Symbol is returned.**It returns the same value when called with the same string**

    let s = Symbol.for("shared");
    let t = Symbol.for("shared");
    s === t // true
    s.toString() // => "Symbol(shared)"
    Symbol.keyFor(t) // => "shared"

    let sym1 = Symbol("name")
    let sym2 = Symbol("name")
    sym1 === sym2 // false