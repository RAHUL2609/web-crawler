# web-crawler

Its a simple web-crawler that takes request from the Clients and return a taskUid in Response.
API to Submit a task to Crawl an URL :

curl -X POST -H 'Content-Type: application/json' -i http://localhost:8095/crawl/submit --data '{"url" : "http://news.yahoo.com/xyz"}'
Expected Response :
{
  "taskId": "3d65b138-109c-4945-b276-fcf95577e834"
}


Later clients can make another to fetch the dervired Urls from the TaskId.

curl -X GET -i http://localhost:8095/crawl/read/3d65b138-109c-4945-b276-fcf95577e834
Expected Response:
[
  "http://news.yahoo.com/xyz/tiger",
  "http://news.yahoo.com/xyz/cat"
]

Note : The deriveUrls which different domain names are ignored.

# Steps to Build and Deploy on K8s environment

Prerequisite : You must a setup per-build with docker and k8s along

Steps :

  1. cd build
  2. sh comp-build.sh (This would generate an docker image with name web-crawler)
  3. kubectl create namespace demo (Create a namespace under which with service would be deployed)
  4. kubectl delete -f service.yaml -n demo (delete older service if already exists)
  5. kubectl delete -f deployment.yaml -n demo (delete older deployments if already exists)
  6. validate health status : kubectl get svc -n demo | grep web-crawler
  7. validate health of the pod : kubectl get pods -n demo | grep web-crawler
