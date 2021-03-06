# Invadicons

Invadicons are 8-bit avatars that look kind of like Space Invaders.

This is an open-source clone of [Gravatar's retro avatars](https://en.gravatar.com/site/implement/images/#default-image), written in Python. Invadicons are served by a [Flask](http://flask.pocoo.org/) app.

## Examples

These are some of the invadicons you get when you hash `00000000000000000000000000000000` repeatedly.

![11ac68eee8398ae00e9f6b11b22f7efd](https://github.com/benwebber/invadicons/raw/master/doc/img/11ac68eee8398ae00e9f6b11b22f7efd.png)
![125a98d55e6aa3b1621f0c73554ec38d](https://github.com/benwebber/invadicons/raw/master/doc/img/125a98d55e6aa3b1621f0c73554ec38d.png)
![dff5a32d020cd482320a7d7ef3dde23c](https://github.com/benwebber/invadicons/raw/master/doc/img/dff5a32d020cd482320a7d7ef3dde23c.png)

## Usage

Request an invadicon by passing an MD5 hash to the application. Invadicons are served as PNG or SVG files. The default format is PNG, and the default size is 100px by 100px.

### PNG

Optionally specify the `.png` suffix, or a `size` in pixels.

    http://invadicons.example.org/11ac68eee8398ae00e9f6b11b22f7efd
    http://invadicons.example.org/11ac68eee8398ae00e9f6b11b22f7efd.png
    http://invadicons.example.org/11ac68eee8398ae00e9f6b11b22f7efd?size=128

### SVG

Specify the `.svg` suffix.

    http://invadicons.example.org/11ac68eee8398ae00e9f6b11b22f7efd.svg
