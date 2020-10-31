# kube-ovn deploy guide

- Namespace: kube-ovn

Labels the OVN master nodes:
> See [how to](https://github.com/alauda/kube-ovn/blob/master/docs/high-available.md)
run ovn db in HA
```shell script
kubectl label node <Node on which to deploy OVN DB> kube-ovn/role=master
```

Create secret to use SSL in OVN.
```shell script
docker run --rm -v "$PWD"/ssl:/etc/ovn kubeovn/kube-ovn:v1.6.0 bash generate-ssl.sh
sudo chown -R $UID ssl
kubectl create secret generic kube-ovn-tls \
--from-file=cacert=ssl/cacert.pem \
--from-file=cert=ssl/ovn-cert.pem \
--from-file=key=ssl/ovn-privkey.pem
```

As of now we can't edit container env vars in a base resource.
We need to generate all and edit it manually.
Generate `all.yaml`.
```shell script
kustomize build > all.yaml && \
go run enable_ssl.go
```

Apply!
```shell script
kubectl apply -f all.yaml
```