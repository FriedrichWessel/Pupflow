Pupflow is a plugin for [Blender](http://blender.org) to enable a user to control his rig with common USB joysticks ([Video](http://vimeo.com/40471709)).

# Folders
## Blender
This folder contains a test scene with a slighty modified [Big Buck Bunny](http://www.bigbuckbunny.org/) rig to match or requirements.
The scene contains a lot of "controllers", which are merely null objects. The translation (or rotation) of these null objects is (sometimes indirectly) mapped to the position and/or rotation of the bunny's bones. This way, the movement of a single controller can result in more complex animations. Take a look at how the ears work for example.

In the `scripts` folder resides `pupflow.py`. This script opens a UDP server socket and accepts JSON-formatted input. Each field of the JSON has to correspond to a named object of the scene. The code is stolen from Delicode's NI Mate Kinect Plugin but has been modified as well.

The script is loaded via another in-scene script. Under Mac OS X this mechanism only works if you start Blender from the console in this `blender` directory.

## inputserver
This daemon written in [Go](http://www.golang.org) uses [SDL](http://libsdl.org) to interface the connected joysticks and polls every axis, button and hat. Each value is looked up in the configuration, packed into JSON and sent to the server. Configuration is done via a webinterface which is available on <http://localhost:8181> by default. A axis can be mapped to any named scene object and the values can be remapped before being sent to Blender (effectively allowing the user to configure the mid-point, enable inverting etc). Rows are deleted by setting the target object to an empty string.

## screenkey
This is dirty. Oh god, this is so dirty. This is a last-minute hack to have some kind of greenscreen for a rehearsal shoot we did a while ago. Imagine two 640x480 rectangles sitting in the top left corner of your screen. This incredibly ugly Java app continuously grabs those 2 rectangles, keys the color green in the left image and layers the result over the right image. It works. But don't take this thing seriously.

If you need any help, feel free to contact us :)

# Tutorial
Coming up...

# Next up

* [OSC](https://en.wikipedia.org/wiki/Open_Sound_Control) support
* Easier configuration
* Pretty interface
