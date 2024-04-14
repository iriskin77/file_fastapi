#!/bin/bash

until aerich migrate && aerich upgrade
do
    echo "Waiting for db to be ready..."
    sleep 10
done
uvicorn apps.main:app --reload --host 0.0.0.0 --port 8090
#gunicorn main:app --workers 4 --worker-class uvicorn.workers.UvicornWorker  --bind=0.0.0.0:8090
