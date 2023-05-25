minikube delete
minikube start

curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.17.2 sh -
cd istio-1.17.2
export PATH=$PWD/bin:$PATH
istioctl manifest apply --set profile=demo -y
istioctl install --set profile=demo --set meshConfig.outboundTrafficPolicy.mode=REGISTRY_ONLY -y
cd ..
kubectl label namespace default istio-injection=enabled

kubectl apply -f joke-service.yml
kubectl apply -f ingress-gateway.yml
kubectl apply -f egress-gateway.yml
kubectl apply -f nginx-service.yml

minikube tunnel
