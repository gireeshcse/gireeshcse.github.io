#### Basics of Canvas

    <canvas></canvas>

* Requires closing tag.
* Default size 300 X 150 (Width X Height)
* width and height properties are important(Not CSS height and Width)

        var canvas = document.getElementById('canvas');
        var context = canvas.getContext('2d'); // d is small letter
        context.fillStyle = rgba(0,0,200,0.5);
        context.fillRect(10,10,50,50);
        context.fillStyle = rgba(200,0,0,0.5);
        context.fillRect(30,30,50,50);

