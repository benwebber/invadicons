#!/usr/bin/env python

from __future__ import division

import hashlib
import numpy
import png
import StringIO
from PIL import Image

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

def choose_foreground(md5hash):
    """
    Placeholder function for something more advanced.
    """
    return pixelize(md5hash[0:6])

def choose_background(md5hash):
    """
    Placeholder function for something more advanced.
    """
    return pixelize(md5hash[6:12])
    
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

def generate_avatar(md5hash, size):
    """
    Generate an 8-bit avatar from a given MD5 hash. Returns a file-like object.
    """
    # Choose the avatar palette
    fg = choose_foreground(md5hash)
    bg = choose_background(md5hash)
    palette = [bg, fg]   
    # Populate the pixel array
    coords = get_coords(md5hash)
    imgarray = fill_pixel_array(coords)
    # Write the image to a file-like object
    strIO = StringIO.StringIO()
    w = png.Writer(palette = palette, bitdepth=1, size=(10,10))
    w.write(strIO, imgarray)
    strIO.seek(0)
    # Resize the image to a given size
    img = Image.open(strIO)
    img = img.resize((size, size))
    outfile = StringIO.StringIO()
    img.save(outfile, format='png')
    return outfile
