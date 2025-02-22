# Getting Started

## Before you Begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must be configured to communicate with your
cluster. If you do not already have a cluster, you can create one by installing [minikube] or [kind], or you can use one
of these Kubernetes playgrounds:

* [Katacoda](https://www.katacoda.com/courses/kubernetes/playground)
* [Play with Kubernetes](http://labs.play-with-k8s.com/)

You also need the `starboard` command to be installed, e.g. from the [binary releases](./installation/binary-releases.md).
By default, it will use the same configuration as kubectl to communicate with the cluster.

## Scanning Workloads

The easiest way to get started with Starboard is to use an imperative `starboard` command, which allows ad hoc scanning
of Kubernetes workloads deployed in your cluster.

To begin with, execute the following one-time setup command:

```
starboard install
```

The `install` subcommand creates the `starboard` namespace, in which Starboard executes Kubernetes jobs to perform
scans. It also sends custom security resources definitions to the Kubernetes API:

```console
$ kubectl api-resources --api-group aquasecurity.github.io
NAME                   SHORTNAMES    APIGROUP                 NAMESPACED   KIND
ciskubebenchreports    kubebench     aquasecurity.github.io   false        CISKubeBenchReport
configauditreports     configaudit   aquasecurity.github.io   true         ConfigAuditReport
kubehunterreports      kubehunter    aquasecurity.github.io   false        KubeHunterReport
vulnerabilityreports   vulns,vuln    aquasecurity.github.io   true         VulnerabilityReport
```

!!! tip
    There's also a `starboard uninstall` subcommand, which can be used to remove all resources created by Starboard.

As an example let's run in the current namespace an old version of `nginx` that we know has vulnerabilities:

```
kubectl create deployment nginx --image nginx:1.16
```

Run the vulnerability scanner to generate vulnerability reports:

```
starboard scan vulnerabilityreports deployment/nginx
```

Behind the scenes, by default this uses [Trivy][trivy] in Standalone mode to identify vulnerabilities in the container
images associated with the specified deployment. Once this has been done, you can retrieve the latest vulnerability
reports for this workload:

```
starboard get vulnerabilityreports deployment/nginx -o yaml
```

!!! tip
    Starboard relies on labels and label selectors to associate vulnerability reports with the specified Deployment.
    For a Deployment with *N* container images Starboard creates *N* instances of `vulnerabilityreports.aquasecurity.github.io`
    resources. In addition, each instance has the `starboard.container.name` label to associate it with a particular
    container's image. This means that the same data retrieved by the `starboard get vulnerabilities` subcommand can be
    fetched with the standard `kubectl get` command:

    ```console
    $ kubectl get vulnerabilityreports -o wide \
       -l starboard.resource.kind=Deployment,starboard.resource.name=nginx
    NAME                     REPOSITORY      TAG    SCANNER   AGE    CRITICAL   HIGH   MEDIUM   LOW   UNKNOWN
    deployment-nginx-nginx   library/nginx   1.16   Trivy     2m6s   3          40     24       90    0
    ```

    In this example, the `nginx` deployment has a single container called `nginx`, hence only one instance of the
    `vulnerabilityreports.aquasecurity.github.io` resource is created with the label `starboard.container.name=nginx`.

To read more about custom resources and label selectors check [custom resource definitions].

[trivy]: ./../integrations/vulnerability-scanners/trivy.md
[custom resource definitions]: ./../crds/index.md

Moving forward, let's take the same `nginx` Deployment and audit its Kubernetes configuration. As you remember we've
created it with the `kubectl create deployment` command which applies the default settings to the deployment descriptors.
However, we also know that in Kubernetes the defaults are usually the least secure.

Run the scanner to audit the configuration using [Polaris](./../integrations/config-checkers/polaris.md), which is the
default configuration checker:

```
starboard scan configauditreports deployment/nginx
```

Retrieve the configuration audit report:

```
starboard get configauditreports deployment/nginx -o yaml
```

or

```console
$ kubectl get configauditreport -o wide \
    -l starboard.resource.kind=Deployment,starboard.resource.name=nginx
NAME               SCANNER   AGE   DANGER   WARNING   PASS
deployment-nginx   Polaris   5s    0        8         9
```

## Generating HTML Reports

Once you scanned the `nginx` Deployment for vulnerabilities and checked its configuration you can generate an HTML
report of identified risks:

```
starboard report deployment/nginx > nginx.deploy.html
```

```
open nginx.deploy.html
```

![HTML Report](../images/html-report.png)

## What's Next?

To learn more about the available Starboard commands and scanners, such as [kube-bench][aqua-kube-bench] or
[kube-hunter][aqua-kube-hunter], use `starboard help`.

[minikube]: https://minikube.sigs.k8s.io/docs/
[kind]: https://kind.sigs.k8s.io/docs/
[polaris]: https://github.com/FairwindsOps/polaris
[aqua-kube-bench]: https://github.com/aquasecurity/kube-bench
[aqua-kube-hunter]: https://github.com/aquasecurity/kube-hunter
