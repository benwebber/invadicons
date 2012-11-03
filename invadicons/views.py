#!/usr/bin/env python

# Standard library
import re

# Flask
from flask import Flask, send_file, request, make_response, render_template

# Invadicons
from invadicons import app
from invadicons.avatars import Invadicon

def validate_hash(md5hash):
    """
    Validate an MD5 hash.
    """
    expr = r'[a-f0-9]{32}'
    if re.match(expr, md5hash):
        return True
    return False

@app.route('/')
def index():
    return 'Must enter MD5 hash', 400

@app.route('/<md5hash>')
@app.route('/<md5hash>.png')
def show_avatar(md5hash):
    if not validate_hash(md5hash):
        return 'Invalid MD5 hash', 400
    
    invadicon = Invadicon(md5hash)
    
    if request.args.get('size'):
        invadicon.size = int(request.args.get('size'))

    strIO = invadicon.save()
    strIO.seek(0)
    return send_file(strIO, mimetype='image/png')

@app.route('/<md5hash>.svg')
def show_svg(md5hash):
    if not validate_hash(md5hash):
        return 'Invalid MD5 hash', 400
    invadicon = Invadicon(md5hash)
    r = make_response(render_template('svg.xml', invadicon = invadicon))
    r.mimetype = 'image/svg+xml'
    return r
