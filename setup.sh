#!/bin/bash
set -euxo pipefail

python3 -m venv .venv
source .venv/bin/activate
.venv/bin/python3 -m pip install --upgrade pip
pip3 install -r requirements.txt
deactivate

