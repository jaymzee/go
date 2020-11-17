Demonstrate having several go routines providing data to
the main GUI thread for display.
e.g. different parts of the screen could have different framerates.

In this example the draw method does not erase the screen.
Instead the erasing is left to DrawSevenSegment() function.
This way, only the part of the screen that pertains to the channel that
received data needs to be redrawn.
