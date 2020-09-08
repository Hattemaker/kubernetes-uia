#apply the pod resource
pod: service
	kubectl apply -f pod.yaml -n uia

#apply the service
service:
	kubectl apply -f service.yaml -n uia

#delete the pod resource
delete-pod:
	kubectl delete -f pod.yaml -n uia

#apply the depoyment resource
deployment: delete-pod service
	kubectl apply -f deployment.yaml -n uia

#upgrade the version in the pod
upgrade-pod:
	kubectl set image pod greeter greeter=stromland/greeter:2

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
