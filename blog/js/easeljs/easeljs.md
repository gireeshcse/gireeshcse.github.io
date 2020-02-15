### About EaselJS

* In EaselJs a **Stage** is root level container for *display list*.
* It wraps a CANVAS element and adds **DiplayObject** intances as children.
* Supports
    * **Bitmap**(Images)
    * **Shape** and **Graphics** (Vector Graphics)
    * Animated bitmaps using **SpriteSheet** and **Sprite**
    * Simple text instances using **Text**
    * Containers that hold oter DisplayObjects using **Container**
    * Control HTML DOM elements using **DOMElement**

**Note**: All display objects can be added to the stage as children, or drawn to a canvas directly.

### Example ()

        <script src="https://code.createjs.com/1.0.0/easeljs.min.js"></script>


        //Create a stage by getting a reference to the canvas
        var stage = new createjs.Stage("demoCanvas");

        //Create a Shape DisplayObject.
        circle = new createjs.Shape();
        circle.graphics.beginFill("red").drawCircle(0, 0, 40);

        //Set position of Shape instance.
        circle.x = circle.y = 50;

        //Add Shape instance to stage display list.
        stage.addChild(circle);

        //Update stage will render next frame
        stage.update();

        displayObject.addEventListener("click", handleClick);

        function handleClick(event){
            // Click happenened
        }

        displayObject.addEventListener("mousedown", handlePress);

        function handlePress(event) {
            // A mouse press happened.
            // Listen for mouse move while the mouse is down:
            event.addEventListener("mousemove", handleMove);
        }

        function handleMove(event) {
            
        }

         //Update stage will render next frame
         createjs.Ticker.addEventListener("tick", handleTick);

        function handleTick() {
        //Circle will move 10 units to the right.
            circle.x += 10;
            //Will cause the circle to wrap back
            if (circle.x > stage.canvas.width) { circle.x = 0; }
            stage.update();
        }



