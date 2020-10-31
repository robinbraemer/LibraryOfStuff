# Open Virtual Network hosted on Kubernetes

- Namespace: `ovn`

## OVN central

Label the nodes where you want to run the OVN-central pods:
```shell script
kubectl label node <node> node-role.kubernetes.io/ovn-central="true"
```

Create SSL secret for OVN. (recommended!)
```shell script
docker run --rm -v "$PWD"/ssl:/etc/ovn kubeovn/kube-ovn:v1.6.0 bash generate-ssl.sh
sudo chown -R $UID ssl
kubectl -n ovn create secret generic ovn-tls \
--from-file=cacert=ssl/cacert.pem \
--from-file=cert=ssl/ovn-cert.pem \
--from-file=key=ssl/ovn-privkey.pem
```

Apply!
```shell script
kubectl apply -k .
```