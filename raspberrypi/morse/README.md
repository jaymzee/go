# morse code for the Raspberry Pi
## Usage
First use wiringPi's gpio command to add the GPIO pin
to the exports in the `/sys/class/gpio` directory.
For example:
```
gpio export 6 out
```

To send a message on the exported GPIO pin using morse code:
```
morse -p 6 hello world.
```
The above command will use GPIO pin 6 and send the message "hello world."
