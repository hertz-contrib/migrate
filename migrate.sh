#!/usr/bin/env bash

git clone https://github.com/hertz-contrib/migrate.git /tmp/hertz_migrate__
python3 /tmp/hertz_migrate__/migrate.py "$PWD"
rm -rf /tmp/hertz_migrate__
