# KubeAware - A Kubernetes aware application
In order help applications who have yet to make it in the fast lane of _kubernetes ready_, 
this application looks at bridging the gap.

This application is made aware of the kubernetes environment and can react as you define it.

## Technically application
With Kubeaware, it can be used to fire events within other third party applications running.
This could be used to sync deployments, enforce configmap updates across deplyoments, use HTTP
requests to send config updates.

The benefit of using KubeAware with applications that do not have this functionality in built is
that it gives you the freedom of being language and platform agnostic.
## Example Deployment
Yiou can use KubeAware to wrap the running of the application so that you can ensure your 
deployments always run with the most up to date settings.
```Dockerfile
FROM moviestoreguy/kubeaware:0.0.1

# Copy in application
COPY . .

# Run all the install steps
RUN ...

# By default, the Entrypoint is set to KubeAware
# You can overwrite the CMD to run your application
CMD --run /path/to/application/start
```

Once you have created the container, you can define environment variables so that KubeAware knows
what should happen once different events happen.

You can also run KubeAware along side your deployments so that it will notify your app with your defined method.
