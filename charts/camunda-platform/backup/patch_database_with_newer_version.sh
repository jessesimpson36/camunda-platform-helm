#!/bin/bash

kubectl set image statefulset/cpt-postgresql postgresql=docker.io/bitnami/postgresql:15.1.0-debian-11-r0
