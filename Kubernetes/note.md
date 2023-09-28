# First you have to make a YAML file :
```vim pod.yaml```
# example of YAML file :
```apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels: 
    app: nginx
    tier: frotend
spec:
  containers: 
  - name: nginx
    image: nginx

```
# Then you can check your YAML file by this Command : 
```cat pod.yaml```

# to apply command :
```kubectl apply -f pod.yaml```
# to check the status : 
```kubectl get pods```

# for more details use :
```kubectl describe pod nginx```
# create nginx yaml file:
```apiVersion: v1
kind: Pod
metadata:
  name: nginx-2
  labels:
    env: production
spec: 
  containers:
    - name: nginx
      image: nginx```


