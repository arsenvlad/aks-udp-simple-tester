# AKS UDP Simple Tester Code Snippets

Very simple UDP client and server code snippets for trying UDP in AKS

client/server Go code snippets are based on <https://github.com/jpoon/kubernetes-udp>

```bash
# Configure SSH on nodes (if needed)
az vmss extension set  --resource-group MC_avaks1_avaks1_eastus2 --vmss-name aks-agentpool-96171998-vmss --name VMAccessForLinux --publisher Microsoft.OSTCExtensions --version 1.4 --protected-settings "{\"username\":\"azureuser\", \"ssh_key\":\"PUBLIC SSH KEY\"}"
az vmss update-instances --instance-ids * --resource-group MC_avaks1_avaks1_eastus2 --name aks-agentpool-96171998-vmss
```

```bash
# Run SSH node
kubectl run -it --rm aks-ssh --image=debian
apt-get update && apt-get install openssh-client -y
kubectl cp /.ssh/id_rsa aks-ssh:/id_rsa
```

```bash
# Build test apps
cd client
docker build -t avcode613/udp-client:v4 .
docker push avcode613/udp-client:v4

cd server
docker build -t avcode613/udp-server:v4 .
docker push avcode613/udp-server:v4
```

```bash
# Deploy server
kubectl apply -f server.yaml

# Look for server LoadBalancer IP
kubectl get services
```

```bash
# Update client.yaml with the public service IP of the server and deploy
kubectl apply -f client.yaml
```

```bash
# Look at server and client logs
kubectl get pods
kubectl logs --tail=-1 --follow=true <pod-name>
```