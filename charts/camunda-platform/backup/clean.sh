#!/bin/bash
kubectl delete jobs postgresql-backup
kubectl delete jobs postgresql-restore
kubectl delete jobs migrate-pv-data-cpt-postgresql-0
kubectl delete pvc postgresql-backup
kubectl delete pv postgresql-backup
kubectl delete pvc old-postgresql-data-dir
kubectl delete pv old-postgresql-data-dir
