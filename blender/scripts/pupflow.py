# bl_info = {
#     "name": "Delicode NI mate",
#     "author": "Janne Karhu (jahka)",
#     "version": (1, 0),
#     "blender": (2, 6, 1),
#     "api": 35622,
#     "location": "Toolbar > Delicode NI mate",
#     "description": "Receives OSC data from the Delicode NI mate program",
#     "category": "Animation",
#     'wiki_url': '',
#     'tracker_url': ''
#     }

import bpy
from bpy.props import *
import socket
import math
from mathutils import Vector
import json

def remap(s1, s2, t1, t2, v):
    return (v - s1)/(s2-s1) * (t2-t1) + t1

def deg2rad(x):
    return x/360 * 2 * math.pi

class JoyReceiver():
    def run(self, objects):
        try:
            data = str(self.sock.recv(1024), "utf-8")
        except:
            return {'PASS_THROUGH'}

        while(True):
            lines = data.split('\n')

            for line in lines[:-1]:
                obj = json.loads(line)
                name = obj['Name']
                axis = obj['Axis']
                val = obj['Value']

                if not name in bpy.data.objects:
                    continue

                if not obj['Rotation']:
                    cnt = bpy.data.objects[name].location
                else:
                    cnt = bpy.data.objects[name].rotation_euler

                # Apparently, manipulating Vector elements is mighty expensive
                if axis == "x":
                    cnt = Vector((val, cnt[1], cnt[2]))
                elif axis == "y":
                    cnt = Vector((cnt[0], val, cnt[2]))
                elif axis == "z":
                    cnt = Vector((cnt[0], cnt[1], val))

                if not obj['Rotation']:
                    bpy.data.objects[name].location = cnt
                    pass
                else:
                    bpy.data.objects[name].rotation_euler = cnt
                    pass

            data = lines[-1]
            try:
                data += str(self.sock.recv(1024), "utf-8")
            except:
                break


    def __init__(self, UDP_PORT):
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.sock.setblocking(0)
        self.sock.bind(("localhost", UDP_PORT))

    def __del__(self):
        self.sock.close()

class PupflowStart(bpy.types.Operator):
    bl_idname = "wm.pupflowstart"
    bl_label = "PupFlow Start"
    bl_options = {'REGISTER'}

    enabled = False
    receiver = None
    timer = None

    def modal(self, context, event):
        if not __class__.enabled:
            return self.cancel(context)
        if event.type == 'TIMER':
            self.receiver.run(bpy.data.objects)
        return {'PASS_THROUGH'}

    def execute(self, context):
        __class__.enabled = True
        self.receiver = JoyReceiver(13370)

        self.timer = context.window_manager.event_timer_add(1/context.scene.render.fps, context.window)
        context.window_manager.modal_handler_add(self)
        return {'RUNNING_MODAL'}

    def cancel(self, context):
        __class__.enabled = False
        context.window_manager.event_timer_remove(self.timer)
        del self.receiver
        return {'CANCELLED'}

    @classmethod
    def disable(cls):
        cls.enabled = False

class PupflowStop(bpy.types.Operator):
    bl_idname = "wm.pupflowstop"
    bl_label = "PupFlow Stop"
    bl_options = {'REGISTER'}

    def execute(self, context):
        PupflowStart.disable()
        return {'FINISHED'}

def register():
    bpy.utils.register_module(__name__)

def unregister():
    bpy.utils.unregister_module(__name__)

if __name__ == "__main__":
    register()

