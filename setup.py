#!/bin/sh

# This trick helps get around problem systems where
# the shebang isn't interpreted, or where paths are non-standard
# As long as python is in the path, and a bourne compatible
# shell is executing this (or python is executing it directly) this should work.

magic='--calling-python-from-/bin/sh--'
"""exec" python -E "$0" "$@" """#$magic"
if __name__ == '__main__':
	import sys
	if sys.argv[-1] == '#%s' % magic:
		del sys.argv[-1]
del magic

from setuptools import setup

setup(
    name='grt',
    version='0.1',
    description='Gerrit Git Repository Helper',
    author='Matt Henkel',
    author_email='matt@menagerie.cc',
    url='http://guildencrantz.github.com/grt/',
    packages = ['grt', 'grt.commands'],
    scripts = ['grt/grt'],
    install_requires = ['gitpython', 'paramiko'],
    classifiers = [
        "Development Status :: 3 - Alpha",
        "Environment :: Console",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: Apache Software License",
    ],
)


# vim: ts=4 sw=4 sts=4 et ai :
