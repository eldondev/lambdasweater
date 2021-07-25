# A lambda-based proxy with aspirations of being a more full-featured vpn

For motivation, check out [awslambdaproxy](/dan-v/awslambdaproxy) . The goal of this project 
is similar. However, rather than running a persistent ec2 machine, the proxy client will run locally.
Similarly, the proxy client will receive inbound HTTP/HTTPS/SOCKS connections (on localhost). The proxy client
will attempt normal nat-hole-punching techniques, such that the lambda can achieve a raw TCP connections.
As a last resort, the proxy client may rendezvous with the lambda client via a turn server.
