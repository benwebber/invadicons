#!/usr/bin/env python

from __future__ import division

import hashlib
import numpy
import png
import StringIO

from PIL import Image

class Invadicon(object):
    """
    An invadicon.
    """
    def __init__(self, md5hash, *args, **kwargs):
        self.md5hash = md5hash
        self.coords = get_coords(self.md5hash)
        self.imgarray = fill_pixel_array(self.coords)
        self.palette = choose_palette(self.md5hash)
        self.bg = '#%s' % self.md5hash[6:12]
        self.fg = '#%s' % self.md5hash[0:6]
        self.size = 100
    def save(self, *args, **kwargs):
        strIO = StringIO.StringIO()
        w = png.Writer(palette = self.palette, bitdepth=1, size=(10,10))
        w.write(strIO, self.imgarray)
        strIO.seek(0)
        # Resize the image to a given size
        img = Image.open(strIO)
        img = img.resize((self.size, self.size))
        outfile = StringIO.StringIO()
        img.save(outfile, format='png')
        return outfile
    def __repr__(self):
        return u"<Invadicon: %s>" % (self.md5hash)

def chunk(s, n):
    """
    Chunks a string into fragments n characters long.
    """
    for i in range(0, len(s), n):
        yield s[i:i+n]

def map_coords(s, xmax = 4, ymax = 8):
    """
    Transforms two hexadecimal digits (s) to (x, y) coordinates, e.g.,
    
        10 => (1, 0)
        b3 => (1, 3)
        2f => (2, 5)
    
    For values greater than xmax or ymax, use the remainder.
    """
    s = list(s)
    x = s[0]
    y = s[1]
    if x > xmax:
        x = int(x, 16) % xmax
    if y > ymax:
        y = int(y, 16) % ymax
    return (x,y)

def get_coords(md5hash):
    chunked = list(chunk(md5hash, 2))
    coords = []
    for c in chunked:
        coords.append(map_coords(c))
    return list(set(coords))

def choose_palette(md5hash):
    """
    Placeholder function for something more advanced.
    """
    return (pixelize(md5hash[6:12]), pixelize(md5hash[0:6]))

def pixelize(colour):
    """
    Takes a colour in #RRGGBB and returns a pixel value as an RGB tuple.
    """
    components = list(chunk(colour, 2))
    return tuple([int(x, 16) for x in components])
    
def fill_pixel_array(coords):
    """
    Construct an avatar pixel array 10px x 10px.
    """
    arr = numpy.zeros((10, 5))
    for c in coords:
        x = c[0]
        y = c[1]
        arr[y+1, x+1] = 1
    arr_mirror = numpy.fliplr(arr)
    arr = numpy.hstack((arr, arr_mirror))
    return arr
