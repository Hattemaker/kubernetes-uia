#apply the pod resource
pod: namespace service
	kubectl apply -f pod.yaml -n uia

#apply the service
service:
	kubectl apply -f service.yaml -n uia

#delete the pod resource
delete-pod:
	kubectl delete -f pod.yaml -n uia

delete-deployment :
	kubectl delete -f deployment.yaml -n uia

#apply the depoyment resource
deployment: delete-pod service
	kubectl apply -f deployment.yaml -n uia

#upgrade the version in the pod
pod2:
	kubectl set image pod greeter greeter=stromland/greeter:2

pod1:
	kubectl set image pod greeter greeter=stromland/greeter:1


pod3:
	kubectl set image pod greeter greeter=stromland/greeter:3


#set the image running to 1
v1:
	kubectl set image deployment greeter greeter=stromland/greeter:1

#set the image running to 2
v2:
	kubectl set image deployment greeter greeter=stromland/greeter:2

#set the image running to 3
v3:
	kubectl set image deployment greeter greeter=stromland/greeter:3

#every 0.5 second send a http request to port 8080
watch:
	watch -n 0.5 http :8080

#Watch all the logs for all pods starting with greeter
logs:
	stern greeter -n uia

get-pod:
	kubectl get pod greeter -o json | jq '.'

get-pods:
	kubectl get pod -o json | jq '.'

delete:
	kubectl delete ns uia

get-service:
	kubectl get service -o json | jq '.'
get-rs:
	kubectl get replicaset

get-endpoint:
	kubectl get endpoints

context:
	kubectl config get-contexts

namespace:
	kubectl create namespace uia

