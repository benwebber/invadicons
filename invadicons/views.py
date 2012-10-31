#!/usr/bin/env python

# Standard library
import re

# Flask
from flask import Flask, send_file, request

# Invadicons
from invadicons import app
from invadicons.avatars import generate_avatar

def validate_hash(hash):
    """
    Validate an MD5 hash.
    """
    expr = r'[a-f0-9]{32}'
    if re.match(expr, hash):
        return True
    return False

@app.route('/')
def index():
    return 'Must enter MD5 hash', 400

@app.route('/<hash>')
@app.route('/<hash>.png')
def show_avatar(hash):
    if not validate_hash(hash):
        return 'Invalid MD5 hash', 400
    if request.args.get('size'):
        size = int(request.args.get('size'))
    else:
        size = 100
    strIO = generate_avatar(hash, size)
    strIO.seek(0)
    return send_file(strIO, mimetype='image/png')
