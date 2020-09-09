#apply the pod resource
pod: namespace service
	kubectl apply -f pod.yaml

#apply the service
service:
	kubectl apply -f service.yaml

#delete the pod resource
delete-pod:
	kubectl delete -f pod.yaml

# delete the deployment resource
delete-deployment :
	kubectl delete -f deployment.yaml

#apply the depoyment resource
deployment: delete-pod service
	kubectl apply -f deployment.yaml

#upgrade the version in the pod that will fail, not existant versjon
pod-fail:
	kubectl set image pod greeter greeter=stromland/greeter:0

#set the image running for a deployment
v1:
	kubectl set image deployment greeter greeter=stromland/greeter:1
v2:
	kubectl set image deployment greeter greeter=stromland/greeter:2

#this should fail because the application will not start
v3:
	kubectl set image deployment greeter greeter=stromland/greeter:3

#this fails becase the image does not exist
v4:
	kubectl set image deployment greeter greeter=stromland/greeter:4

#every 0.5 second call repeater
watch:
	watch -n 0.5 http :8080

#Watch all the logs for all pods starting with greeter
logs:
	stern greeter -n uia

# Fetch a pod as json
get-pod:
	kubectl get pod greeter -o json | jq '.'

# Fetch several pods as json
get-pods:
	kubectl get pod -o json | jq '.'

# Delete the entire uia namespace
delete:
	kubectl delete ns uia

#Fetch a service as json
get-service:
	kubectl get service -o json | jq '.'

#Fetch all replicasets
get-rs:
	kubectl get replicaset

#Fetch all endpoints
get-endpoint:
	kubectl get endpoints

#Fetch the contexts
context:
	kubectl config get-contexts

# Create the namespace
namespace:
	kubectl create namespace uia

#Deplye greeter deployment and service
repeater:
	kubectl apply -f repeater.yaml

logs-repeater:
	stern repeater -n uia
