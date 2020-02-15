#### Basics of Canvas

    <canvas></canvas>

* Requires closing tag.
* Default size 300 X 150 (Width X Height)
* width and height properties are important(Not CSS height and Width)

        var canvas = document.getElementById('canvas');
        var context = canvas.getContext('2d'); // d is small letter
        context.fillStyle = "rgba(0,0,200,0.5)"; // Quotes are required
        context.fillRect(10,10,50,50);
        context.fillStyle = "rgba(200,0,0,0.5)";
        context.fillRect(30,30,50,50);

* Canvas Uses grid system which starts with (0,0)
* 1 Unit =  1 Pixel
* CANVAS only supports rectangles and paths(list of points connected by lines)
* **fillRect** Filled Rectangle
* **strokeRect** Rectangular outline.
* **clearRect** clears specified rectangular area.(transparent)
* **beginPath()** creates new path
* **closePath**
* **stroke()** draws otline 
* **fill()** draws a solid shape by filling the paths content area.
* **lineTo** draws a line from current drawing position specified bt x and y.

**Triangle**

    context.beginPath();
    context.moveTo(100,50);
    context.lineTo(200,75);
    context.lineTo(200,25);
    // Any open shapes are closed automatically when fill() is called. No need to call closePath()
    context.fill(); 

**arc(x,y,radius,startAngle,endAngle,antiClockwise)**

* Center - x,y
* antiClockwise - boolean(true/false)
Angles have to be given in radians 
radian = ( Math.PI / 180 ) * degrees

            var startAngle = 0 * (Math.PI/180);
            var endAngle = 360 * (Math.PI/180);
            context.beginPath();
            context.arc(100,100,50,startAngle,endAngle,false);
            //Important to draw otline not even closePath draws the arc
            context.stroke();

**arcTo(x1,y1,x2,y2,radius)**

The  arc is tangential to both segments

            context.beginPath();
            context.strokeStyle = 'black';
            context.lineWidth = 5;
            // 1st Segment (200,20) & (200,130)
            // 2nd segment (50,20) & (200,130)
            context.moveTo(200, 20);
            context.arcTo(200,130, 50,20, 40);
            context.stroke();

**rect(x,y,width,height)**

    context.beginPath();
    context.rect(250,10,50,50);
    context.fill();

**fillText(text,x,y[,maxWidth])**

    context.font = '12px serif';
    context.fillText('(100,100)',100,100);

**Loading image**

    context.drawImage(image,x,y);
    context.drawImage(image,x,y,width,height);// For scaling image
    // Geting somepart of the image loading it on to the canvas
    drawImage(image,sourceX,sourceY,destinationX,destinationY,width,height);


    // Example
    var canvasImg = document.getElementById('canvas_image');
    var contextImg = canvasImg.getContext('2d');
    
    var img = new Image();
    img.addEventListener('load',function(){
        contextImg.drawImage(img,0,0,300,150);
    },false);
    img.src = "url_of_image";




