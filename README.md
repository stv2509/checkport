# checkport checkport-daemonset.yml
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: checkport-config
  namespace: kube-system
  labels:
    app: checkport
data:
  # ======================================================
  config.hcl: |
    ConnectTimeout = 2
    RepeatTimeout = 60
    
    host "google.com" {
      proto = "udp"
      ports = ["53","68", 79]
    }
    
    host "example.com" {
      proto = "udp"
      ports = ["53","68", 6579]
    }
    
    host "ya.ru" {
      proto = "tcp"
      ports = [443,80, 8080]
    }

---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: checkport
  namespace: kube-system
  labels:
    application: checkport

spec:
  selector:
    matchLabels:
      name: checkport
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  template:
    metadata:
      labels:
        application: checkport
    spec:
      containers:
      - name: checkport
        image: dockerhub.test.local/library/checkport/checkport:0.1
        imagePullPolicy: Always
        volumeMounts:
        - name: checkport-config
          mountPath: /etc/checkport/
          readOnly: true
      terminationGracePeriodSeconds: 10
      volumes:
      - name: checkport-config
        configMap:
          name: checkport-config
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Exists
        effect: NoSchedule
