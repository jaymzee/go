# morse code
## Usage
Use the wiringPi gpio command to export GPIO pins to
the `/sys/class/gpio` directory.

To send a message with morse:
```
morse -p 6 hello world.
```
The above will use GPIO pin 6 and send the message "hello world."
by writing to the file `/sys/class/gpio/gpio6/value`
