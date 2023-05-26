#!/bin/bash

./clean.sh
./scale_down_components.sh
./apply.sh

kubectl scale --replicas=0 statefulset cpt-postgresql

# copy existing PVC to new PVC
kubectl create -f duplicate_pvc.yaml

./migrate-pvc.sh data-cpt-postgresql-0 old-postgresql-data-dir

kubectl delete jobs migrate-pv-data-cpt-postgresql-0

# should we delete the statefulset?
kubectl delete pvc data-cpt-postgresql-0

./patch_database_with_newer_version.sh
kubectl scale --replicas=1 statefulset cpt-postgresql

# wait for new database to come up which should create new pvc

./restore.sh

./scale_up_components.sh
